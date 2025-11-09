package repository

import (
	"context"
	"fmt"
	"runtime"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func (r *Repository) ExistsSameSave(ctx context.Context, userID string, playtime int64) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM v2_save_data WHERE user_id=? AND playtime=?)`
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
INSERT INTO v2_save_data (
    user_id, legacy, version,
    credit, credit_all, medal_in, medal_get,
    ball_get, ball_chain, slot_start, slot_startfev,
    slot_hit, slot_getfev, sqr_get, sqr_step,
    jack_get, jack_startmax, jack_totalmax,
    ult_get, ult_combomax, ult_totalmax,
    rmshbi_get, bstp_step, bstp_rwd, buy_total, skill_point, blackbox, blackbox_total, sp_use,
    hide_record, cpm_max, jack_totalmax_v2, ult_totalmax_v2,
    palball_get, pallot_lot_t0, pallot_lot_t1, pallot_lot_t2, pallot_lot_t3, pallot_lot_t4,
    jacksp_get_all, jacksp_get_t0, jacksp_get_t1, jacksp_get_t2, jacksp_get_t3, jacksp_get_t4,
    jacksp_startmax, jacksp_totalmax, task_cnt, totem_altars, totem_altars_credit, buy_shbi,
    firstboot, lastsave, playtime
	) VALUES (
	    ?, ?, ?,
	    ?, ?, ?, ?,
	    ?, ?, ?, ?,
	    ?, ?, ?, ?,
	    ?, ?, ?,
	    ?, ?, ?,
	    ?, ?, ?, ?, ?, ?, ?, ?,
	    ?, ?, ?, ?,
	    ?, ?, ?, ?, ?, ?,
	    ?, ?, ?, ?, ?, ?,
	    ?, ?, ?, ?, ?, ?,
	    ?, ?, ?
	)`,
		sd.UserId, sd.Legacy, sd.Version,
		sd.Credit, sd.CreditAll, sd.MedalIn, sd.MedalGet,
		sd.BallGet, sd.BallChain, sd.SlotStart, sd.SlotStartFev,
		sd.SlotHit, sd.SlotGetFev, sd.SqrGet, sd.SqrStep,
		sd.JackGet, sd.JackStartMax, sd.JackTotalMax,
		sd.UltGet, sd.UltComboMax, sd.UltTotalMax,
		sd.RmShbiGet, sd.BstpStep, sd.BstpRwd, sd.BuyTotal, sd.SkillPoint, sd.BlackBox, sd.BlackBoxTotal, sd.SpUse,
		sd.HideRecord, sd.CpMMax, sd.JackTotalMaxV2, sd.UltimateTotalMaxV2,
		sd.PalettaBallGet, sd.PalettaLotteryAttemptTier0, sd.PalettaLotteryAttemptTier1, sd.PalettaLotteryAttemptTier2, sd.PalettaLotteryAttemptTier3, sd.PalettaLotteryAttemptTier4,
		sd.JackpotSuperGetTotal, sd.JackpotSuperGetTier0, sd.JackpotSuperGetTier1, sd.JackpotSuperGetTier2, sd.JackpotSuperGetTier3, sd.JackpotSuperGetTier4,
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
	// achievements
	for _, aid := range sd.LAchieve {
		if _, err := tx.ExecContext(ctx, `INSERT INTO v2_save_data_achievements(save_id, achievement_id) VALUES(?,?)`, saveID, aid); err != nil {
			return err
		}
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

	return tx.Commit()
}

