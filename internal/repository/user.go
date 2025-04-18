package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type GameData struct {
	ID           uuid.UUID `db:"id"`
	UserID       string    `db:"user_id"`
	Version      string    `db:"version"`
	InMedal      int       `db:"in_medal"`
	OutMedal     int       `db:"out_medal"`
	SlotHit      int       `db:"slot_hit"`
	GetShirbe    int       `db:"get_shirbe"`
	StartSlot    int       `db:"start_slot"`
	ShirbeBuy300 int       `db:"shirbe_buy300"`
	Medal1       int       `db:"medal_1"`
	Medal2       int       `db:"medal_2"`
	Medal3       int       `db:"medal_3"`
	Medal4       int       `db:"medal_4"`
	Medal5       int       `db:"medal_5"`
	RMedal       int       `db:"R_medal"`
	Second       float32   `db:"second"`
	Minute       int       `db:"minute"`
	Hour         int       `db:"hour"`
}

func (r *Repository) InsertGameData(ctx context.Context, data GameData) error {
	_, err := r.db.ExecContext(ctx, `
	INSERT INTO game_data (
		id, user_id, version, in_medal, out_medal, slot_hit,
		get_shirbe, start_slot, shirbe_buy300, medal_1, medal_2,
		medal_3, medal_4, medal_5, R_medal, second, minute, hour
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		data.ID, data.UserID, data.Version, data.InMedal, data.OutMedal, data.SlotHit,
		data.GetShirbe, data.StartSlot, data.ShirbeBuy300, data.Medal1, data.Medal2,
		data.Medal3, data.Medal4, data.Medal5, data.RMedal, data.Second, data.Minute, data.Hour,
	)

	if err != nil {
		return fmt.Errorf("insert game data: %w", err)
	}
	return nil
}
