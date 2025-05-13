// internal/domain/data_v2.go
package domain

import (
	"encoding/json"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
	"net/url"
	"strconv"
)

// SaveData holds parsed v2 save data
type SaveData struct {
	UserId       string
	Legacy       int
	Version      int
	Credit       int64
	CreditAll    int64
	MedalIn      int
	MedalGet     int
	BallGet      int
	BallChain    int
	SlotStart    int
	SlotStartFev int
	SlotHit      int
	SlotGetFev   int
	SqrGet       int
	SqrStep      int
	JackGet      int
	JackStartMax int
	JackTotalMax int
	UltGet       int
	UltComboMax  int
	UltTotalMax  int
	RmShbiGet    int
	BuyShbi      int
	FirstBoot    int64
	LastSave     int64
	Playtime     int64
	DCMedalGet   map[string]int
	DCBallGet    map[string]int
	DCBallChain  map[string]int
	LAchieve     []string
}

// ParseSaveData decodes URL-encoded JSON into SaveData
func ParseSaveData(raw string) (*SaveData, error) {
	// URL デコード
	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		return nil, err
	}
	// JSONパース
	var m models.SaveDataV2
	if err := json.Unmarshal([]byte(decoded), &m); err != nil {
		return nil, err
	}
	sd := &SaveData{
		Legacy:       int(getInt(m.Legacy)),
		Version:      getInt(m.Version),
		Credit:       getInt64(m.Credit),
		CreditAll:    getInt64(m.CreditAll),
		MedalIn:      getInt(m.MedalIn),
		MedalGet:     getInt(m.MedalGet),
		BallGet:      getInt(m.BallGet),
		BallChain:    getInt(m.BallChain),
		SlotStart:    getInt(m.SlotStart),
		SlotStartFev: getInt(m.SlotStartfev),
		SlotHit:      getInt(m.SlotHit),
		SlotGetFev:   getInt(m.SlotGetfev),
		SqrGet:       getInt(m.SqrGet),
		SqrStep:      getInt(m.SqrStep),
		JackGet:      getInt(m.JackGet),
		JackStartMax: getInt(m.JackStartmax),
		JackTotalMax: getInt(m.JackTotalmax),
		UltGet:       getInt(m.UltGet),
		UltComboMax:  getInt(m.UltCombomax),
		UltTotalMax:  getInt(m.UltTotalmax),
		RmShbiGet:    getInt(m.RmshbiGet),
		BuyShbi:      getInt(m.BuyShbi),
		FirstBoot:    parseUnixString(m.Firstboot),
		LastSave:     parseUnixString(m.Lastsave),
		Playtime:     getInt64(m.Playtime),
		DCMedalGet:   *m.DcMedalGet,
		DCBallGet:    *m.DcBallGet,
		DCBallChain:  *m.DcBallChain,
		LAchieve:     *m.LAchieve,
	}
	return sd, nil
}

func getInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
func getInt64(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}
func parseUnixString(s *string) int64 {
	if s == nil {
		return 0
	}
	i, _ := strconv.ParseInt(*s, 10, 64)
	return i
}
func parseIntString(s *string) int {
	if s == nil {
		return 0
	}
	i, _ := strconv.Atoi(*s)
	return i
}
