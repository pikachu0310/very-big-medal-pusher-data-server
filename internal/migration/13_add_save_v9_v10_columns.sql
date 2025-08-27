-- +goose Up
ALTER TABLE save_data_v2
    /* Save Version 9 */
    ADD COLUMN hide_record INT NOT NULL DEFAULT 0 AFTER sp_use,
    ADD COLUMN cpm_max DOUBLE NOT NULL DEFAULT 0.0 AFTER hide_record,
    /* Save Version 10 */
    ADD COLUMN jack_totalmax_v2 DOUBLE NOT NULL DEFAULT 0.0 AFTER cpm_max,
    ADD COLUMN ult_totalmax_v2 DOUBLE NOT NULL DEFAULT 0.0 AFTER jack_totalmax_v2,
    ADD COLUMN palball_get INT NOT NULL DEFAULT 0 AFTER ult_totalmax_v2,
    ADD COLUMN pallot_lot_t0 INT NOT NULL DEFAULT 0 AFTER palball_get,
    ADD COLUMN pallot_lot_t1 INT NOT NULL DEFAULT 0 AFTER pallot_lot_t0,
    ADD COLUMN pallot_lot_t2 INT NOT NULL DEFAULT 0 AFTER pallot_lot_t1,
    ADD COLUMN pallot_lot_t3 INT NOT NULL DEFAULT 0 AFTER pallot_lot_t2,
    ADD COLUMN jacksp_get_all DOUBLE NOT NULL DEFAULT 0.0 AFTER pallot_lot_t3,
    ADD COLUMN jacksp_get_t0 DOUBLE NOT NULL DEFAULT 0.0 AFTER jacksp_get_all,
    ADD COLUMN jacksp_get_t1 DOUBLE NOT NULL DEFAULT 0.0 AFTER jacksp_get_t0,
    ADD COLUMN jacksp_get_t2 DOUBLE NOT NULL DEFAULT 0.0 AFTER jacksp_get_t1,
    ADD COLUMN jacksp_get_t3 DOUBLE NOT NULL DEFAULT 0.0 AFTER jacksp_get_t2,
    ADD COLUMN jacksp_startmax DOUBLE NOT NULL DEFAULT 0.0 AFTER jacksp_get_t3,
    ADD COLUMN jacksp_totalmax DOUBLE NOT NULL DEFAULT 0.0 AFTER jacksp_startmax,
    ADD COLUMN task_cnt INT NOT NULL DEFAULT 0 AFTER jacksp_totalmax;

-- 新しい子テーブルを作成
CREATE TABLE save_data_v2_palball_get
(
    save_id  INT          NOT NULL,
    ball_id  VARCHAR(255) NOT NULL,
    count    INT          NOT NULL,
    PRIMARY KEY (save_id, ball_id),
    INDEX idx_save_data_v2_palball_get_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE save_data_v2_palball_jp
(
    save_id  INT          NOT NULL,
    ball_id  VARCHAR(255) NOT NULL,
    count    INT          NOT NULL,
    PRIMARY KEY (save_id, ball_id),
    INDEX idx_save_data_v2_palball_jp_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE save_data_v2_perks
(
    save_id INT NOT NULL,
    perk_id INT NOT NULL,
    level   INT NOT NULL,
    PRIMARY KEY (save_id, perk_id),
    INDEX idx_save_data_v2_perks_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE save_data_v2_perks_credit
(
    save_id INT NOT NULL,
    perk_id INT NOT NULL,
    credits INT NOT NULL,
    PRIMARY KEY (save_id, perk_id),
    INDEX idx_save_data_v2_perks_credit_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS save_data_v2_perks_credit;
DROP TABLE IF EXISTS save_data_v2_perks;
DROP TABLE IF EXISTS save_data_v2_palball_jp;
DROP TABLE IF EXISTS save_data_v2_palball_get;

ALTER TABLE save_data_v2
    DROP COLUMN task_cnt,
    DROP COLUMN jacksp_totalmax,
    DROP COLUMN jacksp_startmax,
    DROP COLUMN jacksp_get_t3,
    DROP COLUMN jacksp_get_t2,
    DROP COLUMN jacksp_get_t1,
    DROP COLUMN jacksp_get_t0,
    DROP COLUMN jacksp_get_all,
    DROP COLUMN pallot_lot_t3,
    DROP COLUMN pallot_lot_t2,
    DROP COLUMN pallot_lot_t1,
    DROP COLUMN pallot_lot_t0,
    DROP COLUMN palball_get,
    DROP COLUMN ult_totalmax_v2,
    DROP COLUMN jack_totalmax_v2,
    DROP COLUMN cpm_max,
    DROP COLUMN hide_record;
