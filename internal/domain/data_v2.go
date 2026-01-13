package domain

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// SaveData holds parsed v2 save data *and* DB metadata.
type SaveData struct {
	ID                          int64     `db:"id"`
	UserId                      string    `db:"user_id"`
	Legacy                      int       `db:"legacy"`
	Version                     int       `db:"version"`
	Credit                      int64     `db:"credit"`
	CreditAll                   int64     `db:"credit_all"`
	MedalIn                     int       `db:"medal_in"`
	MedalGet                    int64     `db:"medal_get"`
	BallGet                     int64     `db:"ball_get"`
	BallChain                   int       `db:"ball_chain"`
	SlotStart                   int64     `db:"slot_start"`
	SlotStartFev                int64     `db:"slot_startfev"`
	SlotHit                     int64     `db:"slot_hit"`
	SlotGetFev                  int64     `db:"slot_getfev"`
	SqrGet                      int64     `db:"sqr_get"`
	SqrStep                     int64     `db:"sqr_step"`
	JackGet                     int64     `db:"jack_get"`
	JackStartMax                int64     `db:"jack_startmax"`
	JackTotalMax                int       `db:"jack_totalmax"`
	UltGet                      int       `db:"ult_get"`
	UltComboMax                 int       `db:"ult_combomax"`
	UltTotalMax                 int       `db:"ult_totalmax"`
	RmShbiGet                   int       `db:"rmshbi_get"`
	BuyShbi                     int       `db:"buy_shbi"`
	FirstBoot                   int64     `db:"firstboot"`
	LastSave                    int64     `db:"lastsave"`
	Playtime                    int64     `db:"playtime"`
	BstpStep                    int64     `db:"bstp_step"`
	BstpRwd                     int64     `db:"bstp_rwd"`
	BuyTotal                    int       `db:"buy_total"`
	SkillPoint                  int       `db:"skill_point"`
	BlackBox                    int       `db:"blackbox"`
	BlackBoxTotal               int64     `db:"blackbox_total"`
	SpUse                       int64     `db:"sp_use"`
	HideRecord                  int       `db:"hide_record"`
	CpMMax                      float64   `db:"cpm_max"`
	JackTotalMaxV2              int64     `db:"jack_totalmax_v2"`
	UltimateTotalMaxV2          int64     `db:"ult_totalmax_v2"`
	PalettaBallGet              int       `db:"palball_get"`
	PalettaLotteryAttemptTier0  int       `db:"pallot_lot_t0"`
	PalettaLotteryAttemptTier1  int       `db:"pallot_lot_t1"`
	PalettaLotteryAttemptTier2  int       `db:"pallot_lot_t2"`
	PalettaLotteryAttemptTier3  int       `db:"pallot_lot_t3"`
	PalettaLotteryAttemptTier4  int       `db:"pallot_lot_t4"`
	JackpotSuperGetTotal        int       `db:"jacksp_get_all"`
	JackpotSuperGetTier0        int       `db:"jacksp_get_t0"`
	JackpotSuperGetTier1        int       `db:"jacksp_get_t1"`
	JackpotSuperGetTier2        int       `db:"jacksp_get_t2"`
	JackpotSuperGetTier3        int       `db:"jacksp_get_t3"`
	JackpotSuperGetTier4        int       `db:"jacksp_get_t4"`
	JackpotSuperStartMax        int64     `db:"jacksp_startmax"`
	JackpotSuperTotalMax        int64     `db:"jacksp_totalmax"`
	TaskCompleteCount           int       `db:"task_cnt"`
	TotemAltarUnlockCount       int       `db:"totem_altars"`
	TotemAltarUnlockUsedCredits int64     `db:"totem_altars_credit"`
	CreatedAt                   time.Time `db:"created_at"`
	UpdatedAt                   time.Time `db:"updated_at"`

	// child tables (not loaded via SELECT *)
	DCMedalGet           map[string]int   `db:"-"`
	DCBallGet            map[string]int64 `db:"-"`
	DCBallChain          map[string]int   `db:"-"`
	LAchieve             []string         `db:"-"`
	DCPalettaBallGet     map[string]int   `db:"-"`
	DCPalettaBallJackpot map[string]int   `db:"-"`
	LPerkLevels          []int            `db:"-"`
	LPerkUsedCredits     []int64          `db:"-"`
	LTotemLevels         []int            `db:"-"`
	LTotemUsedCredits    []int64          `db:"-"`
	LTotemPlacements     []int            `db:"-"`
}

