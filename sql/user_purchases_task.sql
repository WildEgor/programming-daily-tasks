CREATE DATABASE product_db;

CREATE SCHEMA IF NOT EXISTS product;

CREATE TABLE IF NOT EXISTS product.user (
    id BIGSERIAL PRIMARY KEY,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    birth DATE
);

CREATE TABLE IF NOT EXISTS product.purchases (
    id BIGSERIAL PRIMARY KEY,
    sku BIGSERIAL,
    price NUMERIC,
    user_id INTEGER REFERENCES product.user (id),
    date date
);

CREATE TABLE IF NOT EXISTS product.ban_list (
    user_id INTEGER REFERENCES product.user (id) unique,
    date_from date
);

INSERT INTO product.user (firstname, lastname, birth)
VALUES
    ('Ivan', 'Petrov', '1996-05-01'),
    ('Alla', 'Petrova', '1999-06-01'),
    ('Anna', 'Petrova', '1990-10-02');

TRUNCATE product.purchases;

INSERT INTO product.purchases (sku, price, user_id, date)
VALUES
    (1, 5500, 1, '2021-02-15'),
    (1, 5700, 1, '2021-01-15'),
    (2, 4000, 1, '2021-02-14'),
    (3, 8000, 2, '2021-03-01'),
    (4, 400,  2,  '2021-03-02');

INSERT INTO product.ban_list (user_id, date_from)
VALUES
    (1, '2021-03-08');

-- Get all user-sku id pairs for all purchases made by user before ban. Sort by user id and sku
SELECT p.sku, p.user_id, p."date" FROM product.purchases p LEFT JOIN product.ban_list bl on p.user_id = bl.user_id WHERE bl.user_id IS NULL OR p."date" <= bl.date_from;
SELECT p.sku, p.user_id, p."date" FROM product.purchases p LEFT JOIN product.ban_list bl on p.user_id = bl.user_id and p."date" <= bl.date_from order by p.user_id, p.sku;

-- Find all users, which make purchases sum > 5000. Show in format [user_id] | [firstname] | [lastname] | [sum]
SELECT * FROM product."user" u INNER JOIN (SELECT user_id, SUM(price) as sum FROM product.purchases GROUP BY user_id) as user_purchases_sum ON u.id = user_purchases_sum.user_id AND user_purchases_sum.sum > 5000;

