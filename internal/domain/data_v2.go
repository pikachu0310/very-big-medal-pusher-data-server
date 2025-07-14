package domain

import (
	"encoding/json"
	"fmt"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
	"net/url"
	"strconv"
	"time"
)

// SaveData holds parsed v2 save data *and* DB metadata.
type SaveData struct {
	ID           int64     `db:"id"`
	UserID       string    `db:"user_id"`
	Legacy       int       `db:"legacy"`
	Version      int       `db:"version"`
	Credit       int64     `db:"credit"`
	CreditAll    int64     `db:"credit_all"`
	MedalIn      int       `db:"medal_in"`
	MedalGet     int       `db:"medal_get"`
	BallGet      int       `db:"ball_get"`
	BallChain    int       `db:"ball_chain"`
	SlotStart    int       `db:"slot_start"`
	SlotStartFev int       `db:"slot_startfev"`
	SlotHit      int       `db:"slot_hit"`
	SlotGetFev   int       `db:"slot_getfev"`
	SqrGet       int       `db:"sqr_get"`
	SqrStep      int       `db:"sqr_step"`
	JackGet      int       `db:"jack_get"`
	JackStartMax int       `db:"jack_startmax"`
	JackTotalMax int       `db:"jack_totalmax"`
	UltGet       int       `db:"ult_get"`
	UltComboMax  int       `db:"ult_combomax"`
	UltTotalMax  int       `db:"ult_totalmax"`
	RmShbiGet    int       `db:"rmshbi_get"`
	BuyShbi      int       `db:"buy_shbi"`
	FirstBoot    int64     `db:"firstboot"`
	LastSave     int64     `db:"lastsave"`
	Playtime     int64     `db:"playtime"`
	BstpStep     int       `db:"bstp_step"`
	BstpRwd      int       `db:"bstp_rwd"`
	BuyTotal     int       `db:"buy_total"`
	SpUse        int       `db:"sp_use"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`

	// child tables (not loaded via SELECT *)
	DCMedalGet  map[string]int `db:"-"`
	DCBallGet   map[string]int `db:"-"`
	DCBallChain map[string]int `db:"-"`
	LAchieve    []string       `db:"-"`
}

// ParseSaveData decodes URL-encoded JSON into a minimal SaveData for insert.
// It does *not* fill ID/CreatedAt/UpdatedAt — those come from the DB.
func ParseSaveData(raw string) (*SaveData, error) {
	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		return nil, err
	}

	var m struct {
		Legacy       *int           `json:"legacy"`
		Version      *int           `json:"version"`
		Credit       *int64         `json:"credit"`
		CreditAll    *int64         `json:"credit_all"`
		MedalIn      *int           `json:"medal_in"`
		MedalGet     *int           `json:"medal_get"`
		BallGet      *int           `json:"ball_get"`
		BallChain    *int           `json:"ball_chain"`
		SlotStart    *int           `json:"slot_start"`
		SlotStartFev *int           `json:"slot_startfev"`
		SlotHit      *int           `json:"slot_hit"`
		SlotGetFev   *int           `json:"slot_getfev"`
		SqrGet       *int           `json:"sqr_get"`
		SqrStep      *int           `json:"sqr_step"`
		JackGet      *int           `json:"jack_get"`
		JackStartMax *int           `json:"jack_startmax"`
		JackTotalMax *int           `json:"jack_totalmax"`
		UltGet       *int           `json:"ult_get"`
		UltComboMax  *int           `json:"ult_combomax"`
		UltTotalMax  *int           `json:"ult_totalmax"`
		RmShbiGet    *int           `json:"rmshbi_get"`
		BuyShbi      *int           `json:"buy_shbi"`
		FirstBoot    *string        `json:"firstboot"`
		LastSave     *string        `json:"lastsave"`
		Playtime     *int64         `json:"playtime"`
		BstpStep     *int           `json:"bstp_step"`
		BstpRwd      *int           `json:"bstp_rwd"`
		BuyTotal     *int           `json:"buy_total"`
		SpUse        *int           `json:"sp_use"`
		DCMedalGet   map[string]int `json:"dc_medal_get"`
		DCBallGet    map[string]int `json:"dc_ball_get"`
		DCBallChain  map[string]int `json:"dc_ball_chain"`
		LAchieve     []interface{}  `json:"l_achieve"`
		UserId       *string        `json:"user_id"`
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
		UserID:       safeString(m.UserId),
		Legacy:       safeInt(m.Legacy),
		Version:      safeInt(m.Version),
		Credit:       safeInt64(m.Credit),
		CreditAll:    safeInt64(m.CreditAll),
		MedalIn:      safeInt(m.MedalIn),
		MedalGet:     safeInt(m.MedalGet),
		BallGet:      safeInt(m.BallGet),
		BallChain:    safeInt(m.BallChain),
		SlotStart:    safeInt(m.SlotStart),
		SlotStartFev: safeInt(m.SlotStartFev),
		SlotHit:      safeInt(m.SlotHit),
		SlotGetFev:   safeInt(m.SlotGetFev),
		SqrGet:       safeInt(m.SqrGet),
		SqrStep:      safeInt(m.SqrStep),
		JackGet:      safeInt(m.JackGet),
		JackStartMax: safeInt(m.JackStartMax),
		JackTotalMax: safeInt(m.JackTotalMax),
		UltGet:       safeInt(m.UltGet),
		UltComboMax:  safeInt(m.UltComboMax),
		UltTotalMax:  safeInt(m.UltTotalMax),
		RmShbiGet:    safeInt(m.RmShbiGet),
		BuyShbi:      safeInt(m.BuyShbi),
		FirstBoot:    parseUnix(m.FirstBoot),
		LastSave:     parseUnix(m.LastSave),
		Playtime:     safeInt64(m.Playtime),
		BstpStep:     safeInt(m.BstpStep),
		BstpRwd:      safeInt(m.BstpRwd),
		BuyTotal:     safeInt(m.BuyTotal),
		SpUse:        safeInt(m.SpUse),
		DCMedalGet:   m.DCMedalGet,
		DCBallGet:    m.DCBallGet,
		DCBallChain:  m.DCBallChain,
		LAchieve:     ach,
	}

	return sd, nil
}

