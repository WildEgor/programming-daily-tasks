CREATE DATABASE point_db;

CREATE SCHEMA IF NOT EXISTS point;

-- Represent pickpoint discount
CREATE TABLE IF NOT EXISTS point.pickpoint_tariff (
    id BIGSERIAL PRIMARY KEY,
    pickpoint_id BIGSERIAL NOT NULL,
    percent DECIMAL(5, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TRUNCATE TABLE point.pickpoint_tariff;

-- Percents for pickpoints by time
INSERT INTO point.pickpoint_tariff (pickpoint_id, percent, created_at)
VALUES
    (1, 2.0, '2023-02-10 00:00:00'),
    (2, 2.0, '2023-02-10 00:00:00'),
    (3, 2.0, '2023-02-10 00:00:00'),
    (3, 5.0, '2023-03-10 00:00:00'),
    (4, 2.0, '2023-02-10 00:00:00'),
    (5, 2.0, '2023-02-10 00:00:00'),
    (6, 3.0, '2023-02-10 00:00:00'),
    (7, 1.0, '2023-03-10 00:00:00'),
    (8, 5.0, '2023-03-10 00:00:00'),
    (8, 8.0, '2023-04-10 00:00:00');

-- Get
SELECT pt.*
FROM point.pickpoint_tariff pt
         INNER JOIN (
    -- Get last (current) percent for point
    SELECT pickpoint_id, MAX(created_at) AS max_created_at
    FROM point.pickpoint_tariff
    GROUP BY pickpoint_id
) max_dates ON pt.pickpoint_id = max_dates.pickpoint_id AND pt.created_at = max_dates.max_created_at
ORDER BY created_at;