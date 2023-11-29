-- +goose Up
-- goose postgres "host=localhost user=admin database=db_WbTech0 password=root sslmode=disable" up
-- +goose StatementBegin
CREATE TABLE deliveries
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(250) NOT NULL,
    phone   VARCHAR(30)  NOT NULL,
    zip     VARCHAR(250) NOT NULL,
    city    VARCHAR(250) NOT NULL,
    address VARCHAR(250) NOT NULL,
    region  VARCHAR(250) NOT NULL,
    email   VARCHAR(250) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS deliveries;
-- +goose StatementEnd
