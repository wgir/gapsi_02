package database

import (
	"context"

	"github.com/user/gapsi_orders_api/internal/infra/database/sqlc"
)

type UserQuerier interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, email string) (sqlc.User, error)
	GetUserByID(ctx context.Context, id string) (sqlc.User, error)
}

type OrderQuerier interface {
	CreateOrder(ctx context.Context, arg sqlc.CreateOrderParams) error
	ListOrders(ctx context.Context, arg sqlc.ListOrdersParams) ([]sqlc.Order, error)
	CountOrders(ctx context.Context, arg sqlc.CountOrdersParams) (int64, error)
	GetTotalOrders(ctx context.Context) (int64, error)
	GetOrdersWithErrorsCount(ctx context.Context) (int64, error)
	GetOrdersByCanal(ctx context.Context) ([]sqlc.GetOrdersByCanalRow, error)
	GetOrdersByFulfillmentType(ctx context.Context) ([]sqlc.GetOrdersByFulfillmentTypeRow, error)
	GetOrdersByProductType(ctx context.Context) ([]sqlc.GetOrdersByProductTypeRow, error)
}
