-- +goose Up
-- Update statistics fields to BIGINT to support larger values
-- These fields are used in v3 endpoints and need to support Long type values
-- Using MODIFY COLUMN to preserve existing data

-- Update medal_get from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN medal_get BIGINT NOT NULL DEFAULT 0;

-- Update ball_get from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN ball_get BIGINT NOT NULL DEFAULT 0;

-- Update slot_start from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN slot_start BIGINT NOT NULL DEFAULT 0;

-- Update slot_startfev from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN slot_startfev BIGINT NOT NULL DEFAULT 0;

-- Update slot_hit from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN slot_hit BIGINT NOT NULL DEFAULT 0;

-- Update slot_getfev from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN slot_getfev BIGINT NOT NULL DEFAULT 0;

-- Update sqr_step from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN sqr_step BIGINT NOT NULL DEFAULT 0;

-- Update jack_startmax from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN jack_startmax BIGINT NOT NULL DEFAULT 0;

-- Update jack_totalmax_v2 from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN jack_totalmax_v2 BIGINT NOT NULL DEFAULT 0;

-- Update ult_totalmax_v2 from INT to BIGINT
ALTER TABLE save_data_v2 MODIFY COLUMN ult_totalmax_v2 BIGINT NOT NULL DEFAULT 0;

-- +goose Down
-- Revert changes back to INT
-- Using MODIFY COLUMN to preserve existing data

ALTER TABLE save_data_v2 MODIFY COLUMN medal_get INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN ball_get INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN slot_start INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN slot_startfev INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN slot_hit INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN slot_getfev INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN sqr_step INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN jack_startmax INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN jack_totalmax_v2 INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 MODIFY COLUMN ult_totalmax_v2 INT NOT NULL DEFAULT 0;
