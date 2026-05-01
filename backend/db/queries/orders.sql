-- name: CreateOrder :exec
INSERT INTO orders (
    id, canal, cantidad, company, cp, created_at, days_to_delivery, error, error_message,
    fecha_compra, fecha_estimada, fulfillment_type, is_flash, is_marketplace,
    no_pedido, plan, product_type, sku, store_selected, tipo_pago, edd1, edd2
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22
) ON CONFLICT (id) DO NOTHING;

-- name: ListOrders :many
SELECT * FROM orders
WHERE 
    (canal = $1 OR $1 = '') AND
    (company = $2 OR $2 = '') AND
    (fulfillment_type = $3 OR $3 IS NULL OR $3 = '') AND
    (product_type = $4 OR $4 IS NULL OR $4 = '')
ORDER BY id
LIMIT $5 OFFSET $6;

-- name: CountOrders :one
SELECT COUNT(*) FROM orders
WHERE 
    (canal = $1 OR $1 = '') AND
    (company = $2 OR $2 = '') AND
    (fulfillment_type = $3 OR $3 IS NULL OR $3 = '') AND
    (product_type = $4 OR $4 IS NULL OR $4 = '');

-- name: GetTotalOrders :one
SELECT COUNT(*) FROM orders;

-- name: GetOrdersWithErrorsCount :one
SELECT COUNT(*) FROM orders WHERE error IS NOT NULL AND error <> '' AND error <> '0';

-- name: GetOrdersByCanal :many
SELECT canal, COUNT(*) as count FROM orders GROUP BY canal;

-- name: GetOrdersByFulfillmentType :many
SELECT fulfillment_type, COUNT(*) as count FROM orders GROUP BY fulfillment_type;

-- name: GetOrdersByProductType :many
SELECT product_type, COUNT(*) as count FROM orders GROUP BY product_type;

-- name: GetDistinctCanals :many
SELECT DISTINCT canal FROM orders ORDER BY canal;

-- name: GetDistinctCompanies :many
SELECT DISTINCT company FROM orders ORDER BY company;

-- name: GetDistinctFulfillmentTypes :many
SELECT DISTINCT fulfillment_type FROM orders ORDER BY fulfillment_type;

-- name: GetDistinctProductTypes :many
SELECT DISTINCT product_type FROM orders ORDER BY product_type;
