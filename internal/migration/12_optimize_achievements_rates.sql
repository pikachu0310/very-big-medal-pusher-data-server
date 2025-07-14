-- +goose Up
-- Index for achievement_id to optimize GROUP BY queries
CREATE INDEX idx_save_data_v2_achievements_achievement_id ON save_data_v2_achievements (achievement_id);

-- Composite index for user_id and achievement_id for better JOIN performance
CREATE INDEX idx_save_data_v2_achievements_user_achievement ON save_data_v2_achievements (save_id, achievement_id);

-- +goose Down
DROP INDEX IF EXISTS idx_save_data_v2_achievements_user_achievement ON save_data_v2_achievements;
DROP INDEX IF EXISTS idx_save_data_v2_achievements_achievement_id ON save_data_v2_achievements;