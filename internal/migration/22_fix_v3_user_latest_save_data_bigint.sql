-- +goose Up
-- Fix v3_user_latest_save_data to support BIGINT values for jack_totalmax_v2 and ult_totalmax_v2
-- This fixes the "Out of range value" error when S-JP values exceed INT max (2,147,483,647)

ALTER TABLE v3_user_latest_save_data MODIFY COLUMN jack_totalmax_v2 BIGINT NOT NULL DEFAULT 0;
ALTER TABLE v3_user_latest_save_data MODIFY COLUMN ult_totalmax_v2 BIGINT NOT NULL DEFAULT 0;

-- +goose Down
-- Revert back to INT (not recommended as this will cause data loss for large values)

ALTER TABLE v3_user_latest_save_data MODIFY COLUMN jack_totalmax_v2 INT(11) NOT NULL DEFAULT 0;
ALTER TABLE v3_user_latest_save_data MODIFY COLUMN ult_totalmax_v2 INT(11) NOT NULL DEFAULT 0;

