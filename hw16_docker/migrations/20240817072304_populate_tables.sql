-- +goose Up
-- +goose StatementBegin
INSERT INTO Users (name, email, password)
VALUES ('Иванов И. И.', 'ivanov_i_i@mail.ru', 'Omakdmu768as'),
       ('Петров П. П.', 'petroff@gmail.com', 'JNlkalOliu56'),
       ('Синицина Д. Д.', 'sinitsyna_d@mail.ru', 'NKhaopdplmp123');

INSERT INTO Products (name, price)
VALUES ('Хлеб', 50),
       ('Масло', 200),
       ('Молоко', 60),
       ('Сыр', 250),
       ('Конфеты', 150),
       ('Фарш', 300),
       ('Сливки', 120),
       ('Арбуз', 275);

INSERT INTO Orders (user_id, order_date, total_amount)
VALUES (1, '2024-07-22 10:23:00+05', 250),
       (1, '2024-07-25 15:45:00+05', 330),
       (1, '2024-07-28 17:24:00+05', 400),
       (2, '2024-07-24 19:09:00+05', 525),
       (2, '2024-07-26 11:54:00+05', 400),
       (3, '2024-07-29 15:56:00+05', 400);

INSERT INTO OrderProducts (order_id, product_id)
VALUES (1, 1),
       (1, 2),
       (2, 3),
       (2, 5),
       (2, 7),
       (3, 4),
       (3, 8),
       (4, 1),
       (4, 2),
       (4, 5),
       (5, 4),
       (5, 5),
       (6, 4),
       (6, 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE Users CASCADE;
TRUNCATE Orders CASCADE;
TRUNCATE Products CASCADE;
TRUNCATE OrderProducts CASCADE;
-- +goose StatementEnd
