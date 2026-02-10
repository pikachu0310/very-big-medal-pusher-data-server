-- +goose Up
-- Save Version 19 (Ferretta) columns and ranking cache fields

ALTER TABLE v2_save_data
    ADD COLUMN IF NOT EXISTS ferball_get INT NOT NULL DEFAULT 0 AFTER jacksp_totalmax,
    ADD COLUMN IF NOT EXISTS ferlot_lot INT NOT NULL DEFAULT 0 AFTER ferball_get,
    ADD COLUMN IF NOT EXISTS jackfr_get_all INT NOT NULL DEFAULT 0 AFTER ferlot_lot,
    ADD COLUMN IF NOT EXISTS jackfr_get_t0 INT NOT NULL DEFAULT 0 AFTER jackfr_get_all,
    ADD COLUMN IF NOT EXISTS jackfr_get_t1 INT NOT NULL DEFAULT 0 AFTER jackfr_get_t0,
    ADD COLUMN IF NOT EXISTS jackfr_get_t2 INT NOT NULL DEFAULT 0 AFTER jackfr_get_t1,
    ADD COLUMN IF NOT EXISTS jackfr_get_t3 INT NOT NULL DEFAULT 0 AFTER jackfr_get_t2,
    ADD COLUMN IF NOT EXISTS jackfr_get_t4 INT NOT NULL DEFAULT 0 AFTER jackfr_get_t3,
    ADD COLUMN IF NOT EXISTS jackfr_startmax BIGINT NOT NULL DEFAULT 0 AFTER jackfr_get_t4,
    ADD COLUMN IF NOT EXISTS jackfr_totalmax BIGINT NOT NULL DEFAULT 0 AFTER jackfr_startmax,
    ADD COLUMN IF NOT EXISTS ferlot_hit INT NOT NULL DEFAULT 0 AFTER jackfr_totalmax,
    ADD COLUMN IF NOT EXISTS ferlot_lose INT NOT NULL DEFAULT 0 AFTER ferlot_hit,
    ADD COLUMN IF NOT EXISTS ferlot_chance INT NOT NULL DEFAULT 0 AFTER ferlot_lose,
    ADD COLUMN IF NOT EXISTS ferlot_act INT NOT NULL DEFAULT 0 AFTER ferlot_chance,
    ADD COLUMN IF NOT EXISTS ferlot_lines INT NOT NULL DEFAULT 0 AFTER ferlot_act,
    ADD COLUMN IF NOT EXISTS bbox_shop INT NOT NULL DEFAULT 0 AFTER ferlot_lines;

