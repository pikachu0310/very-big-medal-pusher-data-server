-- +goose Up
-- Save Version 22 medal tower counter

ALTER TABLE v2_save_data
    ADD COLUMN IF NOT EXISTS get_medaltower INT NOT NULL DEFAULT 0 AFTER bbox_used_ferlot;

-- +goose Down
ALTER TABLE v2_save_data
    DROP COLUMN IF EXISTS get_medaltower;
