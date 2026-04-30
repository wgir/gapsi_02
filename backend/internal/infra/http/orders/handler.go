package orders

import (
	"net/http"

	"github.com/user/gapsi_orders_api/internal/domain"
	"github.com/user/gapsi_orders_api/internal/infra/http/common"
)

type OrderHandler struct {
	orderService domain.OrderService
}

func NewOrderHandler(orderService domain.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	var filters domain.OrderFilters
	if err := common.DecodeJSON(r, &filters); err != nil {
		// If body is empty or invalid, we use default filters
		filters = domain.OrderFilters{}
	}

	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.PageSize < 1 {
		filters.PageSize = 10
	}

	orders, total, err := h.orderService.ListOrders(r.Context(), filters)
	if err != nil {
		common.RespondWithError(w, r, err)
		return
	}

	common.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"data":      orders,
		"total":     total,
		"page":      filters.Page,
		"page_size": filters.PageSize,
	})
}

func (h *OrderHandler) Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.orderService.GetStats(r.Context())
	if err != nil {
		common.RespondWithError(w, r, err)
		return
	}

	common.RespondWithJSON(w, http.StatusOK, stats)
}
