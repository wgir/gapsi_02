package database

import (
	"context"

	"github.com/user/gapsi_orders_api/internal/domain"
)

type orderRepo struct {
	queries OrderQuerier
}

func NewOrderRepository(q OrderQuerier) domain.OrderRepository {
	return &orderRepo{queries: q}
}

func (r *orderRepo) List(ctx context.Context, filters domain.OrderFilters) ([]domain.Order, int64, error) {
	offset := (filters.Page - 1) * filters.PageSize

	dbOrders, err := r.queries.ListOrders(ctx, ListOrdersParams{
		Canal:           filters.Canal,
		Company:         filters.Company,
		FulfillmentType: filters.FulfillmentType,
		ProductType:     filters.ProductType,
		Limit:           int32(filters.PageSize),
		Offset:          int32(offset),
	})
	if err != nil {
		return nil, 0, err
	}

	count, err := r.queries.CountOrders(ctx, CountOrdersParams{
		Canal:           filters.Canal,
		Company:         filters.Company,
		FulfillmentType: filters.FulfillmentType,
		ProductType:     filters.ProductType,
	})
	if err != nil {
		return nil, 0, err
	}

	orders := make([]domain.Order, len(dbOrders))
	for i, dbO := range dbOrders {
		orders[i] = domain.Order{
			ID:              dbO.ID,
			Canal:           dbO.Canal,
			Cantidad:        int(dbO.Cantidad),
			Company:         dbO.Company,
			CP:              dbO.Cp,
			CreatedAt:       dbO.CreatedAt,
			DaysToDelivery:  dbO.DaysToDelivery.String,
			Error:           dbO.Error.String,
			ErrorMessage:    dbO.ErrorMessage.String,
			FechaCompra:     dbO.FechaCompra.String,
			FechaEstimada:   dbO.FechaEstimada.String,
			FulfillmentType: dbO.FulfillmentType.String,
			IsFlash:         dbO.IsFlash,
			IsMarketplace:   dbO.IsMarketplace,
			NoPedido:        dbO.NoPedido.String,
			Plan:            dbO.Plan.String,
			ProductType:     dbO.ProductType.String,
			SKU:             dbO.Sku.String,
			StoreSelected:   dbO.StoreSelected.String,
			TipoPago:        dbO.TipoPago.String,
			Edd1:            dbO.Edd1.String,
			Edd2:            dbO.Edd2.String,
		}
	}

	return orders, count, nil
}

func (r *orderRepo) GetStats(ctx context.Context) (*domain.OrderStats, error) {
	total, err := r.queries.GetTotalOrders(ctx)
	if err != nil {
		return nil, err
	}

	withError, err := r.queries.GetOrdersWithErrorsCount(ctx)
	if err != nil {
		return nil, err
	}

	byCanal, err := r.queries.GetOrdersByCanal(ctx)
	if err != nil {
		return nil, err
	}

	byFulfill, err := r.queries.GetOrdersByFulfillmentType(ctx)
	if err != nil {
		return nil, err
	}

	byProduct, err := r.queries.GetOrdersByProductType(ctx)
	if err != nil {
		return nil, err
	}

	stats := &domain.OrderStats{
		TotalOrders:            total,
		BreakdownByCanal:       make(map[string]int64),
		BreakdownByFulfillment: make(map[string]int64),
		BreakdownByProductType: make(map[string]int64),
	}

	if total > 0 {
		stats.PercentageWithErrors = float64(withError) / float64(total) * 100
	}

	for _, c := range byCanal {
		stats.BreakdownByCanal[c.Canal] = c.Count
	}
	for _, f := range byFulfill {
		stats.BreakdownByFulfillment[f.FulfillmentType.String] = f.Count
	}
	for _, p := range byProduct {
		stats.BreakdownByProductType[p.ProductType.String] = p.Count
	}

	return stats, nil
}
