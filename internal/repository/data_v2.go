package repository

import (
	"context"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func (r *Repository) ExistsSameSave(ctx context.Context, userID string, playtime int64) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM save_data_v2 WHERE user_id=? AND playtime=?)`
	var exists int
	err := r.db.QueryRowContext(ctx, q, userID, playtime).Scan(&exists)
	return exists == 1, err
}

// InsertSave persists a SaveData and its child tables
func (r *Repository) InsertSave(ctx context.Context, sd *domain.SaveData) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, `
INSERT INTO save_data_v2 (
    user_id, legacy, version,
    credit, credit_all, medal_in, medal_get,
    ball_get, ball_chain, slot_start, slot_startfev,
    slot_hit, slot_getfev, sqr_get, sqr_step,
    jack_get, jack_startmax, jack_totalmax,
    ult_get, ult_combomax, ult_totalmax,
    rmshbi_get, bstp_step, bstp_rwd, buy_total, sp_use, buy_shbi,
    hide_record, cpm_max, jack_totalmax_v2, ult_totalmax_v2,
    palball_get, pallot_lot_t0, pallot_lot_t1, pallot_lot_t2, pallot_lot_t3,
    jacksp_get_all, jacksp_get_t0, jacksp_get_t1, jacksp_get_t2, jacksp_get_t3,
    jacksp_startmax, jacksp_totalmax, task_cnt,
    firstboot, lastsave, playtime
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		sd.UserId, sd.Legacy, sd.Version,
		sd.Credit, sd.CreditAll, sd.MedalIn, sd.MedalGet,
		sd.BallGet, sd.BallChain, sd.SlotStart, sd.SlotStartFev,
		sd.SlotHit, sd.SlotGetFev, sd.SqrGet, sd.SqrStep,
		sd.JackGet, sd.JackStartMax, sd.JackTotalMax,
		sd.UltGet, sd.UltComboMax, sd.UltTotalMax,
		sd.RmShbiGet, sd.BstpStep, sd.BstpRwd, sd.BuyTotal, sd.SpUse, sd.BuyShbi,
		sd.HideRecord, sd.CpMMax, sd.JackTotalMaxV2, sd.UltimateTotalMaxV2,
		sd.PalettaBallGet, sd.PalettaLotteryAttemptTier0, sd.PalettaLotteryAttemptTier1, sd.PalettaLotteryAttemptTier2, sd.PalettaLotteryAttemptTier3,
		sd.JackpotSuperGetTotal, sd.JackpotSuperGetTier0, sd.JackpotSuperGetTier1, sd.JackpotSuperGetTier2, sd.JackpotSuperGetTier3,
		sd.JackpotSuperStartMax, sd.JackpotSuperTotalMax, sd.TaskCompleteCount,
		sd.FirstBoot, sd.LastSave, sd.Playtime,
	)
	if err != nil {
		return err
	}
	saveID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// medal_get
	for id, cnt := range sd.DCMedalGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_medal_get(save_id, medal_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}
	// ball_get
	for id, cnt := range sd.DCBallGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_ball_get(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}
	// ball_chain
	for id, cnt := range sd.DCBallChain {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_ball_chain(save_id, ball_id, chain_count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}
	// achievements
	for _, aid := range sd.LAchieve {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_achievements(save_id, achievement_id) VALUES(?,?)`, saveID, aid); err != nil {
			return err
		}
	}

	// palball_get
	for id, cnt := range sd.DCPalettaBallGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_palball_get(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}

	// palball_jp
	for id, cnt := range sd.DCPalettaBallJackpot {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_palball_jp(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			return err
		}
	}

	// perks
	for i, level := range sd.LPerkLevels {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_perks(save_id, perk_id, level) VALUES(?,?,?)`, saveID, i, level); err != nil {
			return err
		}
	}

	// perks_credit
	for i, credits := range sd.LPerkUsedCredits {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_perks_credit(save_id, perk_id, credits) VALUES(?,?,?)`, saveID, i, credits); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetLatestSave retrieves the latest SaveData for a user
func (r *Repository) GetLatestSave(ctx context.Context, userID string) (*domain.SaveData, error) {
	var sd domain.SaveData
	// 1) main row
	err := r.db.GetContext(ctx, &sd, `
SELECT * 
FROM save_data_v2 
WHERE user_id = ? 
ORDER BY created_at DESC 
LIMIT 1
`, userID)
	if err != nil {
		return nil, err
	}

	// 2) medal_get map
	sd.DCMedalGet = make(map[string]int)
	rows, err := r.db.QueryxContext(ctx, `
SELECT medal_id, count 
FROM save_data_v2_medal_get 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var cnt int
		if err := rows.Scan(&id, &cnt); err != nil {
			return nil, err
		}
		sd.DCMedalGet[id] = cnt
	}
	rows.Close()

	// 3) ball_get
	sd.DCBallGet = make(map[string]int)
	rows, err = r.db.QueryxContext(ctx, `
SELECT ball_id, count 
FROM save_data_v2_ball_get 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var cnt int
		if err := rows.Scan(&id, &cnt); err != nil {
			return nil, err
		}
		sd.DCBallGet[id] = cnt
	}
	rows.Close()

	// 4) ball_chain
	sd.DCBallChain = make(map[string]int)
	rows, err = r.db.QueryxContext(ctx, `
SELECT ball_id, chain_count 
FROM save_data_v2_ball_chain 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var cnt int
		if err := rows.Scan(&id, &cnt); err != nil {
			return nil, err
		}
		sd.DCBallChain[id] = cnt
	}
	rows.Close()

	// 5) achievements
	rows, err = r.db.QueryxContext(ctx, `
SELECT achievement_id 
FROM save_data_v2_achievements 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aid string
		if err := rows.Scan(&aid); err != nil {
			return nil, err
		}
		sd.LAchieve = append(sd.LAchieve, aid)
	}
	rows.Close()

	// 6) palball_get
	sd.DCPalettaBallGet = make(map[string]int)
	rows, err = r.db.QueryxContext(ctx, `
SELECT ball_id, count 
FROM save_data_v2_palball_get 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var cnt int
		if err := rows.Scan(&id, &cnt); err != nil {
			return nil, err
		}
		sd.DCPalettaBallGet[id] = cnt
	}
	rows.Close()

	// 7) palball_jp
	sd.DCPalettaBallJackpot = make(map[string]int)
	rows, err = r.db.QueryxContext(ctx, `
SELECT ball_id, count 
FROM save_data_v2_palball_jp 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var cnt int
		if err := rows.Scan(&id, &cnt); err != nil {
			return nil, err
		}
		sd.DCPalettaBallJackpot[id] = cnt
	}
	rows.Close()

	// 8) perks
	rows, err = r.db.QueryxContext(ctx, `
SELECT perk_id, level 
FROM save_data_v2_perks 
WHERE save_id = ?
ORDER BY perk_id
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var perkID int
		var level int
		if err := rows.Scan(&perkID, &level); err != nil {
			return nil, err
		}
		// perk_idの順序に合わせて配列を拡張
		for len(sd.LPerkLevels) <= perkID {
			sd.LPerkLevels = append(sd.LPerkLevels, 0)
		}
		sd.LPerkLevels[perkID] = level
	}
	rows.Close()

	// 9) perks_credit
	rows, err = r.db.QueryxContext(ctx, `
SELECT perk_id, credits 
FROM save_data_v2_perks_credit 
WHERE save_id = ?
ORDER BY perk_id
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var perkID int
		var credits int64
		if err := rows.Scan(&perkID, &credits); err != nil {
			return nil, err
		}
		// perk_idの順序に合わせて配列を拡張
		for len(sd.LPerkUsedCredits) <= perkID {
			sd.LPerkUsedCredits = append(sd.LPerkUsedCredits, 0)
		}
		sd.LPerkUsedCredits[perkID] = credits
	}
	rows.Close()

	return &sd, nil
}

// GetStatistics returns combined rankings and total medals.
// internal/repository/data_v2.go
func (r *Repository) GetStatistics(ctx context.Context) (*models.StatisticsV2, error) {
	stats := &models.StatisticsV2{}

	// 1) max_chain_orange (ball_id = '1')
	{
		q := `
SELECT
  sd.user_id,
  MAX(bc.chain_count) AS value,
  MIN(sd.created_at)  AS created_at
FROM save_data_v2_ball_chain AS bc
JOIN save_data_v2           AS sd ON bc.save_id = sd.id
WHERE bc.ball_id = '1'
GROUP BY sd.user_id
ORDER BY value DESC, created_at ASC
LIMIT 500
`
		rows, err := r.db.QueryxContext(ctx, q)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		list := []models.RankingEntry{}
		for rows.Next() {
			var e models.RankingEntry
			if err := rows.Scan(&e.UserId, &e.Value, &e.CreatedAt); err != nil {
				return nil, err
			}
			list = append(list, e)
		}
		stats.MaxChainOrange = &list
	}

	// 2) max_chain_rainbow (ball_id = '3')
	{
		q := `
SELECT
  sd.user_id,
  MAX(bc.chain_count) AS value,
  MIN(sd.created_at)  AS created_at
FROM save_data_v2_ball_chain AS bc
JOIN save_data_v2           AS sd ON bc.save_id = sd.id
WHERE bc.ball_id = '3'
GROUP BY sd.user_id
ORDER BY value DESC, created_at ASC
LIMIT 500
`
		rows, err := r.db.QueryxContext(ctx, q)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		list := []models.RankingEntry{}
		for rows.Next() {
			var e models.RankingEntry
			if err := rows.Scan(&e.UserId, &e.Value, &e.CreatedAt); err != nil {
				return nil, err
			}
			list = append(list, e)
		}
		stats.MaxChainRainbow = &list
	}

	// 3) max_total_jackpot (以前のまま)
	{
		q := `
SELECT
  sd.user_id,
  MAX(sd.jack_totalmax) AS value,
  MIN(sd.created_at)    AS created_at
FROM save_data_v2 AS sd
GROUP BY sd.user_id
ORDER BY value DESC, created_at ASC
LIMIT 500
`
		rows, err := r.db.QueryxContext(ctx, q)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		list := []models.RankingEntry{}
		for rows.Next() {
			var e models.RankingEntry
			if err := rows.Scan(&e.UserId, &e.Value, &e.CreatedAt); err != nil {
				return nil, err
			}
			list = append(list, e)
		}
		stats.MaxTotalJackpot = &list
	}

	// 4) total_medals (以前のまま)
	{
		q := `
SELECT
  COALESCE(SUM(sd.credit),0) AS total_medals
FROM (
  SELECT user_id, MAX(id) AS max_id
  FROM save_data_v2
  GROUP BY user_id
) AS t
JOIN save_data_v2 AS sd
  ON sd.user_id = t.user_id
 AND sd.id      = t.max_id
`
		var total int
		if err := r.db.GetContext(ctx, &total, q); err != nil {
			return nil, err
		}
		stats.TotalMedals = &total
	}

	return stats, nil
}

// GetStatisticsV3 returns the latest statistics for V3 (ランキング上限 1000).
func (r *Repository) GetStatisticsV3(ctx context.Context) (*models.StatisticsV3, error) {
	stats := &models.StatisticsV3{}

	// ------ 共通ヘルパ (匿名関数) -----------------------------------------
	addRanking := func(ptr **[]models.RankingEntry, query string, args ...interface{}) error {
		rows, err := r.db.QueryxContext(ctx, query, args...)
		if err != nil {
			return err
		}
		defer rows.Close()

		list := []models.RankingEntry{}
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

	// 1) max_chain_orange (ball_id = '1')
	if err := addRanking(&stats.MaxChainOrange, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    sd.user_id,
    bc.chain_count AS value,
    sd.created_at,
    ROW_NUMBER() OVER (PARTITION BY sd.user_id ORDER BY bc.chain_count DESC, sd.created_at ASC) AS rn
  FROM save_data_v2_ball_chain AS bc
  JOIN save_data_v2 AS sd ON bc.save_id = sd.id
  WHERE bc.ball_id = '1' AND sd.hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 2) max_chain_rainbow (ball_id = '3')
	if err := addRanking(&stats.MaxChainRainbow, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    sd.user_id,
    bc.chain_count AS value,
    sd.created_at,
    ROW_NUMBER() OVER (PARTITION BY sd.user_id ORDER BY bc.chain_count DESC, sd.created_at ASC) AS rn
  FROM save_data_v2_ball_chain AS bc
  JOIN save_data_v2 AS sd ON bc.save_id = sd.id
  WHERE bc.ball_id = '3' AND sd.hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 3) jack_startmax
	if err := addRanking(&stats.JackStartmax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    jack_startmax AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jack_startmax DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 4) jack_totalmax
	if err := addRanking(&stats.JackTotalmax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    jack_totalmax AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jack_totalmax DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 5) jack_totalmax_v2
	if err := addRanking(&stats.JackTotalmaxV2, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    jack_totalmax_v2 AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jack_totalmax_v2 DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 6) ult_totalmax_v2
	if err := addRanking(&stats.UltTotalmaxV2, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    ult_totalmax_v2 AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY ult_totalmax_v2 DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 7) jacksp_startmax
	if err := addRanking(&stats.JackspStartmax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    jacksp_startmax AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jacksp_startmax DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 8) golden_palball_get (ID 100)
	if err := addRanking(&stats.GoldenPalballGet, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    sd.user_id,
    COALESCE(dpg.count, 0) AS value,
    sd.created_at,
    ROW_NUMBER() OVER (PARTITION BY sd.user_id ORDER BY COALESCE(dpg.count, 0) DESC, sd.created_at ASC) AS rn
  FROM save_data_v2 AS sd
  LEFT JOIN save_data_v2_palball_get AS dpg ON dpg.save_id = sd.id AND dpg.ball_id = '100'
  WHERE sd.hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 9) cpm_max
	if err := addRanking(&stats.CpmMax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    cpm_max AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY cpm_max DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 10) ult_combomax
	if err := addRanking(&stats.UltCombomax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    ult_combomax AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY ult_combomax DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 11) ult_totalmax
	if err := addRanking(&stats.UltTotalmax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    ult_totalmax AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY ult_totalmax DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 12) sp_use
	if err := addRanking(&stats.SpUse, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    sp_use AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY sp_use DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 13) buy_shbi
	if err := addRanking(&stats.BuyShbi, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    buy_shbi AS value,
    created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY buy_shbi DESC, created_at ASC) AS rn
  FROM save_data_v2
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 14) achievements_count
	if err := addRanking(&stats.AchievementsCount, `
SELECT
  sd.user_id,
  COUNT(a.achievement_id) AS value,
  sd.created_at
FROM (
  SELECT user_id, MAX(id) AS max_id
  FROM save_data_v2
  WHERE hide_record = false
  GROUP BY user_id
) AS latest
JOIN save_data_v2 AS sd ON sd.id = latest.max_id
LEFT JOIN save_data_v2_achievements AS a ON a.save_id = sd.id
GROUP BY sd.user_id, sd.created_at
ORDER BY value DESC, sd.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 15) total_medals（最新セーブの credit 合計）
	{
		q := `
SELECT
  COALESCE(SUM(sd.credit),0) AS total_medals
FROM (
  SELECT user_id, MAX(id) AS max_id
  FROM save_data_v2
  WHERE hide_record = false
  GROUP BY user_id
) AS t
JOIN save_data_v2 AS sd
  ON sd.user_id = t.user_id
 AND sd.id      = t.max_id
`
		var total int
		if err := r.db.GetContext(ctx, &total, q); err != nil {
			return nil, err
		}
		stats.TotalMedals = &total
	}

	return stats, nil
}

// GetAchievementRates returns achievement acquisition rates
func (r *Repository) GetAchievementRates(ctx context.Context) (*models.AchievementRates, error) {
	// 最適化されたクエリ：1回のクエリで総ユーザー数と各実績の取得者数を取得
	rows, err := r.db.QueryxContext(ctx, `
WITH users_with_achievements AS (
    SELECT DISTINCT sd.user_id
    FROM save_data_v2_achievements a
    JOIN save_data_v2 sd ON a.save_id = sd.id
)
SELECT 
    achievement_id,
    COUNT(DISTINCT sd.user_id) as user_count,
    (SELECT COUNT(*) FROM users_with_achievements) as total_users
FROM save_data_v2_achievements a
JOIN save_data_v2 sd ON a.save_id = sd.id
GROUP BY achievement_id
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	achievementRates := make(map[string]struct {
		Count *int     `json:"count,omitempty"`
		Rate  *float32 `json:"rate,omitempty"`
	})

	var totalUsers int
	for rows.Next() {
		var achievementID string
		var userCount int
		if err := rows.Scan(&achievementID, &userCount, &totalUsers); err != nil {
			return nil, err
		}

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
