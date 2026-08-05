-- +goose Up
-- Add fraud data disable flag

ALTER TABLE v2_save_data
    ADD COLUMN IF NOT EXISTS disabled TINYINT NOT NULL DEFAULT 0 AFTER hide_record;

-- +goose Down
ALTER TABLE v2_save_data
    DROP COLUMN IF EXISTS disabled;
