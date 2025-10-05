-- +goose Up
-- v3_user_latest_save_data_achievements に既存データを移行
-- v3_user_latest_save_dataから各ユーザーの最新save_idを取得してアチーブメントを移行

INSERT INTO v3_user_latest_save_data_achievements (
    user_id, achievement_id
)
SELECT DISTINCT
    v3.user_id,
    a.achievement_id
FROM v3_user_latest_save_data v3
JOIN v2_save_data_achievements a ON v3.save_id = a.save_id
ON DUPLICATE KEY UPDATE
    created_at = CURRENT_TIMESTAMP;

-- +goose Down
-- v3_user_latest_save_data_achievements のデータを削除

DELETE FROM v3_user_latest_save_data_achievements;
