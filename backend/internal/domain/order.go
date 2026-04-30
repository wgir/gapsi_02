package domain

import (
	"context"
)

type Order struct {
	ID              string `json:"id"`
	Canal           string `json:"canal"`
	Cantidad        int    `json:"cantidad"`
	Company         string `json:"company"`
	CP              string `json:"cp"`
	CreatedAt       string `json:"createdAt"`
	DaysToDelivery  string `json:"daysToDelivery"`
	Error           string `json:"error"`
	ErrorMessage    string `json:"errorMessage"`
	FechaCompra     string `json:"fechaCompra"`
	FechaEstimada   string `json:"fechaEstimada"`
	FulfillmentType string `json:"fulfillmentType"`
	IsFlash         bool   `json:"isFlash"`
	IsMarketplace   bool   `json:"isMarketplace"`
	NoPedido        string `json:"noPedido"`
	Plan            string `json:"plan"`
	ProductType     string `json:"productType"`
	SKU             string `json:"sku"`
	StoreSelected   string `json:"storeSelected"`
	TipoPago        string `json:"tipoPago"`
	Edd1            string `json:"edd1"`
	Edd2            string `json:"edd2"`
}

type OrderFilters struct {
	Canal           string `json:"canal"`
	Company         string `json:"company"`
	FulfillmentType string `json:"fulfillment_type"`
	ProductType     string `json:"product_type"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
}

type OrderStats struct {
	TotalOrders            int64            `json:"total_orders"`
	BreakdownByCanal       map[string]int64 `json:"breakdown_by_canal"`
	BreakdownByFulfillment map[string]int64 `json:"breakdown_by_fulfillment"`
	BreakdownByProductType map[string]int64 `json:"breakdown_by_product_type"`
	PercentageWithErrors   float64          `json:"percentage_with_errors"`
}

type OrderFiltersOptions struct {
	Channels         []string `json:"channels"`
	Companies        []string `json:"companies"`
	FulfillmentTypes []string `json:"fulfillmentTypes"`
	ProductTypes     []string `json:"productTypes"`
}

type OrderRepository interface {
	List(ctx context.Context, filters OrderFilters) ([]Order, int64, error)
	GetStats(ctx context.Context) (*OrderStats, error)
	GetFilters(ctx context.Context) (*OrderFiltersOptions, error)
}

type OrderService interface {
	ListOrders(ctx context.Context, filters OrderFilters) ([]Order, int64, error)
	GetStats(ctx context.Context) (*OrderStats, error)
	GetFilters(ctx context.Context) (*OrderFiltersOptions, error)
}
