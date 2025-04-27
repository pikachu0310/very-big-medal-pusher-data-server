-- +goose Up
ALTER TABLE game_data
    ADD COLUMN max_chain_item   INT NOT NULL DEFAULT 0,
    ADD COLUMN max_chain_orange INT NOT NULL DEFAULT 0,
    ADD COLUMN max_chain_rainbow INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE game_data
    DROP COLUMN max_chain_item,
    DROP COLUMN max_chain_orange,
    DROP COLUMN max_chain_rainbow;
