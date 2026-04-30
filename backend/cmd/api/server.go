package main

import (
	"database/sql"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/user/gapsi_orders_api/internal/app"
	"github.com/user/gapsi_orders_api/internal/infra/config"
	"github.com/user/gapsi_orders_api/internal/infra/database"
	apiHTTP "github.com/user/gapsi_orders_api/internal/infra/http"
	"github.com/user/gapsi_orders_api/internal/infra/http/auth"
	"github.com/user/gapsi_orders_api/internal/infra/http/orders"
	"github.com/user/gapsi_orders_api/internal/infra/logger"
	"github.com/user/gapsi_orders_api/pkg/jwt"
	"go.uber.org/zap"
)

type App struct {
	cfg *config.Config
	log *logger.Logger
	db  *sql.DB
	srv *http.Server
}

func InitializeApp() (*App, error) {
	// 1. Load Configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	// 2. Setup Logger
	zlog, err := logger.NewLogger(cfg.AppEnv)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize logger: %w", err)
	}

	// 3. Database Connection
	dbConfig := database.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.DBSSLMode,
	}
	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 4. Initialize Repositories
	queries := database.New(db)
	userRepo := database.NewUserRepository(queries)
	orderRepo := database.NewOrderRepository(queries)

	// 5. Initialize Services
	tokenHelper := jwt.NewTokenHelper(cfg.JWTSecret)
	authService := app.NewAuthService(
		userRepo,
		tokenHelper,
		time.Duration(cfg.JWTAccessTTL)*time.Minute,
		time.Duration(cfg.JWTRefreshTTL)*24*time.Hour,
	)
	orderService := app.NewOrderService(orderRepo)

	// 6. Initialize Handlers and Router
	authHandler := auth.NewAuthHandler(authService)
	orderHandler := orders.NewOrderHandler(orderService)
	router := apiHTTP.SetupRouter(zlog, tokenHelper, authHandler, orderHandler)

	// Add Health Check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 7. Setup HTTP Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	return &App{
		cfg: cfg,
		log: zlog,
		db:  db,
		srv: srv,
	}, nil
}

func (a *App) Run() {
	a.log.Info("Starting Gapsi Orders API", zap.String("port", a.cfg.ServerPort))
	a.log.Info("Connected to PostgreSQL database")

	serverErrors := make(chan error, 1)

	go func() {
		a.log.Info("server listening", zap.String("addr", a.srv.Addr))
		serverErrors <- a.srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		a.log.Fatal("server error", zap.Error(err))

	case sig := <-shutdown:
		a.log.Info("start shutdown", zap.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := a.srv.Shutdown(ctx); err != nil {
			a.log.Error("graceful shutdown did not complete in time", zap.Error(err))
			if err := a.srv.Close(); err != nil {
				a.log.Fatal("could not stop server gracefully", zap.Error(err))
			}
		}

		if err := a.db.Close(); err != nil {
			a.log.Error("could not close database connection gracefully", zap.Error(err))
		}
	}

	a.log.Info("server stopped gracefully")
	a.log.Sync()
}
