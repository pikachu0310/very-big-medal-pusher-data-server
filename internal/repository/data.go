package repository

import (
	"context"
	"fmt"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func (r *Repository) InsertGameData(ctx context.Context, data models.GameData) error {
	_, err := r.db.ExecContext(ctx, `
    INSERT INTO game_data (
        user_id, version,
        have_medal, in_medal, out_medal, slot_hit,
        get_shirbe, start_slot, shirbe_buy300,
        medal_1, medal_2, medal_3, medal_4, medal_5,
        R_medal, total_play_time, fever,
        max_chain_item, max_chain_orange, max_chain_rainbow
    ) VALUES (
        ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )`,
		data.UserId, data.Version,
		data.HaveMedal, data.InMedal, data.OutMedal, data.SlotHit,
		data.GetShirbe, data.StartSlot, data.ShirbeBuy300,
		data.Medal1, data.Medal2, data.Medal3, data.Medal4, data.Medal5,
		data.RMedal, data.TotalPlayTime, data.Fever,
		data.MaxChainItem, data.MaxChainOrange, data.MaxChainRainbow,
	)
	if err != nil {
		return fmt.Errorf("insert game data: %w", err)
	}
	return nil
}

func (r *Repository) GetRankings(ctx context.Context, sortBy string, limit int) ([]models.GameData, error) {
	validSortColumns := map[string]bool{
		"have_medal":        true,
		"in_medal":          true,
		"out_medal":         true,
		"slot_hit":          true,
		"get_shirbe":        true,
		"start_slot":        true,
		"shirbe_buy300":     true,
		"medal_1":           true,
		"medal_2":           true,
		"medal_3":           true,
		"medal_4":           true,
		"medal_5":           true,
		"R_medal":           true,
		"second":            true,
		"minute":            true,
		"hour":              true,
		"fever":             true,
		"max_chain_item":    true,
		"max_chain_orange":  true,
		"max_chain_rainbow": true,
	}

	if !validSortColumns[sortBy] {
		return nil, fmt.Errorf("invalid sort column: %s", sortBy)
	}

	query := fmt.Sprintf(`
        SELECT gd.*
        FROM game_data gd
        WHERE gd.id = (
            SELECT id
            FROM game_data
            WHERE user_id = gd.user_id
            ORDER BY gd.%[1]s DESC, id DESC
            LIMIT 1
        )
        ORDER BY gd.%[1]s DESC
        LIMIT ?`, sortBy)

	var rankings []models.GameData
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

// ExistsSameGameData returns true if there is already a game_data row
// with the given userId and totalPlayTime.
func (r *Repository) ExistsSameGameData(ctx context.Context, userId string, totalPlayTime int) (bool, error) {
	var count int64
	err := r.db.GetContext(ctx, &count, `
        SELECT COUNT(*) 
        FROM game_data 
        WHERE user_id = ? AND total_play_time = ?
    `, userId, totalPlayTime)
	if err != nil {
		return false, fmt.Errorf("check existing game data: %w", err)
	}
	return count > 0, nil
}
