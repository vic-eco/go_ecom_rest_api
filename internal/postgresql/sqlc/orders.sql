-- name: CreateOrder :one
INSERT INTO orders(
    customer_id
) VALUES ($1) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items(
    order_id,
    product_id,
    quantity,
    price_cents
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: FindOrderByID :one
SELECT *
FROM orders
WHERE id = $1;

-- name: GetOrderItemsByOrderID :many
SELECT
    oi.order_id,
    p.name AS product_name,
    oi.quantity,
    oi.price_cents
FROM order_items oi
JOIN products p ON p.id = oi.product_id
WHERE oi.order_id = $1;