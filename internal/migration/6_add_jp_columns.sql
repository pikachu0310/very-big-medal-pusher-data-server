-- +goose Up
ALTER TABLE game_data
    ADD COLUMN sugoroku_steps    INT NOT NULL DEFAULT 0,
    ADD COLUMN jackpots          INT NOT NULL DEFAULT 0,
    ADD COLUMN max_jackpot_win   INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE game_data
    DROP COLUMN sugoroku_steps,
    DROP COLUMN jackpots,
    DROP COLUMN max_jackpot_win;
