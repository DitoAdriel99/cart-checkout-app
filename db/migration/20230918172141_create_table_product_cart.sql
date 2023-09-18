-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_cart
(
    id          uuid        not null primary key,
    email       varchar     not null,
    product_id  uuid        not null,
    quantity    int        not null,
    status      bool        not null,
    is_checkout bool        not null,
    created_at  timestamp   not null,
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_cart
-- +goose StatementEnd