// ParseSaveData decodes URL-encoded JSON into a minimal SaveData for insert.
// It does *not* fill ID/CreatedAt/UpdatedAt — those come from the DB.
func ParseSaveData(raw string) (*SaveData, error) {
	decoded, err := decodeSavePayload(raw)
	if err != nil {
		return nil, err
	}

	var m struct {
		Legacy                      *int             `json:"legacy"`
		Version                     *int             `json:"version"`
		Credit                      json.RawMessage  `json:"credit"`
		CreditAll                   json.RawMessage  `json:"credit_all"`
		MedalIn                     *int             `json:"medal_in"`
		MedalGet                    *int64           `json:"medal_get"`
		BallGet                     *int64           `json:"ball_get"`
		BallChain                   *int             `json:"ball_chain"`
		SlotStart                   *int64           `json:"slot_start"`
		SlotStartFev                *int64           `json:"slot_startfev"`
		SlotHit                     *int64           `json:"slot_hit"`
		SlotGetFev                  *int64           `json:"slot_getfev"`
		SqrGet                      json.RawMessage  `json:"sqr_get"`
		SqrStep                     *int64           `json:"sqr_step"`
		JackGet                     json.RawMessage  `json:"jack_get"`
		JackStartMax                *int64           `json:"jack_startmax"`
		JackTotalMax                *int             `json:"jack_totalmax"`
		UltGet                      *int             `json:"ult_get"`
		UltComboMax                 *int             `json:"ult_combomax"`
		UltTotalMax                 *int             `json:"ult_totalmax"`
		RmShbiGet                   *int             `json:"rmshbi_get"`
		BuyShbi                     *int             `json:"buy_shbi"`
		FirstBoot                   *json.Number     `json:"firstboot"`
		LastSave                    *json.Number     `json:"lastsave"`
		Playtime                    *int64           `json:"playtime"`
		BstpStep                    json.RawMessage  `json:"bstp_step"`
		BstpRwd                     json.RawMessage  `json:"bstp_rwd"`
		BuyTotal                    *int             `json:"buy_total"`
		SkillPoint                  *float64         `json:"sp"`
		BlackBox                    *float64         `json:"bbox"`
		BlackBoxTotal               *float64         `json:"bbox_all"`
		SpUse                       json.RawMessage  `json:"sp_use"`
		HideRecord                  *int             `json:"hide_record"`
		CpMMax                      *float64         `json:"cpm_max"`
		JackTotalMaxV2              *int64           `json:"jack_totalmax_v2"`
		UltimateTotalMaxV2          *int64           `json:"ult_totalmax_v2"`
		PalettaBallGet              *int             `json:"palball_get"`
		PalettaLotteryAttemptTier0  *float64         `json:"pallot_lot_t0"`
		PalettaLotteryAttemptTier1  *float64         `json:"pallot_lot_t1"`
		PalettaLotteryAttemptTier2  *float64         `json:"pallot_lot_t2"`
		PalettaLotteryAttemptTier3  *float64         `json:"pallot_lot_t3"`
		PalettaLotteryAttemptTier4  *float64         `json:"pallot_lot_t4"`
		JackpotSuperGetTotal        *int             `json:"jacksp_get_all"`
		JackpotSuperGetTier0        *int             `json:"jacksp_get_t0"`
		JackpotSuperGetTier1        *int             `json:"jacksp_get_t1"`
		JackpotSuperGetTier2        *int             `json:"jacksp_get_t2"`
		JackpotSuperGetTier3        *int             `json:"jacksp_get_t3"`
		JackpotSuperGetTier4        *int             `json:"jacksp_get_t4"`
		JackpotSuperStartMax        *int64           `json:"jacksp_startmax"`
		JackpotSuperTotalMax        *int64           `json:"jacksp_totalmax"`
		TaskCompleteCount           *float64         `json:"task_cnt"`
		DCMedalGet                  map[string]int   `json:"dc_medal_get"`
		DCBallGet                   map[string]int64 `json:"dc_ball_get"`
		DCBallChain                 map[string]int   `json:"dc_ball_chain"`
		LAchieve                    []interface{}    `json:"l_achieve"`
		DCPalettaBallGet            map[string]int   `json:"dc_palball_get"`
		DCPalettaBallJackpot        map[string]int   `json:"dc_palball_jp"`
		LPerkLevels                 []interface{}    `json:"l_perks"`
		LPerkUsedCredits            []interface{}    `json:"l_perks_credit"`
		TotemAltarUnlockCount       *int             `json:"totem_altars"`
		TotemAltarUnlockUsedCredits *int64           `json:"totem_altars_credit"`
		LTotemLevels                []interface{}    `json:"l_totems"`
		LTotemUsedCredits           []interface{}    `json:"l_totems_credit"`
		LTotemPlacements            []interface{}    `json:"l_totems_set"`
		UserId                      *string          `json:"user_id"`
	}
	if err := json.Unmarshal([]byte(decoded), &m); err != nil {
		return nil, err
	}

	// ---------- ここで数値⇔文字列を吸収 ----------
	var ach []string
	for _, v := range m.LAchieve {
		switch val := v.(type) {
		case string:
			ach = append(ach, val)
		case float64:
			ach = append(ach, strconv.FormatInt(int64(val), 10))
		default:
			ach = append(ach, fmt.Sprint(val))
		}
	}

	sd := &SaveData{
		UserId:                      safeString(m.UserId),
		Legacy:                      safeInt(m.Legacy),
		Version:                     safeInt(m.Version),
		Credit:                      parseInt64Message(m.Credit),
		CreditAll:                   parseInt64Message(m.CreditAll),
		MedalIn:                     safeInt(m.MedalIn),
		MedalGet:                    safeInt64(m.MedalGet),
		BallGet:                     safeInt64(m.BallGet),
		BallChain:                   safeInt(m.BallChain),
		SlotStart:                   safeInt64(m.SlotStart),
		SlotStartFev:                safeInt64(m.SlotStartFev),
		SlotHit:                     safeInt64(m.SlotHit),
		SlotGetFev:                  safeInt64(m.SlotGetFev),
		SqrGet:                      parseInt64Message(m.SqrGet),
		SqrStep:                     safeInt64(m.SqrStep),
		JackGet:                     parseInt64Message(m.JackGet),
		JackStartMax:                safeInt64(m.JackStartMax),
		JackTotalMax:                safeInt(m.JackTotalMax),
		UltGet:                      safeInt(m.UltGet),
		UltComboMax:                 safeInt(m.UltComboMax),
		UltTotalMax:                 safeInt(m.UltTotalMax),
		RmShbiGet:                   safeInt(m.RmShbiGet),
		BuyShbi:                     safeInt(m.BuyShbi),
		FirstBoot:                   parseUnixFromNumber(m.FirstBoot),
		LastSave:                    parseUnixFromNumber(m.LastSave),
		Playtime:                    safeInt64(m.Playtime),
		BstpStep:                    parseInt64Message(m.BstpStep),
		BstpRwd:                     parseInt64Message(m.BstpRwd),
		BuyTotal:                    safeInt(m.BuyTotal),
		SkillPoint:                  int(safeFloat64(m.SkillPoint)),
		BlackBox:                    int(safeFloat64(m.BlackBox)),
		BlackBoxTotal:               int64(safeFloat64(m.BlackBoxTotal)),
		SpUse:                       parseInt64Message(m.SpUse),
		HideRecord:                  safeInt(m.HideRecord),
		CpMMax:                      safeFloat64(m.CpMMax),
		JackTotalMaxV2:              safeInt64(m.JackTotalMaxV2),
		UltimateTotalMaxV2:          safeInt64(m.UltimateTotalMaxV2),
		PalettaBallGet:              safeInt(m.PalettaBallGet),
		PalettaLotteryAttemptTier0:  int(safeFloat64(m.PalettaLotteryAttemptTier0)),
		PalettaLotteryAttemptTier1:  int(safeFloat64(m.PalettaLotteryAttemptTier1)),
		PalettaLotteryAttemptTier2:  int(safeFloat64(m.PalettaLotteryAttemptTier2)),
		PalettaLotteryAttemptTier3:  int(safeFloat64(m.PalettaLotteryAttemptTier3)),
		PalettaLotteryAttemptTier4:  int(safeFloat64(m.PalettaLotteryAttemptTier4)),
		JackpotSuperGetTotal:        safeInt(m.JackpotSuperGetTotal),
		JackpotSuperGetTier0:        safeInt(m.JackpotSuperGetTier0),
		JackpotSuperGetTier1:        safeInt(m.JackpotSuperGetTier1),
		JackpotSuperGetTier2:        safeInt(m.JackpotSuperGetTier2),
		JackpotSuperGetTier3:        safeInt(m.JackpotSuperGetTier3),
		JackpotSuperGetTier4:        safeInt(m.JackpotSuperGetTier4),
		JackpotSuperStartMax:        safeInt64(m.JackpotSuperStartMax),
		JackpotSuperTotalMax:        safeInt64(m.JackpotSuperTotalMax),
		TaskCompleteCount:           int(safeFloat64(m.TaskCompleteCount)),
		DCMedalGet:                  m.DCMedalGet,
		DCBallGet:                   m.DCBallGet,
		DCBallChain:                 m.DCBallChain,
		LAchieve:                    ach,
		DCPalettaBallGet:            m.DCPalettaBallGet,
		DCPalettaBallJackpot:        m.DCPalettaBallJackpot,
		LPerkLevels:                 parseIntArray(m.LPerkLevels),
		LPerkUsedCredits:            parseInt64Array(m.LPerkUsedCredits),
		TotemAltarUnlockCount:       safeInt(m.TotemAltarUnlockCount),
		TotemAltarUnlockUsedCredits: safeInt64(m.TotemAltarUnlockUsedCredits),
		LTotemLevels:                parseIntArray(m.LTotemLevels),
		LTotemUsedCredits:           parseInt64Array(m.LTotemUsedCredits),
		LTotemPlacements:            parseIntArray(m.LTotemPlacements),
	}
	return sd, nil
}

