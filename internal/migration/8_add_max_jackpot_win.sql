-- +goose Up
ALTER TABLE game_data
    ADD COLUMN max_total_jackpot  INT NOT NULL DEFAULT 0,
    ADD COLUMN max_total_ultimate INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE game_data
    DROP COLUMN max_total_jackpot,
    DROP COLUMN max_total_ultimate;
