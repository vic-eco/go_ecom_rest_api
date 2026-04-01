-- name: ListProducts :many
SELECT *
FROM products;

-- name: FindProductByID :one
SELECT *
FROM products
WHERE id = $1;

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE(sqlc.narg(name), name),
    price_in_cents = COALESCE(sqlc.narg(price_in_cents), price_in_cents),
    quantity = COALESCE(sqlc.narg(quantity), quantity)
WHERE id = sqlc.arg(id)
    RETURNING *;

-- name: CreateProduct :one
INSERT INTO products(
    name,
    price_in_cents,
    quantity
) VALUES ($1, $2, $3) RETURNING *;