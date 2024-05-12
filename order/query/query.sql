-- name: CreateOrder :exec
INSERT INTO "orders" ("total_price", "user_id", "updated_at")
VALUES ($1, $2, $3);

-- name: CreateOrderItems :exec
INSERT INTO "order_items" ("product_id", "order_id", "amount", "price")
VALUES ($1, $2, $3, $4);

-- name: GetOrder :many
SELECT * FROM "orders" 
LEFT JOIN "order_items" on "orders"."id" = "order_items"."order_id"
WHERE "orders"."id" = $1 ORDER BY "order_items"."id";    

-- name: GetOrders :many
SELECT * FROM "orders" 
JOIN "order_items" on "orders"."id" = "order_items"."order_id" ORDER BY "orders"."id" LIMIT $1 OFFSET $2 ;



-- name: GetOrderItems :many
SELECT * FROM "order_items" WHERE "id" = $1;

-- name: UpdateOrder :exec
UPDATE "orders" 
SET "total_price" = $1, "user_id" = $2, "created_at" = $3, "updated_at" = $4
WHERE "id" = $5;

-- name: UpdateOrderItems :exec
UPDATE "order_items" 
SET "product_id" = $1, "amount" = $2, "price" = $3
WHERE "id" = $4;

-- name: DeleteOrder :exec
DELETE FROM "orders" WHERE "id" = $1;

-- name: DeleteOrderItems :exec
DELETE FROM "order_items" WHERE "id" = $1;
