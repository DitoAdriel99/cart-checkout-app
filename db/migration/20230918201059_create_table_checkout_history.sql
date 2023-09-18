-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS checkout_history
(
    id                  uuid        NOT NULL PRIMARY KEY,
    email               VARCHAR     NOT NULL,
    cart_id             uuid        NOT NULL,
    quantity            INT         NOT NULL,
    quantity_request    INT         NOT NULL,
    price               INT         NOT NULL,
    total_price         INT         NOT NULL,
    created_at          timestamp   not null,
    updated_at          timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS checkout_history
-- +goose StatementEnd