func decodeSavePayload(raw string) (string, error) {
	if decoded, ok := tryDecodeBase64(raw); ok {
		return decoded, nil
	}
	return url.QueryUnescape(raw)
}

func tryDecodeBase64(raw string) (string, bool) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", false
	}
	encodings := []*base64.Encoding{
		base64.RawURLEncoding,
		base64.URLEncoding,
		base64.RawStdEncoding,
		base64.StdEncoding,
	}
	for _, enc := range encodings {
		decoded, err := enc.DecodeString(trimmed)
		if err != nil {
			continue
		}
		str := string(decoded)
		if looksLikeJSON(str) {
			return str, true
		}
	}
	return "", false
}

func looksLikeJSON(s string) bool {
	t := strings.TrimSpace(s)
	return strings.HasPrefix(t, "{") || strings.HasPrefix(t, "[")
}

// ToModel converts a domain.SaveData into the OpenAPI model SaveDataV2.
func (sd *SaveData) ToModel() *models.SaveDataV2 {
	// メモリ効率を改善するため、一時変数を使用してポインタを作成
	var (
		legacy            = sd.Legacy
		version           = sd.Version
		creditStr         = strconv.FormatInt(sd.Credit, 10)
		creditAllStr      = strconv.FormatInt(sd.CreditAll, 10)
		medalIn           = sd.MedalIn
		medalGet          = sd.MedalGet
		ballGet           = sd.BallGet
		ballChain         = sd.BallChain
		slotStart         = sd.SlotStart
		slotStartfev      = sd.SlotStartFev
		slotHit           = sd.SlotHit
		slotGetfev        = sd.SlotGetFev
		sqrGet            = sd.SqrGet
		sqrStep           = sd.SqrStep
		jackGet           = sd.JackGet
		jackStartmax      = sd.JackStartMax
		jackTotalmax      = sd.JackTotalMax
		ultGet            = sd.UltGet
		ultCombomax       = sd.UltComboMax
		ultTotalmax       = sd.UltTotalMax
		rmshbiGet         = sd.RmShbiGet
		buyShbi           = sd.BuyShbi
		firstboot         = strconv.FormatInt(sd.FirstBoot, 10)
		lastsave          = strconv.FormatInt(sd.LastSave, 10)
		playtime          = sd.Playtime
		bstpStep          = sd.BstpStep
		bstpRwd           = sd.BstpRwd
		buyTotal          = sd.BuyTotal
		sp                = float64(sd.SkillPoint)
		bbox              = float64(sd.BlackBox)
		bboxAll           = float64(sd.BlackBoxTotal)
		spUse             = sd.SpUse
		hideRecord        = sd.HideRecord
		cpmMax            = sd.CpMMax
		jackTotalmaxV2    = sd.JackTotalMaxV2
		ultTotalmaxV2     = sd.UltimateTotalMaxV2
		palballGet        = float64(sd.PalettaBallGet)
		pallotLotT0       = float64(sd.PalettaLotteryAttemptTier0)
		pallotLotT1       = float64(sd.PalettaLotteryAttemptTier1)
		pallotLotT2       = float64(sd.PalettaLotteryAttemptTier2)
		pallotLotT3       = float64(sd.PalettaLotteryAttemptTier3)
		pallotLotT4       = float64(sd.PalettaLotteryAttemptTier4)
		jackspGetAll      = float64(sd.JackpotSuperGetTotal)
		jackspGetT0       = float64(sd.JackpotSuperGetTier0)
		jackspGetT1       = float64(sd.JackpotSuperGetTier1)
		jackspGetT2       = float64(sd.JackpotSuperGetTier2)
		jackspGetT3       = float64(sd.JackpotSuperGetTier3)
		jackspGetT4       = float64(sd.JackpotSuperGetTier4)
		jackspStartmax    = float64(sd.JackpotSuperStartMax)
		jackspTotalmax    = float64(sd.JackpotSuperTotalMax)
		taskCnt           = float64(sd.TaskCompleteCount)
		totemAltars       = sd.TotemAltarUnlockCount
		totemAltarsCredit = sd.TotemAltarUnlockUsedCredits
	)

	m := &models.SaveDataV2{
		Legacy:            &legacy,
		Version:           &version,
		Credit:            &creditStr,
		CreditAll:         &creditAllStr,
		MedalIn:           &medalIn,
		MedalGet:          &medalGet,
		BallGet:           &ballGet,
		BallChain:         &ballChain,
		SlotStart:         &slotStart,
		SlotStartfev:      &slotStartfev,
		SlotHit:           &slotHit,
		SlotGetfev:        &slotGetfev,
		SqrGet:            &sqrGet,
		SqrStep:           &sqrStep,
		JackGet:           &jackGet,
		JackStartmax:      &jackStartmax,
		JackTotalmax:      &jackTotalmax,
		UltGet:            &ultGet,
		UltCombomax:       &ultCombomax,
		UltTotalmax:       &ultTotalmax,
		RmshbiGet:         &rmshbiGet,
		BuyShbi:           &buyShbi,
		Firstboot:         &firstboot,
		Lastsave:          &lastsave,
		Playtime:          &playtime,
		BstpStep:          &bstpStep,
		BstpRwd:           &bstpRwd,
		BuyTotal:          &buyTotal,
		Sp:                &sp,
		Bbox:              &bbox,
		BboxAll:           &bboxAll,
		SpUse:             &spUse,
		HideRecord:        &hideRecord,
		CpmMax:            &cpmMax,
		JackTotalmaxV2:    &jackTotalmaxV2,
		UltTotalmaxV2:     &ultTotalmaxV2,
		PalballGet:        &palballGet,
		PallotLotT0:       &pallotLotT0,
		PallotLotT1:       &pallotLotT1,
		PallotLotT2:       &pallotLotT2,
		PallotLotT3:       &pallotLotT3,
		PallotLotT4:       &pallotLotT4,
		JackspGetAll:      &jackspGetAll,
		JackspGetT0:       &jackspGetT0,
		JackspGetT1:       &jackspGetT1,
		JackspGetT2:       &jackspGetT2,
		JackspGetT3:       &jackspGetT3,
		JackspGetT4:       &jackspGetT4,
		JackspStartmax:    &jackspStartmax,
		JackspTotalmax:    &jackspTotalmax,
		TaskCnt:           &taskCnt,
		TotemAltars:       &totemAltars,
		TotemAltarsCredit: &totemAltarsCredit,

		// dictionaries and list
		DcMedalGet:    &sd.DCMedalGet,
		DcBallGet:     &sd.DCBallGet,
		DcBallChain:   &sd.DCBallChain,
		LAchieve:      &sd.LAchieve,
		DcPalballGet:  &sd.DCPalettaBallGet,
		DcPalballJp:   &sd.DCPalettaBallJackpot,
		LPerks:        &sd.LPerkLevels,
		LPerksCredit:  &sd.LPerkUsedCredits,
		LTotems:       &sd.LTotemLevels,
		LTotemsCredit: &sd.LTotemUsedCredits,
		LTotemsSet:    &sd.LTotemPlacements,
	}

	return m
}

func safeInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

func safeInt64(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

func safeString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func safeFloat64(p *float64) float64 {
	if p == nil {
		return 0.0
	}
	return *p
}

func parseIntArray(arr []interface{}) []int {
	if arr == nil {
		return []int{}
	}
	var result []int
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			result = append(result, val)
		case float64:
			result = append(result, int(val))
		default:
			result = append(result, 0)
		}
	}
	return result
}

func parseInt64Array(arr []interface{}) []int64 {
	if arr == nil {
		return []int64{}
	}
	var result []int64
	for _, v := range arr {
		switch val := v.(type) {
		case int:
			result = append(result, int64(val))
		case float64:
			result = append(result, int64(val))
		default:
			result = append(result, 0)
		}
	}
	return result
}

func parseInt64Message(raw json.RawMessage) int64 {
	if len(raw) == 0 {
		return 0
	}
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return 0
	}

	var num json.Number
	if err := json.Unmarshal(trimmed, &num); err == nil {
		if i, err := num.Int64(); err == nil {
			return i
		}
		if f, err := num.Float64(); err == nil {
			return int64(f)
		}
	}

	var s string
	if err := json.Unmarshal(trimmed, &s); err == nil {
		if s == "" {
			return 0
		}
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return i
		}
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return int64(f)
		}
	}
	return 0
}

func parseUnixFromNumber(n *json.Number) int64 {
	if n == nil {
		return 0
	}
	i, err := n.Int64()
	if err != nil {
		return 0
	}
	return i
}
