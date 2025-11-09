-- +goose Up
ALTER TABLE v2_save_data
    ADD COLUMN pallot_lot_t4 INT NOT NULL DEFAULT 0 AFTER pallot_lot_t3,
    ADD COLUMN jacksp_get_t4 INT NOT NULL DEFAULT 0 AFTER jacksp_get_t3;

-- +goose Down
ALTER TABLE v2_save_data
    DROP COLUMN jacksp_get_t4,
    DROP COLUMN pallot_lot_t4;
