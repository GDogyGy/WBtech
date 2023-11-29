-- +goose Up
-- +goose StatementBegin
CREATE TABLE products
(
    id           SERIAL PRIMARY KEY,
    chrt_id      int          NOT NULL,
    track_number VARCHAR(250) NOT NULL,
    price        numeric      NOT NULL default 0,
    rid          VARCHAR(250),
    name         VARCHAR(250) NOT NULL,
    sale         numeric      NOT NULL default 0,
    size         VARCHAR(250),
    total_price  numeric      NOT NULL default 0, -- wtf тут поле ошибочно по логике, уточнить
    nm_id        int          NOT NULL,
    brand        VARCHAR(250),
    status       int,
    order_id     int          NOT NULL,
    check ( price > sale ),
    CONSTRAINT products_order_id
        FOREIGN KEY (order_id)
            REFERENCES orders (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
