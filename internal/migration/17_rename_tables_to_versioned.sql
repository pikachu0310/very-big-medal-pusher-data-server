-- +goose Up
-- 既存テーブルをバージョン付き名前にリネーム

-- v1 テーブル
RENAME TABLE game_data TO v1_game_data;

-- v2 テーブル
RENAME TABLE save_data_v2 TO v2_save_data;
RENAME TABLE save_data_v2_achievements TO v2_save_data_achievements;
RENAME TABLE save_data_v2_ball_chain TO v2_save_data_ball_chain;
RENAME TABLE save_data_v2_ball_get TO v2_save_data_ball_get;
RENAME TABLE save_data_v2_medal_get TO v2_save_data_medal_get;
RENAME TABLE save_data_v2_palball_get TO v2_save_data_palball_get;
RENAME TABLE save_data_v2_palball_jp TO v2_save_data_palball_jp;
RENAME TABLE save_data_v2_perks TO v2_save_data_perks;
RENAME TABLE save_data_v2_perks_credit TO v2_save_data_perks_credit;

-- +goose Down
-- テーブル名を元に戻す

-- v2 テーブル
RENAME TABLE v2_save_data_perks_credit TO save_data_v2_perks_credit;
RENAME TABLE v2_save_data_perks TO save_data_v2_perks;
RENAME TABLE v2_save_data_palball_jp TO save_data_v2_palball_jp;
RENAME TABLE v2_save_data_palball_get TO save_data_v2_palball_get;
RENAME TABLE v2_save_data_medal_get TO save_data_v2_medal_get;
RENAME TABLE v2_save_data_ball_get TO save_data_v2_ball_get;
RENAME TABLE v2_save_data_ball_chain TO save_data_v2_ball_chain;
RENAME TABLE v2_save_data_achievements TO save_data_v2_achievements;
RENAME TABLE v2_save_data TO save_data_v2;

-- v1 テーブル
RENAME TABLE v1_game_data TO game_data;
