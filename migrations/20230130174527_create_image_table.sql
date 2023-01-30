-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS image
(
    id BIGSERIAL PRIMARY KEY,
    created_at DATE DEFAULT now(),
    deleted_at DATE default NULL
);

-- +goose NO TRANSACTION
-- +goose Down
DROP TABLE image;