// GetLatestSave retrieves the latest SaveData for a user
func (r *Repository) GetLatestSave(ctx context.Context, userID string) (*domain.SaveData, error) {
	fmt.Printf("[REPO-DEBUG] GetLatestSave START - user_id=%s\n", userID)

	var sd domain.SaveData
	// 1) main row
	fmt.Printf("[REPO-DEBUG] GetLatestSave FETCHING_MAIN_ROW - user_id=%s\n", userID)
	err := r.db.GetContext(ctx, &sd, `
SELECT * 
FROM v2_save_data 
WHERE user_id = ? 
ORDER BY updated_at DESC 
LIMIT 1
`, userID)
	if err != nil {
		fmt.Printf("[REPO-DEBUG] GetLatestSave MAIN_ROW_ERROR - user_id=%s, error=%v\n", userID, err)
		return nil, err
	}

	// 2) medal_get map
	sd.DCMedalGet = make(map[string]int)
	rows, err := r.db.QueryxContext(ctx, `
SELECT medal_id, count 
FROM v2_save_data_medal_get 
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
	sd.DCBallGet = make(map[string]int64)
	rows, err = r.db.QueryxContext(ctx, `
SELECT ball_id, count 
FROM v2_save_data_ball_get 
WHERE save_id = ?
`, sd.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var cnt int64
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
FROM v2_save_data_ball_chain 
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

	// 5) achievements - v3_user_latest_save_data_achievements から取得
	rows, err = r.db.QueryxContext(ctx, `
SELECT achievement_id 
FROM v3_user_latest_save_data_achievements 
WHERE user_id = ?
`, userID)
	if err != nil {
		return nil, err
	}
	// 事前に容量を確保してメモリ効率を改善（最大1000個まで）
	sd.LAchieve = make([]string, 0, 1000)
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
FROM v2_save_data_palball_get 
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
FROM v2_save_data_palball_jp 
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
FROM v2_save_data_perks 
WHERE save_id = ?
ORDER BY perk_id
`, sd.ID)
	if err != nil {
		return nil, err
	}
	// 事前に容量を確保してメモリ効率を改善（最大100個まで）
	sd.LPerkLevels = make([]int, 0, 100)
	for rows.Next() {
		var perkID int
		var level int
		if err := rows.Scan(&perkID, &level); err != nil {
			return nil, err
		}
		// perk_idの順序に合わせて配列を拡張（制限付き）
		for len(sd.LPerkLevels) <= perkID && len(sd.LPerkLevels) < 100 {
			sd.LPerkLevels = append(sd.LPerkLevels, 0)
		}
		if perkID < len(sd.LPerkLevels) {
			sd.LPerkLevels[perkID] = level
		}
	}
	rows.Close()

	// 9) perks_credit
	rows, err = r.db.QueryxContext(ctx, `
SELECT perk_id, credits 
FROM v2_save_data_perks_credit 
WHERE save_id = ?
ORDER BY perk_id
`, sd.ID)
	if err != nil {
		return nil, err
	}
	// 事前に容量を確保してメモリ効率を改善（最大100個まで）
	sd.LPerkUsedCredits = make([]int64, 0, 100)
	for rows.Next() {
		var perkID int
		var credits int64
		if err := rows.Scan(&perkID, &credits); err != nil {
			return nil, err
		}
		// perk_idの順序に合わせて配列を拡張（制限付き）
		for len(sd.LPerkUsedCredits) <= perkID && len(sd.LPerkUsedCredits) < 100 {
			sd.LPerkUsedCredits = append(sd.LPerkUsedCredits, 0)
		}
		if perkID < len(sd.LPerkUsedCredits) {
			sd.LPerkUsedCredits[perkID] = credits
		}
	}
	rows.Close()

	// 10) totems
	rows, err = r.db.QueryxContext(ctx, `
SELECT totem_id, level 
FROM v2_save_data_totems 
WHERE save_id = ?
ORDER BY totem_id
`, sd.ID)
	if err != nil {
		return nil, err
	}
	sd.LTotemLevels = make([]int, 0, 100)
	for rows.Next() {
		var totemID int
		var level int
		if err := rows.Scan(&totemID, &level); err != nil {
			return nil, err
		}
		for len(sd.LTotemLevels) <= totemID && len(sd.LTotemLevels) < 100 {
			sd.LTotemLevels = append(sd.LTotemLevels, 0)
		}
		if totemID < len(sd.LTotemLevels) {
			sd.LTotemLevels[totemID] = level
		}
	}
	rows.Close()

	// 11) totems_credit
	rows, err = r.db.QueryxContext(ctx, `
SELECT totem_id, credits 
FROM v2_save_data_totems_credit 
WHERE save_id = ?
ORDER BY totem_id
`, sd.ID)
	if err != nil {
		return nil, err
	}
	sd.LTotemUsedCredits = make([]int64, 0, 100)
	for rows.Next() {
		var totemID int
		var credits int64
		if err := rows.Scan(&totemID, &credits); err != nil {
			return nil, err
		}
		for len(sd.LTotemUsedCredits) <= totemID && len(sd.LTotemUsedCredits) < 100 {
			sd.LTotemUsedCredits = append(sd.LTotemUsedCredits, 0)
		}
		if totemID < len(sd.LTotemUsedCredits) {
			sd.LTotemUsedCredits[totemID] = credits
		}
	}
	rows.Close()

	// 12) totems_placement
	rows, err = r.db.QueryxContext(ctx, `
SELECT placement_idx, totem_id 
FROM v2_save_data_totems_placement 
WHERE save_id = ?
ORDER BY placement_idx
`, sd.ID)
	if err != nil {
		return nil, err
	}
	sd.LTotemPlacements = make([]int, 0, 100)
	for rows.Next() {
		var placementIdx int
		var totemID int
		if err := rows.Scan(&placementIdx, &totemID); err != nil {
			return nil, err
		}
		for len(sd.LTotemPlacements) <= placementIdx && len(sd.LTotemPlacements) < 100 {
			sd.LTotemPlacements = append(sd.LTotemPlacements, 0)
		}
		if placementIdx < len(sd.LTotemPlacements) {
			sd.LTotemPlacements[placementIdx] = totemID
		}
	}
	rows.Close()

	fmt.Printf("[REPO-DEBUG] GetLatestSave SUCCESS - user_id=%s, achievements=%d, perks=%d, totems=%d\n", userID, len(sd.LAchieve), len(sd.LPerkLevels), len(sd.LTotemLevels))
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
  MAX(sd.updated_at)  AS created_at
FROM v2_save_data_ball_chain AS bc
JOIN v2_save_data           AS sd ON bc.save_id = sd.id
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
  MAX(sd.updated_at)  AS created_at
FROM v2_save_data_ball_chain AS bc
JOIN v2_save_data           AS sd ON bc.save_id = sd.id
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
  MAX(sd.updated_at)    AS created_at
FROM v2_save_data AS sd
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
  FROM v2_save_data
  GROUP BY user_id
) AS t
JOIN v2_save_data AS sd
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

