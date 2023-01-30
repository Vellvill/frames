-- +goose Up
CREATE TABLE IF NOT EXISTS good
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT CHECK (length(name) < 100),
    created_at DATE DEFAULT now(),
    deleted_at DATE default NULL
);

create unique index on good(name);

-- +goose Down
DROP TABLE good;

