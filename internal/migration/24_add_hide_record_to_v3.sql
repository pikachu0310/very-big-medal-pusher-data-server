-- +goose Up
-- v3_user_latest_save_data テーブルに hide_record カラムを追加
ALTER TABLE v3_user_latest_save_data
    ADD COLUMN hide_record INT NOT NULL DEFAULT 0 AFTER sp_use;

-- 既存データの hide_record を v2_save_data から取得して更新
UPDATE v3_user_latest_save_data v3
JOIN v2_save_data v2 ON v3.save_id = v2.id
SET v3.hide_record = v2.hide_record;

-- hide_record を使用したランキング用インデックスを追加
-- hide_record = 0 のデータのみを効率的に取得するため
CREATE INDEX idx_v3_latest_hide_record ON v3_user_latest_save_data (hide_record);

-- 各ランキング用の複合インデックスを追加（hide_record を含む）
-- achievements_count ランキング用
CREATE INDEX idx_v3_latest_achievements_hide ON v3_user_latest_save_data (hide_record, achievements_count DESC, created_at ASC);

-- jacksp_startmax ランキング用
CREATE INDEX idx_v3_latest_jacksp_hide ON v3_user_latest_save_data (hide_record, jacksp_startmax DESC, created_at ASC);

-- golden_palball_get ランキング用
CREATE INDEX idx_v3_latest_golden_hide ON v3_user_latest_save_data (hide_record, golden_palball_get DESC, created_at ASC);

-- cpm_max ランキング用
CREATE INDEX idx_v3_latest_cpm_hide ON v3_user_latest_save_data (hide_record, cpm_max DESC, created_at ASC);

-- max_chain_rainbow ランキング用
CREATE INDEX idx_v3_latest_chain_hide ON v3_user_latest_save_data (hide_record, max_chain_rainbow DESC, created_at ASC);

-- jack_totalmax_v2 ランキング用
CREATE INDEX idx_v3_latest_jack_hide ON v3_user_latest_save_data (hide_record, jack_totalmax_v2 DESC, created_at ASC);

-- ult_combomax ランキング用
CREATE INDEX idx_v3_latest_ult_combo_hide ON v3_user_latest_save_data (hide_record, ult_combomax DESC, created_at ASC);

-- ult_totalmax_v2 ランキング用
CREATE INDEX idx_v3_latest_ult_total_hide ON v3_user_latest_save_data (hide_record, ult_totalmax_v2 DESC, created_at ASC);

-- sp_use ランキング用
CREATE INDEX idx_v3_latest_sp_hide ON v3_user_latest_save_data (hide_record, sp_use DESC, created_at ASC);

-- +goose Down
-- インデックスを削除
DROP INDEX idx_v3_latest_sp_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_ult_total_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_ult_combo_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_jack_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_chain_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_cpm_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_golden_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_jacksp_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_achievements_hide ON v3_user_latest_save_data;
DROP INDEX idx_v3_latest_hide_record ON v3_user_latest_save_data;

-- hide_record カラムを削除
ALTER TABLE v3_user_latest_save_data
    DROP COLUMN hide_record;

