-- +goose Up
-- Save Version 21 (Ferretta usage) columns and dictionary table

ALTER TABLE v2_save_data
    ADD COLUMN IF NOT EXISTS ferlot_maxln INT NOT NULL DEFAULT 0 AFTER bbox_shop,
    ADD COLUMN IF NOT EXISTS bbox_used_ferlot INT NOT NULL DEFAULT 0 AFTER ferlot_maxln;

CREATE TABLE IF NOT EXISTS v2_save_data_ferlot_useitem (
    save_id INT NOT NULL,
    item_id VARCHAR(255) NOT NULL,
    count INT NOT NULL,
    PRIMARY KEY (save_id, item_id),
    INDEX idx_v2_save_data_ferlot_useitem_save (save_id),
    FOREIGN KEY (save_id) REFERENCES v2_save_data (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS v2_save_data_ferlot_useitem;

ALTER TABLE v2_save_data
    DROP COLUMN IF EXISTS bbox_used_ferlot,
    DROP COLUMN IF EXISTS ferlot_maxln;
