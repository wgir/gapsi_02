package database

import (
	"context"
	"database/sql"
)

type CreateOrderParams struct {
	ID              string
	Canal           string
	Cantidad        int32
	Company         string
	Cp              string
	CreatedAt       string
	DaysToDelivery  sql.NullString
	Error           sql.NullString
	ErrorMessage    sql.NullString
	FechaCompra     sql.NullString
	FechaEstimada   sql.NullString
	FulfillmentType sql.NullString
	IsFlash         bool
	IsMarketplace   bool
	NoPedido        sql.NullString
	Plan            sql.NullString
	ProductType     sql.NullString
	Sku             sql.NullString
	StoreSelected   sql.NullString
	TipoPago        sql.NullString
	Edd1            sql.NullString
	Edd2            sql.NullString
}

const createOrder = `-- name: CreateOrder :exec
INSERT INTO orders (
    id, canal, cantidad, company, cp, created_at, days_to_delivery, error, error_message,
    fecha_compra, fecha_estimada, fulfillment_type, is_flash, is_marketplace,
    no_pedido, plan, product_type, sku, store_selected, tipo_pago, edd1, edd2
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22
) ON CONFLICT (id) DO NOTHING
`

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) error {
	_, err := q.db.ExecContext(ctx, createOrder,
		arg.ID, arg.Canal, arg.Cantidad, arg.Company, arg.Cp, arg.CreatedAt,
		arg.DaysToDelivery, arg.Error, arg.ErrorMessage, arg.FechaCompra,
		arg.FechaEstimada, arg.FulfillmentType, arg.IsFlash, arg.IsMarketplace,
		arg.NoPedido, arg.Plan, arg.ProductType, arg.Sku, arg.StoreSelected,
		arg.TipoPago, arg.Edd1, arg.Edd2,
	)
	return err
}

type ListOrdersParams struct {
	Canal           string
	Company         string
	FulfillmentType string
	ProductType     string
	Limit           int32
	Offset          int32
}

const listOrders = `-- name: ListOrders :many
SELECT id, canal, cantidad, company, cp, created_at, days_to_delivery, error, error_message, fecha_compra, fecha_estimada, fulfillment_type, is_flash, is_marketplace, no_pedido, plan, product_type, sku, store_selected, tipo_pago, edd1, edd2 FROM orders
WHERE 
    (canal = $1 OR $1 = '') AND
    (company = $2 OR $2 = '') AND
    (fulfillment_type = $3 OR $3 = '') AND
    (product_type = $4 OR $4 = '')
ORDER BY id
LIMIT $5 OFFSET $6
`

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders,
		arg.Canal,
		arg.Company,
		arg.FulfillmentType,
		arg.ProductType,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID, &i.Canal, &i.Cantidad, &i.Company, &i.Cp, &i.CreatedAt,
			&i.DaysToDelivery, &i.Error, &i.ErrorMessage, &i.FechaCompra,
			&i.FechaEstimada, &i.FulfillmentType, &i.IsFlash, &i.IsMarketplace,
			&i.NoPedido, &i.Plan, &i.ProductType, &i.Sku, &i.StoreSelected,
			&i.TipoPago, &i.Edd1, &i.Edd2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type CountOrdersParams struct {
	Canal           string
	Company         string
	FulfillmentType string
	ProductType     string
}

const countOrders = `-- name: CountOrders :one
SELECT COUNT(*) FROM orders
WHERE 
    (canal = $1 OR $1 = '') AND
    (company = $2 OR $2 = '') AND
    (fulfillment_type = $3 OR $3 = '') AND
    (product_type = $4 OR $4 = '')
`

func (q *Queries) CountOrders(ctx context.Context, arg CountOrdersParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOrders,
		arg.Canal,
		arg.Company,
		arg.FulfillmentType,
		arg.ProductType,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTotalOrders = `-- name: GetTotalOrders :one
SELECT COUNT(*) FROM orders
`

func (q *Queries) GetTotalOrders(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalOrders)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getOrdersWithErrorsCount = `-- name: GetOrdersWithErrorsCount :one
SELECT COUNT(*) FROM orders WHERE error IS NOT NULL AND error <> '' AND error <> '0'
`

func (q *Queries) GetOrdersWithErrorsCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getOrdersWithErrorsCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

type GetOrdersByCanalRow struct {
	Canal string
	Count int64
}

const getOrdersByCanal = `-- name: GetOrdersByCanal :many
SELECT canal, COUNT(*) as count FROM orders GROUP BY canal
`

func (q *Queries) GetOrdersByCanal(ctx context.Context) ([]GetOrdersByCanalRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersByCanal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrdersByCanalRow
	for rows.Next() {
		var i GetOrdersByCanalRow
		if err := rows.Scan(&i.Canal, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type GetOrdersByFulfillmentTypeRow struct {
	FulfillmentType sql.NullString
	Count           int64
}

const getOrdersByFulfillmentType = `-- name: GetOrdersByFulfillmentType :many
SELECT fulfillment_type, COUNT(*) as count FROM orders GROUP BY fulfillment_type
`

func (q *Queries) GetOrdersByFulfillmentType(ctx context.Context) ([]GetOrdersByFulfillmentTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersByFulfillmentType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrdersByFulfillmentTypeRow
	for rows.Next() {
		var i GetOrdersByFulfillmentTypeRow
		if err := rows.Scan(&i.FulfillmentType, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type GetOrdersByProductTypeRow struct {
	ProductType sql.NullString
	Count       int64
}

const getOrdersByProductType = `-- name: GetOrdersByProductType :many
SELECT product_type, COUNT(*) as count FROM orders GROUP BY product_type
`

func (q *Queries) GetOrdersByProductType(ctx context.Context) ([]GetOrdersByProductTypeRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersByProductType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOrdersByProductTypeRow
	for rows.Next() {
		var i GetOrdersByProductTypeRow
		if err := rows.Scan(&i.ProductType, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
