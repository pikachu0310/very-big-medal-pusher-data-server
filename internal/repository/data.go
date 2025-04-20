package repository

import (
	"context"
	"fmt"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func (r *Repository) InsertGameData(ctx context.Context, data models.GameData) error {
	_, err := r.db.ExecContext(ctx, `
    INSERT INTO game_data (
        user_id, version, have_medal, in_medal, out_medal, slot_hit,
        get_shirbe, start_slot, shirbe_buy300, medal_1, medal_2,
        medal_3, medal_4, medal_5, R_medal, total_play_time, fever
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		data.UserId, data.Version, data.HaveMedal, data.InMedal, data.OutMedal, data.SlotHit,
		data.GetShirbe, data.StartSlot, data.ShirbeBuy300, data.Medal1, data.Medal2,
		data.Medal3, data.Medal4, data.Medal5, data.RMedal, data.TotalPlayTime, data.Fever, // 修正済
	)

	if err != nil {
		return fmt.Errorf("insert game data: %w", err)
	}
	return nil
}

func (r *Repository) GetRankings(ctx context.Context, sortBy string, limit int) ([]models.GameData, error) {
	// カラムのホワイトリスト（SQLインジェクション防止）
	validSortColumns := map[string]bool{
		"have_medal":    true,
		"in_medal":      true,
		"out_medal":     true,
		"slot_hit":      true,
		"get_shirbe":    true,
		"start_slot":    true,
		"shirbe_buy300": true,
		"medal_1":       true,
		"medal_2":       true,
		"medal_3":       true,
		"medal_4":       true,
		"medal_5":       true,
		"R_medal":       true,
		"second":        true,
		"minute":        true,
		"hour":          true,
		"fever":         true,
	}

	if !validSortColumns[sortBy] {
		return nil, fmt.Errorf("invalid sort column: %s", sortBy)
	}

	var rankings []models.GameData
	query := fmt.Sprintf(`
		SELECT * FROM (
			SELECT *,
				ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY %[1]s DESC) AS rn
			FROM game_data
		) AS ranked
		WHERE rn = 1
		ORDER BY %[1]s DESC
		LIMIT ?`, sortBy)

	if err := r.db.SelectContext(ctx, &rankings, query, limit); err != nil {
		return nil, fmt.Errorf("get rankings: %w", err)
	}

	return rankings, nil
}

func (r *Repository) GetUserGameData(ctx context.Context, userId string) (*models.GameData, error) {
	var data models.GameData
	if err := r.db.GetContext(ctx, &data, `SELECT * FROM game_data WHERE user_id = ? ORDER BY created_at DESC LIMIT 1`, userId); err != nil {
		return nil, fmt.Errorf("get user game data: %w", err)
	}
	return &data, nil
}
