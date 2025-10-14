package repository

import (
	"context"
	"fmt"
	"runtime"

	"github.com/jmoiron/sqlx"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// InsertSaveV4 persists a SaveData and its child tables, and updates v3_user_latest_save_data
func (r *Repository) InsertSaveV4(ctx context.Context, sd *domain.SaveData) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// v2_save_data に挿入
	res, err := tx.ExecContext(ctx, `
INSERT INTO v2_save_data (
    user_id, legacy, version,
    credit, credit_all, medal_in, medal_get,
    ball_get, ball_chain, slot_start, slot_startfev,
    slot_hit, slot_getfev, sqr_get, sqr_step,
    jack_get, jack_startmax, jack_totalmax,
    ult_get, ult_combomax, ult_totalmax,
    rmshbi_get, bstp_step, bstp_rwd, buy_total, sp_use,
    hide_record, cpm_max, jack_totalmax_v2, ult_totalmax_v2,
    palball_get, pallot_lot_t0, pallot_lot_t1, pallot_lot_t2, pallot_lot_t3,
    jacksp_get_all, jacksp_get_t0, jacksp_get_t1, jacksp_get_t2, jacksp_get_t3,
    jacksp_startmax, jacksp_totalmax, task_cnt, totem_altars, totem_altars_credit, buy_shbi,
    firstboot, lastsave, playtime
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		sd.UserId, sd.Legacy, sd.Version,
		sd.Credit, sd.CreditAll, sd.MedalIn, sd.MedalGet,
		sd.BallGet, sd.BallChain, sd.SlotStart, sd.SlotStartFev,
		sd.SlotHit, sd.SlotGetFev, sd.SqrGet, sd.SqrStep,
		sd.JackGet, sd.JackStartMax, sd.JackTotalMax,
		sd.UltGet, sd.UltComboMax, sd.UltTotalMax,
		sd.RmShbiGet, sd.BstpStep, sd.BstpRwd, sd.BuyTotal, sd.SpUse,
		sd.HideRecord, sd.CpMMax, sd.JackTotalMaxV2, sd.UltimateTotalMaxV2,
		sd.PalettaBallGet, sd.PalettaLotteryAttemptTier0, sd.PalettaLotteryAttemptTier1, sd.PalettaLotteryAttemptTier2, sd.PalettaLotteryAttemptTier3,
		sd.JackpotSuperGetTotal, sd.JackpotSuperGetTier0, sd.JackpotSuperGetTier1, sd.JackpotSuperGetTier2, sd.JackpotSuperGetTier3,
		sd.JackpotSuperStartMax, sd.JackpotSuperTotalMax, sd.TaskCompleteCount, sd.TotemAltarUnlockCount, sd.TotemAltarUnlockUsedCredits, sd.BuyShbi,
		sd.FirstBoot, sd.LastSave, sd.Playtime,
	)
	if err != nil {
		return err
	}
	saveID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// 関連テーブルに挿入
	// medal_get
	for id, cnt := range sd.DCMedalGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_medal_get(save_id, medal_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}
	// ball_get
	for id, cnt := range sd.DCBallGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_ball_get(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}
	// ball_chain
	for id, cnt := range sd.DCBallChain {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_ball_chain(save_id, ball_id, chain_count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}
	// achievements - 最適化版：新しいアチーブメントのみを追加
	if err := r.insertNewAchievements(ctx, tx, sd.UserId, saveID, sd.LAchieve); err != nil {
		return err
	}

	// palball_get
	for id, cnt := range sd.DCPalettaBallGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_palball_get(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}

	// palball_jp
	for id, cnt := range sd.DCPalettaBallJackpot {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_palball_jp(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}

	// perks
	for i, level := range sd.LPerkLevels {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_perks(save_id, perk_id, level) VALUES(?,?,?)`, saveID, i, level); err != nil {
			return err
		}
	}

	// perks_credit
	for i, credits := range sd.LPerkUsedCredits {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_perks_credit(save_id, perk_id, credits) VALUES(?,?,?)`, saveID, i, credits); err != nil {
			return err
		}
	}

	// totems
	for i, level := range sd.LTotemLevels {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_totems(save_id, totem_id, level) VALUES(?,?,?)`, saveID, i, level); err != nil {
			return err
		}
	}

	// totems_credit
	for i, credits := range sd.LTotemUsedCredits {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_totems_credit(save_id, totem_id, credits) VALUES(?,?,?)`, saveID, i, credits); err != nil {
			return err
		}
	}

	// totems_placement
	for i, totemID := range sd.LTotemPlacements {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_totems_placement(save_id, placement_idx, totem_id) VALUES(?,?,?)`, saveID, i, totemID); err != nil {
			return err
		}
	}

	// v3_user_latest_save_data を更新
	// 実績数を計算
	achievementsCount := len(sd.LAchieve)

	// 最大レインボーチェイン数を計算
	maxChainRainbow := 0
	if chainCount, exists := sd.DCBallChain["3"]; exists {
		maxChainRainbow = chainCount
	}

	// golden_palball_getを計算（ball_id = 100の数）
	goldenPalballGet := 0
	if palballCount, exists := sd.DCPalettaBallGet["100"]; exists {
		goldenPalballGet = palballCount
	}

	// v3_user_latest_save_data に挿入/更新
	_, err = tx.ExecContext(ctx, `
INSERT INTO v3_user_latest_save_data (
    user_id, version, credit_all, playtime, save_id, achievements_count, jacksp_startmax, golden_palball_get,
    cpm_max, max_chain_rainbow, jack_totalmax_v2, ult_combomax, ult_totalmax_v2, sp_use
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)
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
    updated_at = CURRENT_TIMESTAMP`,
		sd.UserId, sd.Version, sd.CreditAll, sd.Playtime, saveID, achievementsCount, sd.JackpotSuperStartMax, goldenPalballGet,
		sd.CpMMax, maxChainRainbow, sd.JackTotalMaxV2, sd.UltComboMax, sd.UltimateTotalMaxV2, sd.SpUse,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetStatisticsV4 returns the latest statistics for V4 using v3_user_latest_save_data (ランキング上限 1000).
func (r *Repository) GetStatisticsV4(ctx context.Context) (*models.StatisticsV4, error) {
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 START\n")
	stats := &models.StatisticsV4{}

	// ------ 共通ヘルパ (匿名関数) -----------------------------------------
	addRanking := func(ptr **[]models.RankingEntry, query string, args ...interface{}) error {
		rows, err := r.db.QueryxContext(ctx, query, args...)
		if err != nil {
			return err
		}
		defer rows.Close()

		// 事前に容量を確保してメモリ効率を改善（最大1000個）
		list := make([]models.RankingEntry, 0, 1000)
		for rows.Next() {
			var e models.RankingEntry
			if err := rows.Scan(&e.UserId, &e.Value, &e.CreatedAt); err != nil {
				return err
			}
			list = append(list, e)
		}
		*ptr = &list
		return nil
	}

	// 1) achievements_count
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_1 - achievements_count\n")
	if err := addRanking(&stats.AchievementsCount, `
SELECT
  user_id,
  achievements_count AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY achievements_count DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 2) jacksp_startmax
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_2 - jacksp_startmax\n")
	if err := addRanking(&stats.JackspStartmax, `
SELECT
  user_id,
  jacksp_startmax AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY jacksp_startmax DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 3) golden_palball_get
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_3 - golden_palball_get\n")
	if err := addRanking(&stats.GoldenPalballGet, `
SELECT
  user_id,
  golden_palball_get AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY golden_palball_get DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 4) cpm_max
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_4 - cpm_max\n")
	if err := addRanking(&stats.CpmMax, `
SELECT
  user_id,
  CAST(cpm_max AS SIGNED) AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY cpm_max DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 5) max_chain_rainbow
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_5 - max_chain_rainbow\n")
	if err := addRanking(&stats.MaxChainRainbow, `
SELECT
  user_id,
  max_chain_rainbow AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY max_chain_rainbow DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 6) jack_totalmax_v2
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_6 - jack_totalmax_v2\n")
	if err := addRanking(&stats.JackTotalmaxV2, `
SELECT
  user_id,
  jack_totalmax_v2 AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY jack_totalmax_v2 DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 7) ult_combomax
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_7 - ult_combomax\n")
	if err := addRanking(&stats.UltCombomax, `
SELECT
  user_id,
  ult_combomax AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY ult_combomax DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 8) ult_totalmax_v2
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_8 - ult_totalmax_v2\n")
	if err := addRanking(&stats.UltTotalmaxV2, `
SELECT
  user_id,
  ult_totalmax_v2 AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY ult_totalmax_v2 DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 9) sp_use
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CREATING_RANKING_9 - sp_use\n")
	if err := addRanking(&stats.SpUse, `
SELECT
  user_id,
  sp_use AS value,
  created_at
FROM v3_user_latest_save_data
ORDER BY sp_use DESC, created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 10) total_medals（全ユーザーの合計）
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 CALCULATING_TOTAL_MEDALS\n")
	var totalMedals int
	{
		q := `
SELECT
  COALESCE(SUM(credit_all), 0) AS total_medals
FROM v3_user_latest_save_data
`
		if err := r.db.GetContext(ctx, &totalMedals, q); err != nil {
			return nil, err
		}
		stats.TotalMedals = &totalMedals
	}

	// 最終段階でGCを促してメモリ使用量を安定化
	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 FINAL_GC_TRIGGER\n")
	runtime.GC()

	fmt.Printf("[REPO-DEBUG] GetStatisticsV4 SUCCESS - total_medals=%d\n", totalMedals)
	return stats, nil
}

// insertNewAchievements は新しいアチーブメントのみを挿入する最適化版
func (r *Repository) insertNewAchievements(ctx context.Context, tx *sqlx.Tx, userID string, saveID int64, newAchievements []string) error {
	if len(newAchievements) == 0 {
		return nil
	}

	// 既存のアチーブメントを取得
	existingAchievements := make(map[string]bool)
	rows, err := tx.QueryContext(ctx, `
		SELECT achievement_id 
		FROM v3_user_latest_save_data_achievements 
		WHERE user_id = ?
	`, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var achievementID string
		if err := rows.Scan(&achievementID); err != nil {
			return err
		}
		existingAchievements[achievementID] = true
	}

	// 新しいアチーブメントのみを処理
	for _, achievementID := range newAchievements {
		if !existingAchievements[achievementID] {
			// v2_save_data_achievements に新しいアチーブメントを追加（ログ用）
			if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_achievements(save_id, achievement_id) VALUES(?,?)`, saveID, achievementID); err != nil {
				return err
			}

			// v3_user_latest_save_data_achievements に追加
			if _, err := tx.ExecContext(ctx, `
				INSERT INTO v3_user_latest_save_data_achievements (user_id, achievement_id) 
				VALUES (?,?)
			`, userID, achievementID); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetAchievementRates returns achievement acquisition rates
func (r *Repository) GetAchievementRates(ctx context.Context) (*models.AchievementRates, error) {
	// 最適化1: 全ユーザー数を先に取得（サブクエリを排除）
	var totalUsers int
	err := r.db.GetContext(ctx, &totalUsers, `
SELECT COUNT(DISTINCT user_id) FROM v3_user_latest_save_data_achievements
`)
	if err != nil {
		return nil, err
	}

	// 最適化2: アチーブメント別ユーザー数のみを取得（重複なしなのでCOUNT(DISTINCT)は不要）
	rows, err := r.db.QueryxContext(ctx, `
SELECT 
    achievement_id,
    COUNT(user_id) as user_count
FROM v3_user_latest_save_data_achievements
GROUP BY achievement_id
ORDER BY user_count DESC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 最適化3: 事前に容量を確保（アチーブメント数は限定的）
	achievementRates := make(map[string]struct {
		Count *int     `json:"count,omitempty"`
		Rate  *float32 `json:"rate,omitempty"`
	}, 1000)

	for rows.Next() {
		var achievementID string
		var userCount int
		if err := rows.Scan(&achievementID, &userCount); err != nil {
			return nil, err
		}

		// 最適化4: 浮動小数点計算を最適化
		rate := float32(0.0)
		if totalUsers > 0 {
			rate = float32(userCount) / float32(totalUsers)
		}

		achievementRates[achievementID] = struct {
			Count *int     `json:"count,omitempty"`
			Rate  *float32 `json:"rate,omitempty"`
		}{
			Count: &userCount,
			Rate:  &rate,
		}
	}

	return &models.AchievementRates{
		TotalUsers:       &totalUsers,
		AchievementRates: &achievementRates,
	}, nil
}
