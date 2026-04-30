package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/user/gapsi_orders_api/internal/infra/http/auth"
	"github.com/user/gapsi_orders_api/internal/infra/http/common"
	"github.com/user/gapsi_orders_api/internal/infra/http/orders"
	"github.com/user/gapsi_orders_api/internal/infra/logger"
	"github.com/user/gapsi_orders_api/pkg/jwt"
)

func SetupRouter(
	log *logger.Logger,
	tokenHelper *jwt.TokenHelper,
	authHandler *auth.AuthHandler,
	orderHandler *orders.OrderHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(common.RequestIDMiddleware)
	r.Use(common.LoggerMiddleware(log))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			auth.RegisterUserRoutes(r, authHandler)
		})

		r.Route("/auth", func(r chi.Router) {
			auth.RegisterAuthRoutes(r, authHandler, tokenHelper)
		})

		r.Route("/orders", func(r chi.Router) {
			orders.RegisterOrderRoutes(r, orderHandler, tokenHelper)
		})
	})

	return r
}