// ToModel converts a domain.SaveData into the OpenAPI model SaveDataV2.
func (sd *SaveData) ToModel() *models.SaveDataV2 {
	// helpers to get pointers
	intPtr := func(i int) *int { return &i }
	int64Ptr := func(i int64) *int64 { return &i }
	strPtr := func(s string) *string { return &s }

	m := &models.SaveDataV2{
		Legacy:       intPtr(sd.Legacy),
		Version:      intPtr(sd.Version),
		Credit:       int64Ptr(sd.Credit),
		CreditAll:    int64Ptr(sd.CreditAll),
		MedalIn:      intPtr(sd.MedalIn),
		MedalGet:     intPtr(sd.MedalGet),
		BallGet:      intPtr(sd.BallGet),
		BallChain:    intPtr(sd.BallChain),
		SlotStart:    intPtr(sd.SlotStart),
		SlotStartfev: intPtr(sd.SlotStartFev),
		SlotHit:      intPtr(sd.SlotHit),
		SlotGetfev:   intPtr(sd.SlotGetFev),
		SqrGet:       intPtr(sd.SqrGet),
		SqrStep:      intPtr(sd.SqrStep),
		JackGet:      intPtr(sd.JackGet),
		JackStartmax: intPtr(sd.JackStartMax),
		JackTotalmax: intPtr(sd.JackTotalMax),
		UltGet:       intPtr(sd.UltGet),
		UltCombomax:  intPtr(sd.UltComboMax),
		UltTotalmax:  intPtr(sd.UltTotalMax),
		RmshbiGet:    intPtr(sd.RmShbiGet),
		BuyShbi:      intPtr(sd.BuyShbi),
		Firstboot:    strPtr(strconv.FormatInt(sd.FirstBoot, 10)),
		Lastsave:     strPtr(strconv.FormatInt(sd.LastSave, 10)),
		Playtime:     int64Ptr(sd.Playtime),
		BstpStep:     intPtr(sd.BstpStep),
		BstpRwd:      intPtr(sd.BstpRwd),
		BuyTotal:     intPtr(sd.BuyTotal),
		SpUse:        intPtr(sd.SpUse),

		// dictionaries and list
		DcMedalGet:  &sd.DCMedalGet,
		DcBallGet:   &sd.DCBallGet,
		DcBallChain: &sd.DCBallChain,
		LAchieve:    &sd.LAchieve,
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
func parseUnix(s *string) int64 {
	if s == nil {
		return 0
	}
	i, _ := strconv.ParseInt(*s, 10, 64)

	return i
}
