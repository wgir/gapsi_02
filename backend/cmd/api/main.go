package main

import (
	"context"
	"fmt"
	"log"
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

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// 2. Setup Logger
	zlog, err := logger.NewLogger(cfg.AppEnv)
	if err != nil {
		log.Fatalf("cannot initialize logger: %v", err)
	}
	defer zlog.Sync()
	zlog.Info("Starting Gapsi Orders API", zap.String("port", cfg.ServerPort))

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
		zlog.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close()
	zlog.Info("Connected to PostgreSQL database")

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

	// 7. Setup Server with Graceful Shutdown
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		zlog.Info("server listening", zap.String("addr", srv.Addr))
		serverErrors <- srv.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		zlog.Fatal("server error", zap.Error(err))

	case sig := <-shutdown:
		zlog.Info("start shutdown", zap.String("signal", sig.String()))

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Asking listener to shut down and load shed.
		if err := srv.Shutdown(ctx); err != nil {
			zlog.Error("graceful shutdown did not complete in time", zap.Error(err))
			if err := srv.Close(); err != nil {
				zlog.Fatal("could not stop server gracefully", zap.Error(err))
			}
		}

		if err := db.Close(); err != nil {
			zlog.Error("could not close database connection gracefully", zap.Error(err))
		}
	}

	zlog.Info("server stopped gracefully")
}
