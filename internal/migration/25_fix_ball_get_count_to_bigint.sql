-- +goose Up
-- Fix v2_save_data_ball_get.count to BIGINT to prevent overflow
-- This fixes the issue where ball_get count exceeds INT max (2,147,483,647)

ALTER TABLE v2_save_data_ball_get MODIFY COLUMN count BIGINT NOT NULL;

-- +goose Down
-- Revert back to INT (not recommended as this will cause data loss for large values)

ALTER TABLE v2_save_data_ball_get MODIFY COLUMN count INT(11) NOT NULL;

