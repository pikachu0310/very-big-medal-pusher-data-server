-- +goose Up
-- 既存テーブルの照合順序を utf8mb4_unicode_ci に統一
-- これにより、v3系テーブルとの照合順序不一致エラーを解決

-- 段階的実行のため、各テーブルを個別に処理
-- 1. メインテーブルから開始
ALTER TABLE v2_save_data CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 2. 関連テーブル（外部キー制約のため順序重要）
ALTER TABLE v2_save_data_achievements CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_ball_chain CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_ball_get CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_medal_get CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_palball_get CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_palball_jp CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_perks CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE v2_save_data_perks_credit CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- +goose Down
-- 照合順序を元に戻す（通常は実行しない）
-- ALTER TABLE v2_save_data CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_achievements CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_ball_chain CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_ball_get CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_medal_get CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_palball_get CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_palball_jp CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_perks CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
-- ALTER TABLE v2_save_data_perks_credit CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_uca1400_ai_ci;
