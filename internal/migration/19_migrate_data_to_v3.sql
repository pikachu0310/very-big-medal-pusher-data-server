-- +goose Up
-- v2_save_data の内容を v3_user_latest_save_data に移行

INSERT INTO v3_user_latest_save_data (
    user_id, save_id, version, credit_all, playtime, achievements_count, jacksp_startmax, golden_palball_get,
    cpm_max, max_chain_rainbow, jack_totalmax_v2, ult_combomax, ult_totalmax_v2, sp_use
)
SELECT 
    latest.user_id,
    latest.save_id,
    latest.version,
    latest.credit_all,
    latest.playtime,
    COALESCE(ach_count.achievements_count, 0) as achievements_count,
    latest.jacksp_startmax,
    COALESCE(golden_palball.golden_palball_get, 0) as golden_palball_get,
    latest.cpm_max,
    COALESCE(rainbow_chain.max_chain_rainbow, 0) as max_chain_rainbow,
    latest.jack_totalmax_v2,
    latest.ult_combomax,
    latest.ult_totalmax_v2,
    latest.sp_use
FROM (
    -- 各ユーザーの最新セーブデータを取得
    SELECT 
        user_id,
        id as save_id,
        version,
        credit_all,
        playtime,
        jacksp_startmax,
        cpm_max,
        jack_totalmax_v2,
        ult_combomax,
        ult_totalmax_v2,
        sp_use
    FROM v2_save_data
    WHERE id IN (
        SELECT MAX(id) 
        FROM v2_save_data 
        GROUP BY user_id
    )
) latest
LEFT JOIN (
    -- 各ユーザーの最新セーブデータの実績数を計算
    SELECT 
        sd.user_id,
        COUNT(a.achievement_id) as achievements_count
    FROM v2_save_data sd
    JOIN v2_save_data_achievements a ON sd.id = a.save_id
    WHERE sd.id IN (
        SELECT MAX(id) 
        FROM v2_save_data 
        GROUP BY user_id
    )
    GROUP BY sd.user_id
) ach_count ON latest.user_id = ach_count.user_id
LEFT JOIN (
    -- 各ユーザーの最新セーブデータの最大レインボーチェイン数を計算
    SELECT 
        sd.user_id,
        MAX(bc.chain_count) as max_chain_rainbow
    FROM v2_save_data sd
    JOIN v2_save_data_ball_chain bc ON sd.id = bc.save_id
    WHERE sd.id IN (
        SELECT MAX(id) 
        FROM v2_save_data 
        GROUP BY user_id
    )
    AND bc.ball_id = '3'
    GROUP BY sd.user_id
       ) rainbow_chain ON latest.user_id = rainbow_chain.user_id
       LEFT JOIN (
           -- 各ユーザーの最新セーブデータのgolden_palball_getを計算
           SELECT 
               sd.user_id,
               COALESCE(SUM(pg.count), 0) as golden_palball_get
           FROM v2_save_data sd
           JOIN v2_save_data_palball_get pg ON sd.id = pg.save_id
           WHERE sd.id IN (
               SELECT MAX(id) 
               FROM v2_save_data 
               GROUP BY user_id
           )
           AND pg.ball_id = '100'
           GROUP BY sd.user_id
       ) golden_palball ON latest.user_id = golden_palball.user_id
ON DUPLICATE KEY UPDATE
    version = VALUES(version),
    credit_all = VALUES(credit_all),
    playtime = VALUES(playtime),
    save_id = VALUES(save_id),
    achievements_count = VALUES(achievements_count),
    jacksp_startmax = VALUES(jacksp_startmax),
    golden_palball_get = VALUES(golden_palball_get),
    cpm_max = VALUES(cpm_max),
    max_chain_rainbow = VALUES(max_chain_rainbow),
    jack_totalmax_v2 = VALUES(jack_totalmax_v2),
    ult_combomax = VALUES(ult_combomax),
    ult_totalmax_v2 = VALUES(ult_totalmax_v2),
    sp_use = VALUES(sp_use),
    updated_at = CURRENT_TIMESTAMP;

-- +goose Down
-- v3_user_latest_save_data のデータを削除

DELETE FROM v3_user_latest_save_data;
