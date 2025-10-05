-- +goose Up
-- v3_user_latest_save_data テーブル作成
-- ユーザーごとの最新セーブデータを保持し、統計計算を高速化する

CREATE TABLE v3_user_latest_save_data (
    user_id VARCHAR(255) PRIMARY KEY,
    save_id INT(11) NOT NULL,
    
    -- 基本データ（v2_save_dataから取得）
    version INT(11) NOT NULL DEFAULT 0,
    credit_all BIGINT(20) NOT NULL DEFAULT 0,
    playtime BIGINT(20) NOT NULL DEFAULT 0,
    
    -- v4/statistics で使用する前計算済み値
    achievements_count INT(11) NOT NULL DEFAULT 0,
    jacksp_startmax BIGINT(20) NOT NULL DEFAULT 0,
    golden_palball_get INT(11) NOT NULL DEFAULT 0,
    cpm_max DOUBLE NOT NULL DEFAULT 0,
    max_chain_rainbow INT(11) NOT NULL DEFAULT 0,
    jack_totalmax_v2 INT(11) NOT NULL DEFAULT 0,
    ult_combomax INT(11) NOT NULL DEFAULT 0,
    ult_totalmax_v2 INT(11) NOT NULL DEFAULT 0,
    sp_use INT(11) NOT NULL DEFAULT 0,
    
    -- メタデータ
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- 統計計算用インデックス
    INDEX idx_v3_user_latest_save_data_credit_all (credit_all DESC),
    INDEX idx_v3_user_latest_save_data_achievements_count (achievements_count DESC),
    INDEX idx_v3_user_latest_save_data_jacksp_startmax (jacksp_startmax DESC),
    INDEX idx_v3_user_latest_save_data_golden_palball_get (golden_palball_get DESC),
    INDEX idx_v3_user_latest_save_data_cpm_max (cpm_max DESC),
    INDEX idx_v3_user_latest_save_data_max_chain_rainbow (max_chain_rainbow DESC),
    INDEX idx_v3_user_latest_save_data_jack_totalmax_v2 (jack_totalmax_v2 DESC),
    INDEX idx_v3_user_latest_save_data_ult_combomax (ult_combomax DESC),
    INDEX idx_v3_user_latest_save_data_ult_totalmax_v2 (ult_totalmax_v2 DESC),
    INDEX idx_v3_user_latest_save_data_sp_use (sp_use DESC),
    INDEX idx_v3_user_latest_save_data_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- v3_user_latest_save_data テーブルを削除

DROP TABLE IF EXISTS v3_user_latest_save_data;