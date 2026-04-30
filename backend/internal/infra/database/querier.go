package database

import (
	"context"
	"database/sql"
)

type UserQuerier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
}

type OrderQuerier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) error
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	CountOrders(ctx context.Context, arg CountOrdersParams) (int64, error)
	GetTotalOrders(ctx context.Context) (int64, error)
	GetOrdersWithErrorsCount(ctx context.Context) (int64, error)
	GetOrdersByCanal(ctx context.Context) ([]GetOrdersByCanalRow, error)
	GetOrdersByFulfillmentType(ctx context.Context) ([]GetOrdersByFulfillmentTypeRow, error)
	GetOrdersByProductType(ctx context.Context) ([]GetOrdersByProductTypeRow, error)
}

type Querier interface {
	UserQuerier
	OrderQuerier
}

var _ Querier = (*Queries)(nil)

type Queries struct {
	db *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}
