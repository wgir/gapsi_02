package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/user/gapsi_orders_api/internal/infra/http/common"
	"github.com/user/gapsi_orders_api/pkg/jwt"
)

func RegisterUserRoutes(r chi.Router, handler *AuthHandler) {
	r.Post("/register", handler.Register)
}

func RegisterAuthRoutes(r chi.Router, handler *AuthHandler, tokenHelper *jwt.TokenHelper) {
	r.Post("/login", handler.Login)
	r.Post("/refresh", handler.Refresh)
	r.Post("/logout", handler.Logout)

	r.With(common.AuthMiddleware(tokenHelper)).Get("/me", handler.Me)
}
