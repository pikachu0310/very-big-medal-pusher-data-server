-- +goose Up
-- Fix type mismatches for Save v9 and v10 columns
-- Change jack_totalmax_v2 and ult_totalmax_v2 from DOUBLE to INT
-- Note: Using MySQL-specific syntax for goose compatibility

-- Drop and recreate columns with correct types
ALTER TABLE save_data_v2 DROP COLUMN jack_totalmax_v2;
ALTER TABLE save_data_v2 ADD COLUMN jack_totalmax_v2 INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 DROP COLUMN ult_totalmax_v2;
ALTER TABLE save_data_v2 ADD COLUMN ult_totalmax_v2 INT NOT NULL DEFAULT 0;

-- Change jacksp_get_* columns from DOUBLE to INT
ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_all;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_all INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t0;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t0 INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t1;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t1 INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t2;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t2 INT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t3;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t3 INT NOT NULL DEFAULT 0;

-- Change jacksp_startmax and jacksp_totalmax from DOUBLE to BIGINT
ALTER TABLE save_data_v2 DROP COLUMN jacksp_startmax;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_startmax BIGINT NOT NULL DEFAULT 0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_totalmax;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_totalmax BIGINT NOT NULL DEFAULT 0;

-- Change save_data_v2_perks_credit.credits from INT to BIGINT
ALTER TABLE save_data_v2_perks_credit DROP COLUMN credits;
ALTER TABLE save_data_v2_perks_credit ADD COLUMN credits BIGINT NOT NULL DEFAULT 0;

-- +goose Down
-- Revert type changes back to original types
ALTER TABLE save_data_v2 DROP COLUMN jack_totalmax_v2;
ALTER TABLE save_data_v2 ADD COLUMN jack_totalmax_v2 DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN ult_totalmax_v2;
ALTER TABLE save_data_v2 ADD COLUMN ult_totalmax_v2 DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_all;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_all DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t0;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t0 DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t1;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t1 DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t2;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t2 DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_get_t3;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_get_t3 DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_startmax;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_startmax DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2 DROP COLUMN jacksp_totalmax;
ALTER TABLE save_data_v2 ADD COLUMN jacksp_totalmax DOUBLE NOT NULL DEFAULT 0.0;

ALTER TABLE save_data_v2_perks_credit DROP COLUMN credits;
ALTER TABLE save_data_v2_perks_credit ADD COLUMN credits INT NOT NULL DEFAULT 0;
