-- +goose Up
ALTER TABLE game_data
    ADD INDEX idx_game_data_user_created_at  (user_id, created_at),
    ADD INDEX idx_game_data_user_total_play_time (user_id, total_play_time);

-- +goose Down
ALTER TABLE game_data
    DROP INDEX idx_game_data_user_total_play_time,
    DROP INDEX idx_game_data_user_created_at;
