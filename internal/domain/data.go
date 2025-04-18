package domain

import "github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"

// GetDataParams から GameData を生成する
func GetDataParamsToGameData(params models.GetDataParams) models.GameData {
	return models.GameData{
		Version:      &params.Version,
		UserId:       &params.UserId,
		HaveMedal:    &params.HaveMedal,
		InMedal:      &params.InMedal,
		OutMedal:     &params.OutMedal,
		SlotHit:      &params.SlotHit,
		GetShirbe:    &params.GetShirbe,
		StartSlot:    &params.StartSlot,
		ShirbeBuy300: &params.ShirbeBuy300,
		Medal1:       &params.Medal1,
		Medal2:       &params.Medal2,
		Medal3:       &params.Medal3,
		Medal4:       &params.Medal4,
		Medal5:       &params.Medal5,
		RMedal:       &params.RMedal,
		Second:       &params.Second,
		Minute:       &params.Minute,
		Hour:         &params.Hour,
		Fever:        &params.Fever,
	}
}

/*
// GameData defines model for GameData.
type GameData struct {
	RMedal       *int    `db:"r_medal" json:"R_medal,omitempty"`
	Fever        *int    `json:"fever,omitempty"`
	GetShirbe    *int    `db:"get_shirbe" json:"get_shirbe,omitempty"`
	HaveMedal    *int    `db:"have_medal" json:"have_medal,omitempty"`
	Hour         *int    `json:"hour,omitempty"`
	InMedal      *int    `db:"in_medal" json:"in_medal,omitempty"`
	Medal1       *int    `db:"medal_1" json:"medal_1,omitempty"`
	Medal2       *int    `db:"medal_2" json:"medal_2,omitempty"`
	Medal3       *int    `db:"medal_3" json:"medal_3,omitempty"`
	Medal4       *int    `db:"medal_4" json:"medal_4,omitempty"`
	Medal5       *int    `db:"medal_5" json:"medal_5,omitempty"`
	Minute       *int    `json:"minute,omitempty"`
	OutMedal     *int    `db:"out_medal" json:"out_medal,omitempty"`
	Second       *int    `json:"second,omitempty"`
	ShirbeBuy300 *int    `db:"shirbe_buy300" json:"shirbe_buy300,omitempty"`
	SlotHit      *int    `db:"slot_hit" json:"slot_hit,omitempty"`
	StartSlot    *int    `db:"start_slot" json:"start_slot,omitempty"`
	UserId       *string `db:"user_id" json:"user_id,omitempty"`
	Version      *string `json:"version,omitempty"`
}

// GetDataParams defines parameters for GetData.
type GetDataParams struct {
	Version      string `form:"version" json:"version"`
	UserId       string `form:"user_id" json:"user_id"`
	HaveMedal    int    `form:"have_medal" json:"have_medal"`
	InMedal      int    `form:"in_medal" json:"in_medal"`
	OutMedal     int    `form:"out_medal" json:"out_medal"`
	SlotHit      int    `form:"slot_hit" json:"slot_hit"`
	GetShirbe    int    `form:"get_shirbe" json:"get_shirbe"`
	StartSlot    int    `form:"start_slot" json:"start_slot"`
	ShirbeBuy300 int    `form:"shirbe_buy300" json:"shirbe_buy300"`
	Medal1       int    `form:"medal_1" json:"medal_1"`
	Medal2       int    `form:"medal_2" json:"medal_2"`
	Medal3       int    `form:"medal_3" json:"medal_3"`
	Medal4       int    `form:"medal_4" json:"medal_4"`
	Medal5       int    `form:"medal_5" json:"medal_5"`
	RMedal       int    `form:"R_medal" json:"R_medal"`
	Second       int    `form:"second" json:"second"`
	Minute       int    `form:"minute" json:"minute"`
	Hour         int    `form:"hour" json:"hour"`
	Fever        int    `form:"fever" json:"fever"`

	// Sig HMAC-SHA256署名（順序固定・user_id込みで生成）
	Sig string `form:"sig" json:"sig"`
}
*/
