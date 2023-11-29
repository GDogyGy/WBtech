-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    id                 SERIAL PRIMARY KEY,
    order_uid          VARCHAR(255) NOT NULL,
    track_number       VARCHAR(255) NOT NULL,
    entry              VARCHAR(50)  NOT NULL,
    locale             VARCHAR(50),
    internal_signature VARCHAR(50),
    customer_id        VARCHAR(50),
    delivery_service   VARCHAR(255),
    shardkey           VARCHAR(50)  not null,
    sm_id              integer      not null,
    date_created       TIMESTAMP    not null default NOW(),
    oof_shard          VARCHAR(255),
    payments_id        integer      NOT NULL,
    delivery_id        integer      NOT NULL,
    CONSTRAINT fk_orders_payment_id
        FOREIGN KEY (payments_id)
            REFERENCES payments (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_delivery_id_id
        FOREIGN KEY (delivery_id)
            REFERENCES deliveries (id)
            ON DELETE CASCADE

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
