-- +goose Up
ALTER TABLE v2_save_data
    /* Save Version 12 - Totem System */
    ADD COLUMN IF NOT EXISTS totem_altars INT NOT NULL DEFAULT 0 AFTER task_cnt,
    ADD COLUMN IF NOT EXISTS totem_altars_credit BIGINT NOT NULL DEFAULT 0 AFTER totem_altars;

-- トーテムレベルテーブルを作成
CREATE TABLE IF NOT EXISTS v2_save_data_totems
(
    save_id     INT NOT NULL,
    totem_id    INT NOT NULL,
    level       INT NOT NULL,
    PRIMARY KEY (save_id, totem_id),
    INDEX idx_v2_save_data_totems_save (save_id),
    FOREIGN KEY (save_id) REFERENCES v2_save_data (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- トーテムごとの消費クレジットテーブルを作成
CREATE TABLE IF NOT EXISTS v2_save_data_totems_credit
(
    save_id     INT    NOT NULL,
    totem_id    INT    NOT NULL,
    credits     BIGINT NOT NULL,
    PRIMARY KEY (save_id, totem_id),
    INDEX idx_v2_save_data_totems_credit_save (save_id),
    FOREIGN KEY (save_id) REFERENCES v2_save_data (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- セット中のトーテムIDリストテーブルを作成
CREATE TABLE IF NOT EXISTS v2_save_data_totems_placement
(
    save_id       INT NOT NULL,
    placement_idx INT NOT NULL,
    totem_id      INT NOT NULL,
    PRIMARY KEY (save_id, placement_idx),
    INDEX idx_v2_save_data_totems_placement_save (save_id),
    FOREIGN KEY (save_id) REFERENCES v2_save_data (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS v2_save_data_totems_placement;
DROP TABLE IF EXISTS v2_save_data_totems_credit;
DROP TABLE IF EXISTS v2_save_data_totems;

ALTER TABLE v2_save_data
    DROP COLUMN IF EXISTS totem_altars_credit,
    DROP COLUMN IF EXISTS totem_altars;

