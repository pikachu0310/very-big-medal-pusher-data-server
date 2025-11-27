-- +goose Up
-- Extend counters to BIGINT for long-compatible fields
-- Targets: sqr_get, jack_get, bstp_step, bstp_rwd, sp_use

ALTER TABLE v2_save_data
    MODIFY COLUMN sqr_get BIGINT NOT NULL DEFAULT 0,
    MODIFY COLUMN jack_get BIGINT NOT NULL DEFAULT 0,
    MODIFY COLUMN bstp_step BIGINT NOT NULL DEFAULT 0,
    MODIFY COLUMN bstp_rwd BIGINT NOT NULL DEFAULT 0,
    MODIFY COLUMN sp_use BIGINT NOT NULL DEFAULT 0;

ALTER TABLE v3_user_latest_save_data
    MODIFY COLUMN sp_use BIGINT NOT NULL DEFAULT 0;

-- +goose Down
-- Revert fields back to INT (may truncate large values)

ALTER TABLE v3_user_latest_save_data
    MODIFY COLUMN sp_use INT NOT NULL DEFAULT 0;

ALTER TABLE v2_save_data
    MODIFY COLUMN sqr_get INT NOT NULL DEFAULT 0,
    MODIFY COLUMN jack_get INT NOT NULL DEFAULT 0,
    MODIFY COLUMN bstp_step INT NOT NULL DEFAULT 0,
    MODIFY COLUMN bstp_rwd INT NOT NULL DEFAULT 0,
    MODIFY COLUMN sp_use INT NOT NULL DEFAULT 0;
