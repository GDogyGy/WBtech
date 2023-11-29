-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments
(
    id            SERIAL PRIMARY KEY,
    transaction   VARCHAR(255) NOT NULL,
    request_id    VARCHAR(255),
    currency      VARCHAR(50)  NOT NULL,
    provider      VARCHAR(50)  NOT NULL,
    amount        integer      not null default 0,
    payment_dt    BIGINT      not null,
    bank          VARCHAR(255),
    delivery_cost integer      not null default 0,
    goods_total   NUMERIC      not null default 0,
    custom_fee    NUMERIC      not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