// GetStatisticsV3 returns the latest statistics for V3 (ランキング上限 500).
func (r *Repository) GetStatisticsV3(ctx context.Context) (*models.StatisticsV3, error) {
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 START\n")
	stats := &models.StatisticsV3{}

	// ------ 共通ヘルパ (匿名関数) -----------------------------------------
	addRanking := func(ptr **[]models.RankingEntry, query string, args ...interface{}) error {
		rows, err := r.db.QueryxContext(ctx, query, args...)
		if err != nil {
			return err
		}
		defer rows.Close()

		// 事前に容量を確保してメモリ効率を改善（最大500個）
		list := make([]models.RankingEntry, 0, 500)
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
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 CREATING_RANKING_1 - max_chain_orange\n")
	if err := addRanking(&stats.MaxChainOrange, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    sd.user_id,
    bc.chain_count AS value,
    sd.updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY sd.user_id ORDER BY bc.chain_count DESC, sd.updated_at ASC) AS rn
  FROM v2_save_data_ball_chain AS bc
  JOIN v2_save_data AS sd ON bc.save_id = sd.id
  WHERE bc.ball_id = '1' AND sd.hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
`); err != nil {
		return nil, err
	}

	// 2) max_chain_rainbow (ball_id = '3')
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 CREATING_RANKING_2 - max_chain_rainbow\n")
	if err := addRanking(&stats.MaxChainRainbow, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    sd.user_id,
    bc.chain_count AS value,
    sd.updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY sd.user_id ORDER BY bc.chain_count DESC, sd.updated_at ASC) AS rn
  FROM v2_save_data_ball_chain AS bc
  JOIN v2_save_data AS sd ON bc.save_id = sd.id
  WHERE bc.ball_id = '3' AND sd.hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jack_startmax DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jack_totalmax DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jack_totalmax_v2 DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY ult_totalmax_v2 DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
`); err != nil {
		return nil, err
	}

	// 中間地点でGCを促してメモリ使用量を安定化
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 MIDPOINT_GC_TRIGGER\n")
	runtime.GC()

	// 7) jacksp_startmax
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 CREATING_RANKING_7 - jacksp_startmax\n")
	if err := addRanking(&stats.JackspStartmax, `
SELECT
  ranked.user_id,
  ranked.value,
  ranked.created_at
FROM (
  SELECT
    user_id,
    jacksp_startmax AS value,
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY jacksp_startmax DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    sd.updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY sd.user_id ORDER BY COALESCE(dpg.count, 0) DESC, sd.updated_at ASC) AS rn
  FROM v2_save_data AS sd
  LEFT JOIN v2_save_data_palball_get AS dpg ON dpg.save_id = sd.id AND dpg.ball_id = '100'
  WHERE sd.hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    CAST(cpm_max AS SIGNED) AS value,
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY cpm_max DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY ult_combomax DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY ult_totalmax DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY sp_use DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
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
    updated_at AS created_at,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY buy_shbi DESC, updated_at ASC) AS rn
  FROM v2_save_data
  WHERE hide_record = false
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 500
`); err != nil {
		return nil, err
	}

	// 14) achievements_count
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 CREATING_RANKING_14 - achievements_count\n")
	if err := addRanking(&stats.AchievementsCount, `
SELECT
  sd.user_id,
  COUNT(a.achievement_id) AS value,
  sd.updated_at AS created_at
FROM (
  SELECT user_id, MAX(id) AS max_id
  FROM v2_save_data
  WHERE hide_record = false
  GROUP BY user_id
) AS latest
JOIN v2_save_data AS sd ON sd.id = latest.max_id
LEFT JOIN v2_save_data_achievements AS a ON a.save_id = sd.id
GROUP BY sd.user_id, sd.updated_at
ORDER BY value DESC, created_at ASC
LIMIT 500
`); err != nil {
		return nil, err
	}

	// 最終段階でGCを促してメモリ使用量を安定化
	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 FINAL_GC_TRIGGER\n")
	runtime.GC()

	// 15) total_medals（最新セーブの credit_all 合計）
	var totalMedals int
	{
		q := `
SELECT
  COALESCE(SUM(sd.credit_all),0) AS total_medals
FROM (
  SELECT user_id, MAX(id) AS max_id
  FROM v2_save_data
  WHERE hide_record = false
  GROUP BY user_id
) AS t
JOIN v2_save_data AS sd
  ON sd.user_id = t.user_id
 AND sd.id      = t.max_id
`
		if err := r.db.GetContext(ctx, &totalMedals, q); err != nil {
			return nil, err
		}
		stats.TotalMedals = &totalMedals
	}

	fmt.Printf("[REPO-DEBUG] GetStatisticsV3 SUCCESS - total_medals=%d\n", totalMedals)
	return stats, nil
}
