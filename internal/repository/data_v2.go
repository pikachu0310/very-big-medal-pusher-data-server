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
    firstboot, lastsave, playtime
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		sd.UserId, sd.Legacy, sd.Version,
		sd.Credit, sd.CreditAll, sd.MedalIn, sd.MedalGet,
		sd.BallGet, sd.BallChain, sd.SlotStart, sd.SlotStartFev,
		sd.SlotHit, sd.SlotGetFev, sd.SqrGet, sd.SqrStep,
		sd.JackGet, sd.JackStartMax, sd.JackTotalMax,
		sd.UltGet, sd.UltComboMax, sd.UltTotalMax,
		sd.RmShbiGet, sd.BstpStep, sd.BstpRwd, sd.BuyTotal, sd.SpUse, sd.BuyShbi,
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
  WHERE bc.ball_id = '1'
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
  WHERE bc.ball_id = '3'
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
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 5) ult_combomax
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
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 6) ult_totalmax
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
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 7) sp_use
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
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 8) buy_shbi
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
) AS ranked
WHERE ranked.rn = 1
ORDER BY ranked.value DESC, ranked.created_at ASC
LIMIT 1000
`); err != nil {
		return nil, err
	}

	// 9) total_medals（最新セーブの credit 合計）
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
