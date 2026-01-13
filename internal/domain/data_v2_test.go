package domain

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

func TestParseSaveData_Base64AndUrlEncoded(t *testing.T) {
	payload := `{"legacy":1,"version":4,"credit":"100","credit_all":200,"medal_in":3,"medal_get":4,"ball_get":5,"ball_chain":6,"sqr_get":"7","jack_get":8,"firstboot":12345,"lastsave":23456,"playtime":789,"task_cnt":3.9,"l_achieve":["a",2],"l_perks":[1,2],"l_perks_credit":[10,20],"l_totems":[3],"l_totems_credit":[30],"l_totems_set":[2]}`
	base64Payload := base64.RawURLEncoding.EncodeToString([]byte(payload))
	urlPayload := url.QueryEscape(payload)

	assertParsed := func(t *testing.T, raw string) {
		t.Helper()
		sd, err := ParseSaveData(raw)
		if err != nil {
			t.Fatalf("ParseSaveData error: %v", err)
		}
		if sd.Credit != 100 {
			t.Fatalf("Credit: got %d", sd.Credit)
		}
		if sd.CreditAll != 200 {
			t.Fatalf("CreditAll: got %d", sd.CreditAll)
		}
		if sd.TaskCompleteCount != 3 {
			t.Fatalf("TaskCompleteCount: got %d", sd.TaskCompleteCount)
		}
		wantAch := []string{"a", "2"}
		if !reflect.DeepEqual(sd.LAchieve, wantAch) {
			t.Fatalf("LAchieve: got %#v", sd.LAchieve)
		}
		if len(sd.LPerkLevels) != 2 || sd.LPerkLevels[1] != 2 {
			t.Fatalf("LPerkLevels: got %#v", sd.LPerkLevels)
		}
		if len(sd.LTotemPlacements) != 1 || sd.LTotemPlacements[0] != 2 {
			t.Fatalf("LTotemPlacements: got %#v", sd.LTotemPlacements)
		}
	}

	t.Run("base64", func(t *testing.T) {
		assertParsed(t, base64Payload)
	})
	t.Run("urlencoded", func(t *testing.T) {
		assertParsed(t, urlPayload)
	})
}

func TestDecodeSavePayload(t *testing.T) {
	rawJSON := `{"credit":1}`
	base64Payload := base64.RawURLEncoding.EncodeToString([]byte(rawJSON))

	decoded, err := decodeSavePayload(base64Payload)
	if err != nil {
		t.Fatalf("decodeSavePayload error: %v", err)
	}
	if decoded != rawJSON {
		t.Fatalf("decoded base64: got %q", decoded)
	}

	plain, err := decodeSavePayload("not-base64")
	if err != nil {
		t.Fatalf("decodeSavePayload error: %v", err)
	}
	if plain != "not-base64" {
		t.Fatalf("decoded plain: got %q", plain)
	}
}

func TestTryDecodeBase64(t *testing.T) {
	rawJSON := `{"credit":1}`
	base64Payload := base64.RawURLEncoding.EncodeToString([]byte(rawJSON))
	decoded, ok := tryDecodeBase64(base64Payload)
	if !ok {
		t.Fatalf("tryDecodeBase64 expected ok")
	}
	if decoded != rawJSON {
		t.Fatalf("tryDecodeBase64 decoded: got %q", decoded)
	}

	_, ok = tryDecodeBase64(base64.RawURLEncoding.EncodeToString([]byte("not json")))
	if ok {
		t.Fatalf("tryDecodeBase64 expected false for non json")
	}
}

func TestLooksLikeJSON(t *testing.T) {
	if !looksLikeJSON(` {"a":1}`) {
		t.Fatalf("looksLikeJSON expected true")
	}
	if looksLikeJSON("nope") {
		t.Fatalf("looksLikeJSON expected false")
	}
}

func TestParseInt64Message(t *testing.T) {
	cases := []struct {
		name string
		raw  json.RawMessage
		want int64
	}{
		{name: "empty", raw: nil, want: 0},
		{name: "null", raw: json.RawMessage("null"), want: 0},
		{name: "number", raw: json.RawMessage("123"), want: 123},
		{name: "float", raw: json.RawMessage("123.9"), want: 123},
		{name: "string", raw: json.RawMessage(`"456"`), want: 456},
		{name: "string-float", raw: json.RawMessage(`"789.1"`), want: 789},
		{name: "invalid", raw: json.RawMessage(`"bad"`), want: 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := parseInt64Message(tc.raw); got != tc.want {
				t.Fatalf("got %d want %d", got, tc.want)
			}
		})
	}
}

func TestParseIntArrays(t *testing.T) {
	ints := parseIntArray([]interface{}{1, 2.9, "bad"})
	if !reflect.DeepEqual(ints, []int{1, 2, 0}) {
		t.Fatalf("parseIntArray: got %#v", ints)
	}

	int64s := parseInt64Array([]interface{}{1, 2.9, "bad"})
	if !reflect.DeepEqual(int64s, []int64{1, 2, 0}) {
		t.Fatalf("parseInt64Array: got %#v", int64s)
	}
}

func TestSaveDataToModel(t *testing.T) {
	sd := &SaveData{
		Legacy:     1,
		Version:    4,
		Credit:     10,
		CreditAll:  20,
		Playtime:   30,
		LAchieve:   []string{"a"},
		DCMedalGet: map[string]int{"1": 2},
	}

	model := sd.ToModel()
	if model.Credit == nil || *model.Credit != "10" {
		t.Fatalf("Credit: got %#v", model.Credit)
	}
	if model.CreditAll == nil || *model.CreditAll != "20" {
		t.Fatalf("CreditAll: got %#v", model.CreditAll)
	}
	if model.Playtime == nil || *model.Playtime != 30 {
		t.Fatalf("Playtime: got %#v", model.Playtime)
	}
	if model.LAchieve == nil || len(*model.LAchieve) != 1 || (*model.LAchieve)[0] != "a" {
		t.Fatalf("LAchieve: got %#v", model.LAchieve)
	}
}
