-- +goose Up
ALTER TABLE game_data
    ADD COLUMN maxChainItem   INT NOT NULL DEFAULT 0,
    ADD COLUMN maxChainOrange INT NOT NULL DEFAULT 0,
    ADD COLUMN maxChainRainbow INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE game_data
    DROP COLUMN maxChainItem,
    DROP COLUMN maxChainOrange,
    DROP COLUMN maxChainRainbow;
