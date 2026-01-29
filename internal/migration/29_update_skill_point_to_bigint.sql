-- +goose Up
-- Extend skill_point to BIGINT for long-compatible values

ALTER TABLE v2_save_data
    MODIFY COLUMN skill_point BIGINT NOT NULL DEFAULT 0;

-- +goose Down
-- Revert skill_point back to INT (may truncate large values)

ALTER TABLE v2_save_data
    MODIFY COLUMN skill_point INT NOT NULL DEFAULT 0;
