-- +goose Up
ALTER TABLE v2_save_data
    ADD COLUMN skill_point INT NOT NULL DEFAULT 0 AFTER buy_total,
    ADD COLUMN blackbox INT NOT NULL DEFAULT 0 AFTER skill_point,
    ADD COLUMN blackbox_total BIGINT NOT NULL DEFAULT 0 AFTER blackbox;

ALTER TABLE v3_user_latest_save_data
    ADD COLUMN blackbox_total BIGINT NOT NULL DEFAULT 0 AFTER ult_totalmax_v2;

-- 既存データを最新セーブデータから再集計
UPDATE v3_user_latest_save_data v3
    JOIN v2_save_data v2 ON v3.save_id = v2.id
SET v3.blackbox_total = v2.blackbox_total;

CREATE INDEX idx_v3_user_latest_save_data_blackbox_total
    ON v3_user_latest_save_data (blackbox_total DESC);
CREATE INDEX idx_v3_latest_blackbox_total_hide
    ON v3_user_latest_save_data (hide_record, blackbox_total DESC, created_at ASC);

DROP TRIGGER IF EXISTS update_v3_user_latest_save_data_after_insert;

-- +goose StatementBegin
CREATE TRIGGER update_v3_user_latest_save_data_after_insert
AFTER INSERT ON v2_save_data
FOR EACH ROW
BEGIN
    DECLARE v_achievements_count INT(11) DEFAULT 0;
    DECLARE v_max_chain_rainbow INT(11) DEFAULT 0;
    DECLARE v_golden_palball_get INT(11) DEFAULT 0;

    -- 実績数を計算
    SELECT COUNT(*) INTO v_achievements_count
    FROM v2_save_data_achievements
    WHERE save_id = NEW.id;

    -- 最大レインボーチェイン数を計算
    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain
    WHERE save_id = NEW.id AND ball_id = '3';

    -- golden_palball_getを計算（ball_id = 100の数）
    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get
    WHERE save_id = NEW.id AND ball_id = '100';

    -- v3_user_latest_save_data を更新
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

    -- 実績数を計算
    SELECT COUNT(*) INTO v_achievements_count
    FROM v2_save_data_achievements
    WHERE save_id = NEW.id;

    -- 最大レインボーチェイン数を計算
    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain
    WHERE save_id = NEW.id AND ball_id = '3';

    -- golden_palball_getを計算（ball_id = 100の数）
    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get
    WHERE save_id = NEW.id AND ball_id = '100';

    -- v3_user_latest_save_data を更新
    INSERT INTO v3_user_latest_save_data (
        user_id, save_id, version, credit_all, playtime, achievements_count, jacksp_startmax, golden_palball_get,
        cpm_max, max_chain_rainbow, jack_totalmax_v2, ult_combomax, ult_totalmax_v2, sp_use
    ) VALUES (
        NEW.user_id, NEW.id, NEW.version, NEW.credit_all, NEW.playtime, v_achievements_count, NEW.jacksp_startmax, v_golden_palball_get,
        NEW.cpm_max, v_max_chain_rainbow, NEW.jack_totalmax_v2, NEW.ult_combomax, NEW.ult_totalmax_v2, NEW.sp_use
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
        sp_use = NEW.sp_use,
        updated_at = CURRENT_TIMESTAMP;
END;
-- +goose StatementEnd

DROP INDEX idx_v3_latest_blackbox_total_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_user_latest_save_data_blackbox_total ON v3_user_latest_save_data;

ALTER TABLE v3_user_latest_save_data
    DROP COLUMN blackbox_total;

ALTER TABLE v2_save_data
    DROP COLUMN blackbox_total,
    DROP COLUMN blackbox,
    DROP COLUMN skill_point;
