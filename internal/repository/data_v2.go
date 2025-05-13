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
    rmshbi_get, buy_shbi,
    firstboot, lastsave, playtime
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		sd.UserId, sd.Legacy, sd.Version,
		sd.Credit, sd.CreditAll, sd.MedalIn, sd.MedalGet,
		sd.BallGet, sd.BallChain, sd.SlotStart, sd.SlotStartFev,
		sd.SlotHit, sd.SlotGetFev, sd.SqrGet, sd.SqrStep,
		sd.JackGet, sd.JackStartMax, sd.JackTotalMax,
		sd.UltGet, sd.UltComboMax, sd.UltTotalMax,
		sd.RmShbiGet, sd.BuyShbi,
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
WITH latest_chain AS (
  SELECT 
    sd.user_id,
    bc.chain_count AS value,
    sd.created_at
  FROM save_data_v2_ball_chain AS bc
  JOIN save_data_v2 AS sd ON bc.save_id = sd.id
  WHERE bc.ball_id = '1'
),
max_per_user AS (
  SELECT 
    user_id,
    MAX(value) AS max_value
  FROM latest_chain
  GROUP BY user_id
)
SELECT 
  mpu.user_id,
  mpu.max_value AS value,
  MIN(lc.created_at) AS created_at
FROM max_per_user AS mpu
JOIN latest_chain AS lc
  ON lc.user_id = mpu.user_id
 AND lc.value = mpu.max_value
GROUP BY mpu.user_id, mpu.max_value
ORDER BY mpu.max_value DESC, created_at ASC
LIMIT 500;
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
WITH latest_chain AS (
  SELECT 
    sd.user_id,
    bc.chain_count AS value,
    sd.created_at
  FROM save_data_v2_ball_chain AS bc
  JOIN save_data_v2 AS sd ON bc.save_id = sd.id
  WHERE bc.ball_id = '3'
),
max_per_user AS (
  SELECT 
    user_id,
    MAX(value) AS max_value
  FROM latest_chain
  GROUP BY user_id
)
SELECT 
  mpu.user_id,
  mpu.max_value AS value,
  MIN(lc.created_at) AS created_at
FROM max_per_user AS mpu
JOIN latest_chain AS lc
  ON lc.user_id = mpu.user_id
 AND lc.value = mpu.max_value
GROUP BY mpu.user_id, mpu.max_value
ORDER BY mpu.max_value DESC, created_at ASC
LIMIT 500;
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
WITH max_jp AS (
  SELECT user_id, MAX(jack_totalmax) AS value
  FROM save_data_v2
  GROUP BY user_id
)
SELECT 
  mj.user_id,
  mj.value,
  MIN(sd.created_at) AS created_at
FROM max_jp AS mj
JOIN save_data_v2 AS sd
  ON sd.user_id = mj.user_id
 AND sd.jack_totalmax = mj.value
GROUP BY mj.user_id, mj.value
ORDER BY mj.value DESC, created_at ASC
LIMIT 500;
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
SELECT COALESCE(SUM(sd.credit), 0) AS total_medals
FROM save_data_v2 AS sd
JOIN (
  SELECT user_id, MAX(id) AS max_id
  FROM save_data_v2
  GROUP BY user_id
) AS latest
  ON sd.user_id = latest.user_id
 AND sd.id = latest.max_id;
`
		var total int
		if err := r.db.GetContext(ctx, &total, q); err != nil {
			return nil, err
		}
		stats.TotalMedals = &total
	}

	return stats, nil
}
