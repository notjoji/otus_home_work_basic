-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS Orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES Users (id) ON DELETE CASCADE,
    order_date TIMESTAMP WITH TIME ZONE NOT NULL,
    total_amount BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS Products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS OrderProducts (
    order_id BIGINT REFERENCES Orders (id) ON DELETE CASCADE,
    product_id BIGINT REFERENCES Products (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS order_user_idx ON Orders (user_id);
CREATE INDEX IF NOT EXISTS order_product_idx ON OrderProducts (order_id, product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Orders CASCADE;
DROP TABLE IF EXISTS Products CASCADE;
DROP TABLE IF EXISTS OrderProducts CASCADE;
DROP TABLE IF EXISTS Users CASCADE;
-- +goose StatementEnd
