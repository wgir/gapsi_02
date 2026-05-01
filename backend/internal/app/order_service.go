package app

import (
	"context"

	"github.com/user/gapsi_orders_api/internal/domain"
)

type OrderService struct {
	repo domain.OrderRepository
}

func NewOrderService(repo domain.OrderRepository) domain.OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) ListOrders(ctx context.Context, filters domain.OrderFilters) ([]domain.Order, int64, error) {
	if filters.Page <= 0 {
		filters.Page = 1
	}
	if filters.PageSize <= 0 {
		filters.PageSize = 10
	}

	return s.repo.List(ctx, filters)
}

func (s *OrderService) GetStats(ctx context.Context) (*domain.OrderStats, error) {
	return s.repo.GetStats(ctx)
}

func (s *OrderService) GetFilters(ctx context.Context) (*domain.OrderFiltersOptions, error) {
	return s.repo.GetFilters(ctx)
}
