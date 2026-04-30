package orders

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	query := r.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(query.Get("pageSize"))
	if pageSize < 1 {
		pageSize = 10
	}

	filters := domain.OrderFilters{
		Canal:           query.Get("canal"),
		Company:         query.Get("company"),
		FulfillmentType: query.Get("fulfillmentType"),
		ProductType:     query.Get("productType"),
		Page:            page,
		PageSize:        pageSize,
	}

	orders, total, err := h.orderService.ListOrders(r.Context(), filters)
	if err != nil {
		common.RespondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"data":  orders,
		"total": total,
		"page":  page,
		"size":  pageSize,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *OrderHandler) Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.orderService.GetStats(r.Context())
	if err != nil {
		common.RespondWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
