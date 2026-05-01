-- +goose Up
-- Fix: Error 1020 (HY000) + retry-409 race condition in InsertSaveV4

DROP TRIGGER IF EXISTS update_v3_user_latest_save_data_after_insert;
DROP TRIGGER IF EXISTS update_v3_achievements_count_after_insert;
DROP TRIGGER IF EXISTS update_v3_max_chain_rainbow_after_insert;
DROP TRIGGER IF EXISTS update_v3_golden_palball_get_after_insert;

ALTER TABLE v2_save_data
    DROP INDEX idx_save_data_v2_user_playtime,
    ADD UNIQUE INDEX uniq_save_data_v2_user_playtime (user_id, playtime);

-- +goose Down
ALTER TABLE v2_save_data
    DROP INDEX uniq_save_data_v2_user_playtime,
    ADD INDEX idx_save_data_v2_user_playtime (user_id, playtime);

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

-- +goose StatementBegin
CREATE TRIGGER update_v3_achievements_count_after_insert
AFTER INSERT ON v2_save_data_achievements
FOR EACH ROW
BEGIN
    DECLARE v_user_id VARCHAR(255);
    DECLARE v_achievements_count INT(11) DEFAULT 0;
    SELECT user_id INTO v_user_id FROM v2_save_data WHERE id = NEW.save_id;
    SELECT COUNT(*) INTO v_achievements_count FROM v2_save_data_achievements WHERE save_id = NEW.save_id;
    UPDATE v3_user_latest_save_data
    SET achievements_count = v_achievements_count, updated_at = CURRENT_TIMESTAMP
    WHERE user_id = v_user_id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_v3_max_chain_rainbow_after_insert
AFTER INSERT ON v2_save_data_ball_chain
FOR EACH ROW
BEGIN
    DECLARE v_user_id VARCHAR(255);
    DECLARE v_max_chain_rainbow INT(11) DEFAULT 0;
    SELECT user_id INTO v_user_id FROM v2_save_data WHERE id = NEW.save_id;
    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain WHERE save_id = NEW.save_id AND ball_id = '3';
    UPDATE v3_user_latest_save_data
    SET max_chain_rainbow = v_max_chain_rainbow, updated_at = CURRENT_TIMESTAMP
    WHERE user_id = v_user_id;
END;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_v3_golden_palball_get_after_insert
AFTER INSERT ON v2_save_data_palball_get
FOR EACH ROW
BEGIN
    DECLARE v_user_id VARCHAR(255);
    DECLARE v_golden_palball_get INT(11) DEFAULT 0;
    SELECT user_id INTO v_user_id FROM v2_save_data WHERE id = NEW.save_id;
    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get WHERE save_id = NEW.save_id AND ball_id = '100';
    UPDATE v3_user_latest_save_data
    SET golden_palball_get = v_golden_palball_get, updated_at = CURRENT_TIMESTAMP
    WHERE user_id = v_user_id;
END;
-- +goose StatementEnd
