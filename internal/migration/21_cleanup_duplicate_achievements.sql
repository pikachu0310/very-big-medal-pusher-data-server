-- +goose Up
-- v2_save_data_achievements から重複データを削除
-- 各ユーザー・アチーブメントの組み合わせで最初の記録のみを残す

-- 一時テーブルを作成して、各ユーザー・アチーブメントの最初の記録を特定
CREATE TEMPORARY TABLE temp_first_achievements AS
SELECT 
    user_id,
    achievement_id,
    MIN(save_id) as first_save_id
FROM v2_save_data_achievements a
JOIN v2_save_data sd ON a.save_id = sd.id
GROUP BY user_id, achievement_id;

-- 重複データを削除（最初の記録以外）
DELETE a FROM v2_save_data_achievements a
JOIN v2_save_data sd ON a.save_id = sd.id
LEFT JOIN temp_first_achievements tfa ON sd.user_id = tfa.user_id 
    AND a.achievement_id = tfa.achievement_id 
    AND a.save_id = tfa.first_save_id
WHERE tfa.first_save_id IS NULL;

-- 一時テーブルを削除
DROP TEMPORARY TABLE temp_first_achievements;

-- +goose Down
-- このマイグレーションは元に戻せません（データ削除のため）
-- 重複データの削除は不可逆的な操作です
