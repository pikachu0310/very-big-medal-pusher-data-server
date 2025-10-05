-- +goose Up
-- v3_user_latest_save_data_achievements に既存データを移行

INSERT INTO v3_user_latest_save_data_achievements (
    user_id, achievement_id, first_achieved_at
)
SELECT 
    sd.user_id,
    a.achievement_id,
    sd.created_at as first_achieved_at
FROM v2_save_data_achievements a
JOIN v2_save_data sd ON a.save_id = sd.id
ON DUPLICATE KEY UPDATE
    updated_at = CURRENT_TIMESTAMP;

-- +goose Down
-- v3_user_latest_save_data_achievements のデータを削除

DELETE FROM v3_user_latest_save_data_achievements;
