-- +goose Up

ALTER TABLE feeds ADD COLUMN list_fetched_at timestamp;

-- +goose Down

ALTER TABLE users DROP COLUMN list_fetched_at;