CREATE TABLE IF NOT EXISTS v2_save_data_bbox_shop (
    save_id INT NOT NULL,
    item_id VARCHAR(255) NOT NULL,
    count INT NOT NULL,
    PRIMARY KEY (save_id, item_id),
    INDEX idx_v2_save_data_bbox_shop_save (save_id),
    FOREIGN KEY (save_id) REFERENCES v2_save_data (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS v2_save_data_ferlot_item (
    save_id INT NOT NULL,
    item_id VARCHAR(255) NOT NULL,
    count INT NOT NULL,
    PRIMARY KEY (save_id, item_id),
    INDEX idx_v2_save_data_ferlot_item_save (save_id),
    FOREIGN KEY (save_id) REFERENCES v2_save_data (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

ALTER TABLE v3_user_latest_save_data
    ADD COLUMN IF NOT EXISTS jackfr_startmax BIGINT NOT NULL DEFAULT 0 AFTER jacksp_startmax,
    ADD COLUMN IF NOT EXISTS jackfr_totalmax BIGINT NOT NULL DEFAULT 0 AFTER jackfr_startmax,
    ADD COLUMN IF NOT EXISTS ferlot_lines INT NOT NULL DEFAULT 0 AFTER jackfr_totalmax;

UPDATE v3_user_latest_save_data v3
JOIN v2_save_data v2 ON v3.save_id = v2.id
SET
    v3.jackfr_startmax = v2.jackfr_startmax,
    v3.jackfr_totalmax = v2.jackfr_totalmax,
    v3.ferlot_lines = v2.ferlot_lines;

CREATE INDEX idx_v3_user_latest_save_data_jackfr_startmax
    ON v3_user_latest_save_data (jackfr_startmax DESC);
CREATE INDEX idx_v3_user_latest_save_data_jackfr_totalmax
    ON v3_user_latest_save_data (jackfr_totalmax DESC);
CREATE INDEX idx_v3_user_latest_save_data_ferlot_lines
    ON v3_user_latest_save_data (ferlot_lines DESC);
CREATE INDEX idx_v3_latest_jackfr_start_hide
    ON v3_user_latest_save_data (hide_record, jackfr_startmax DESC, created_at ASC);
CREATE INDEX idx_v3_latest_jackfr_total_hide
    ON v3_user_latest_save_data (hide_record, jackfr_totalmax DESC, created_at ASC);
CREATE INDEX idx_v3_latest_ferlot_lines_hide
    ON v3_user_latest_save_data (hide_record, ferlot_lines DESC, created_at ASC);

DROP TRIGGER IF EXISTS update_v3_user_latest_save_data_after_insert;

-- +goose StatementBegin
CREATE TRIGGER update_v3_user_latest_save_data_after_insert
AFTER INSERT ON v2_save_data
FOR EACH ROW
BEGIN
    DECLARE v_achievements_count INT(11) DEFAULT 0;
    DECLARE v_max_chain_rainbow INT(11) DEFAULT 0;
    DECLARE v_golden_palball_get INT(11) DEFAULT 0;

    SELECT COUNT(*) INTO v_achievements_count
    FROM v2_save_data_achievements
    WHERE save_id = NEW.id;

    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain
    WHERE save_id = NEW.id AND ball_id = '3';

    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get
    WHERE save_id = NEW.id AND ball_id = '100';

    INSERT INTO v3_user_latest_save_data (
        user_id, save_id, version, credit_all, playtime, achievements_count,
        jacksp_startmax, jackfr_startmax, jackfr_totalmax, ferlot_lines, golden_palball_get,
        cpm_max, max_chain_rainbow, jack_totalmax_v2, ult_combomax, ult_totalmax_v2, blackbox_total, sp_use
    ) VALUES (
        NEW.user_id, NEW.id, NEW.version, NEW.credit_all, NEW.playtime, v_achievements_count,
        NEW.jacksp_startmax, NEW.jackfr_startmax, NEW.jackfr_totalmax, NEW.ferlot_lines, v_golden_palball_get,
        NEW.cpm_max, v_max_chain_rainbow, NEW.jack_totalmax_v2, NEW.ult_combomax, NEW.ult_totalmax_v2, NEW.blackbox_total, NEW.sp_use
    ) ON DUPLICATE KEY UPDATE
        version = NEW.version,
        credit_all = NEW.credit_all,
        playtime = NEW.playtime,
        save_id = NEW.id,
        achievements_count = v_achievements_count,
        jacksp_startmax = NEW.jacksp_startmax,
        jackfr_startmax = NEW.jackfr_startmax,
        jackfr_totalmax = NEW.jackfr_totalmax,
        ferlot_lines = NEW.ferlot_lines,
        golden_palball_get = v_golden_palball_get,
        cpm_max = NEW.cpm_max,
        max_chain_rainbow = v_max_chain_rainbow,
        jack_totalmax_v2 = NEW.jack_totalmax_v2,
        ult_combomax = NEW.ult_combomax,
        ult_totalmax_v2 = NEW.ult_totalmax_v2,
        blackbox_total = NEW.blackbox_total,
        sp_use = NEW.sp_use,
        updated_at = CURRENT_TIMESTAMP;
END;
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS update_v3_user_latest_save_data_after_insert;

-- 旧定義に戻す
-- +goose StatementBegin
CREATE TRIGGER update_v3_user_latest_save_data_after_insert
AFTER INSERT ON v2_save_data
FOR EACH ROW
BEGIN
    DECLARE v_achievements_count INT(11) DEFAULT 0;
    DECLARE v_max_chain_rainbow INT(11) DEFAULT 0;
    DECLARE v_golden_palball_get INT(11) DEFAULT 0;

    SELECT COUNT(*) INTO v_achievements_count
    FROM v2_save_data_achievements
    WHERE save_id = NEW.id;

    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain
    WHERE save_id = NEW.id AND ball_id = '3';

    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get
    WHERE save_id = NEW.id AND ball_id = '100';

    INSERT INTO v3_user_latest_save_data (
        user_id, save_id, version, credit_all, playtime, achievements_count, jacksp_startmax, golden_palball_get,
        cpm_max, max_chain_rainbow, jack_totalmax_v2, ult_combomax, ult_totalmax_v2, blackbox_total, sp_use
    ) VALUES (
        NEW.user_id, NEW.id, NEW.version, NEW.credit_all, NEW.playtime, v_achievements_count, NEW.jacksp_startmax, v_golden_palball_get,
        NEW.cpm_max, v_max_chain_rainbow, NEW.jack_totalmax_v2, NEW.ult_combomax, NEW.ult_totalmax_v2, NEW.blackbox_total, NEW.sp_use
    ) ON DUPLICATE KEY UPDATE
        version = NEW.version,
        credit_all = NEW.credit_all,
        playtime = NEW.playtime,
        save_id = NEW.id,
        achievements_count = v_achievements_count,
        jacksp_startmax = NEW.jacksp_startmax,
        golden_palball_get = v_golden_palball_get,
        cpm_max = NEW.cpm_max,
        max_chain_rainbow = v_max_chain_rainbow,
        jack_totalmax_v2 = NEW.jack_totalmax_v2,
        ult_combomax = NEW.ult_combomax,
        ult_totalmax_v2 = NEW.ult_totalmax_v2,
        blackbox_total = NEW.blackbox_total,
        sp_use = NEW.sp_use,
        updated_at = CURRENT_TIMESTAMP;
END;
-- +goose StatementEnd

DROP INDEX idx_v3_latest_ferlot_lines_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_jackfr_total_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_jackfr_start_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_user_latest_save_data_ferlot_lines ON v3_user_latest_save_data;
DROP INDEX idx_v3_user_latest_save_data_jackfr_totalmax ON v3_user_latest_save_data;
DROP INDEX idx_v3_user_latest_save_data_jackfr_startmax ON v3_user_latest_save_data;

ALTER TABLE v3_user_latest_save_data
    DROP COLUMN IF EXISTS ferlot_lines,
    DROP COLUMN IF EXISTS jackfr_totalmax,
    DROP COLUMN IF EXISTS jackfr_startmax;

DROP TABLE IF EXISTS v2_save_data_ferlot_item;
DROP TABLE IF EXISTS v2_save_data_bbox_shop;

ALTER TABLE v2_save_data
    DROP COLUMN IF EXISTS bbox_shop,
    DROP COLUMN IF EXISTS ferlot_lines,
    DROP COLUMN IF EXISTS ferlot_act,
    DROP COLUMN IF EXISTS ferlot_chance,
    DROP COLUMN IF EXISTS ferlot_lose,
    DROP COLUMN IF EXISTS ferlot_hit,
    DROP COLUMN IF EXISTS jackfr_totalmax,
    DROP COLUMN IF EXISTS jackfr_startmax,
    DROP COLUMN IF EXISTS jackfr_get_t4,
    DROP COLUMN IF EXISTS jackfr_get_t3,
    DROP COLUMN IF EXISTS jackfr_get_t2,
    DROP COLUMN IF EXISTS jackfr_get_t1,
    DROP COLUMN IF EXISTS jackfr_get_t0,
    DROP COLUMN IF EXISTS jackfr_get_all,
    DROP COLUMN IF EXISTS ferlot_lot,
    DROP COLUMN IF EXISTS ferball_get;
