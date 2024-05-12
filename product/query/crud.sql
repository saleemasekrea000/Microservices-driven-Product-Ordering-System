-- name: GetProductById :one
SELECT * FROM product WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM product LIMIT $1 OFFSET $2;

-- name: CreateProduct :one
INSERT INTO product (name, image, price, updated_at,created_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;
