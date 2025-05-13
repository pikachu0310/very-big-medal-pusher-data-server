// internal/repository/data_v2.go
package repository

import (
	"context"
	"fmt"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// ExistsSameSave checks duplicate by user_id and playtime
func (r *Repository) ExistsSameSave(ctx context.Context, userID string, playtime int64) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM save_data_v2 WHERE user_id=? AND playtime=?)`
	var exists int
	err := r.db.QueryRowContext(ctx, q, userID, playtime).Scan(&exists)
	return exists == 1, err
}

// InsertSave persists SaveData into normalized tables
func (r *Repository) InsertSave(ctx context.Context, sd *domain.SaveData) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
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
		tx.Rollback()
		return err
	}
	saveID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	// medal_get
	for id, cnt := range sd.DCMedalGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_medal_get(save_id, medal_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			tx.Rollback()
			return err
		}
	}
	// ball_get
	for id, cnt := range sd.DCBallGet {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_ball_get(save_id, ball_id, count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			tx.Rollback()
			return err
		}
	}
	// ball_chain
	for id, cnt := range sd.DCBallChain {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_ball_chain(save_id, ball_id, chain_count) VALUES(?,?,?)`, saveID, id, cnt); err != nil {
			tx.Rollback()
			return err
		}
	}
	// achievements
	for _, aid := range sd.LAchieve {
		if _, err := tx.ExecContext(ctx, `INSERT INTO save_data_v2_achievements(save_id, achievement_id) VALUES(?,?)`, saveID, aid); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// GetLatestSave retrieves one SaveData by user_id
func (r *Repository) GetLatestSave(ctx context.Context, userID string) (*models.SaveDataV2, error) {
	// TODO: implement join queries to assemble models.SaveDataV2
	return nil, fmt.Errorf("not implemented")
}

// GetStatistics returns combined rankings and total_medals
func (r *Repository) GetStatistics(ctx context.Context) (*models.StatisticsV2, error) {
	// TODO: query each ranking and sum and assemble models.StatisticsV2
	return nil, fmt.Errorf("not implemented")
}
