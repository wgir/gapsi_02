package orders

import (
	"github.com/go-chi/chi/v5"
	"github.com/user/gapsi_orders_api/internal/infra/http/common"
	"github.com/user/gapsi_orders_api/pkg/jwt"
)

func RegisterOrderRoutes(r chi.Router, handler *OrderHandler, tokenHelper *jwt.TokenHelper) {
	r.Use(common.AuthMiddleware(tokenHelper))
	r.Post("/", handler.List)
	r.Get("/stats", handler.Stats)
	r.Get("/filters", handler.GetFilters)
}
