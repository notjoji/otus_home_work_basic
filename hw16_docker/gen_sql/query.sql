-- name: GetUsersAndProducts :many
SELECT o.id   AS orderId,
       u.name AS userName,
       p.name AS productName,
       p.price
FROM Users u
         JOIN Orders o ON o.user_id = u.id
         JOIN OrderProducts op ON op.order_id = o.id
         JOIN Products p ON p.id = op.product_id
ORDER BY o.id
LIMIT $1 OFFSET $2;

-- name: GetOrdersByUserId :many
SELECT o.*
FROM Orders o
         JOIN Users u ON o.user_id = u.id
WHERE u.id = $1;

-- name: GetUserStatistics :many
SELECT u.name,
       COUNT(o)                                       AS orderNumber,
       SUM(o.total_amount)                            AS orderTotalSum,
       to_char(AVG(o.total_amount), 'FM999999999.00') AS avgSum
FROM Users u
         JOIN Orders o ON o.user_id = u.id
GROUP BY u.name;

-- name: GetUsers :many
SELECT u.*
FROM Users u
ORDER BY u.id
LIMIT $1 OFFSET $2;

-- name: GetUserById :one
SELECT DISTINCT u.*
FROM Users u
WHERE u.id = $1;

-- name: CreateUser :one
INSERT INTO Users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id;

-- name: UpdateUser :one
UPDATE Users
SET name     = $1,
    email    = $2,
    password = $3
WHERE id = $4
RETURNING id;

-- name: DeleteUser :one
DELETE
FROM Users
WHERE id = $1
RETURNING id;

-- name: GetOrders :many
SELECT o.*
FROM Orders o
ORDER BY o.id
LIMIT $1 OFFSET $2;

-- name: GetOrderById :one
SELECT DISTINCT o.*
FROM Orders o
WHERE o.id = $1;

-- name: CreateOrder :one
INSERT INTO Orders (user_id, order_date, total_amount)
VALUES ($1, $2, $3)
RETURNING id;

-- name: UpdateOrder :one
UPDATE Orders
SET user_id      = $1,
    order_date   = $2,
    total_amount = $3
WHERE id = $4
RETURNING id;

-- name: DeleteOrder :one
DELETE
FROM Orders
WHERE id = $1
RETURNING id;

-- name: GetProducts :many
SELECT p.*
FROM Products p
ORDER BY p.id
LIMIT $1 OFFSET $2;

-- name: GetProductById :one
SELECT DISTINCT p.*
FROM Products p
WHERE p.id = $1;

-- name: CreateProduct :one
INSERT INTO Products (name, price)
VALUES ($1, $2)
RETURNING id;

-- name: UpdateProduct :one
UPDATE Products
SET name  = $1,
    price = $2
WHERE id = $3
RETURNING id;

-- name: DeleteProduct :one
DELETE
FROM Products
WHERE id = $1
RETURNING id;

-- name: CreateOrderProduct :one
INSERT INTO orderproducts (order_id, product_id)
VALUES ($1, $2)
RETURNING order_id;