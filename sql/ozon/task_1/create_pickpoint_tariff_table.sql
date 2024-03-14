CREATE TABLE IF NOT EXISTS pickpoint_tariff (
    id BIGSERIAL PRIMARY KEY,
    pickpoint_id BIGSERIAL NOT NULL,
    percent DECIMAL(5, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);