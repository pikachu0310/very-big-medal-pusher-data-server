-- +goose Up
-- v3_user_latest_save_data_achievements テーブル作成
-- ユーザーごとの最新アチーブメントを保持し、重複を避ける

CREATE TABLE v3_user_latest_save_data_achievements (
    user_id VARCHAR(255) NOT NULL,
    achievement_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (user_id, achievement_id),
    
    -- インデックス
    INDEX idx_v3_user_latest_achievements_user_id (user_id),
    INDEX idx_v3_user_latest_achievements_achievement_id (achievement_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- v3_user_latest_save_data_achievements テーブルを削除

DROP TABLE IF EXISTS v3_user_latest_save_data_achievements;
