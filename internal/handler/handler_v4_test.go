package handler

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

const (
	testSaveSecret = "save-secret"
	testLoadSecret = "load-secret"
)

type stubRepo struct {
	existsSameSave bool
	existsErr      error
	insertErr      error
	insertedSave   *domain.SaveData

	latestSave        *domain.SaveData
	latestErr         error
	latestSaveByUser  map[string]*domain.SaveData
	latestErrByUser   map[string]error
	latestSaveUserIDs []string

	saveHistory        []models.SaveHistoryEntry
	saveHistoryHasMore bool
	saveHistoryErr     error
	saveHistoryLimit   int
	saveHistoryBefore  *time.Time
	saveHistoryUserID  string

	achievements      []models.AchievementUnlockEntry
	achievementsTotal int
	achievementsErr   error
	achievementsLimit int
	achievementsUser  string

	statsV4        *models.StatisticsV4
	statsV4Err     error
	statsV4Calls   int
	statsV3        *models.StatisticsV3
	statsV3Err     error
	rankings       []models.GameData
	rankingsErr    error
	totalMedals    int
	totalMedalsErr error

	achievementRates      *models.AchievementRates
	achievementRatesErr   error
	achievementRatesCalls int

	medalTimeseries      *models.MedalTimeseriesResponse
	medalTimeseriesErr   error
	medalTimeseriesCalls []int

	saveActivity      *models.SaveActivityResponse
	saveActivityErr   error
	saveActivityCalls []int

	creditAllDistribution    *models.CreditAllDistributionResponse
	creditAllDistributionErr error
}

func (s *stubRepo) GetRankings(ctx context.Context, sortBy string, limit int) ([]models.GameData, error) {
	return s.rankings, s.rankingsErr
}

func (s *stubRepo) GetTotalMedals(ctx context.Context) (int, error) {
	return s.totalMedals, s.totalMedalsErr
}

func (s *stubRepo) GetStatisticsV3(ctx context.Context) (*models.StatisticsV3, error) {
	return s.statsV3, s.statsV3Err
}

func (s *stubRepo) GetStatisticsV4(ctx context.Context) (*models.StatisticsV4, error) {
	s.statsV4Calls++
	return s.statsV4, s.statsV4Err
}

func (s *stubRepo) GetAchievementRates(ctx context.Context) (*models.AchievementRates, error) {
	s.achievementRatesCalls++
	return s.achievementRates, s.achievementRatesErr
}

func (s *stubRepo) GetMedalTimeseries(ctx context.Context, days int) (*models.MedalTimeseriesResponse, error) {
	s.medalTimeseriesCalls = append(s.medalTimeseriesCalls, days)
	return s.medalTimeseries, s.medalTimeseriesErr
}

func (s *stubRepo) GetSaveActivity(ctx context.Context, hours int) (*models.SaveActivityResponse, error) {
	s.saveActivityCalls = append(s.saveActivityCalls, hours)
	return s.saveActivity, s.saveActivityErr
}

func (s *stubRepo) GetCreditAllDistribution(ctx context.Context) (*models.CreditAllDistributionResponse, error) {
	return s.creditAllDistribution, s.creditAllDistributionErr
}

func (s *stubRepo) ExistsSameSave(ctx context.Context, userID string, playtime int64) (bool, error) {
	return s.existsSameSave, s.existsErr
}

func (s *stubRepo) InsertSaveV4(ctx context.Context, sd *domain.SaveData) error {
	s.insertedSave = sd
	return s.insertErr
}

func (s *stubRepo) GetLatestSave(ctx context.Context, userID string) (*domain.SaveData, error) {
	if s.latestSaveByUser != nil || s.latestErrByUser != nil {
		s.latestSaveUserIDs = append(s.latestSaveUserIDs, userID)
		if s.latestErrByUser != nil {
			if err, ok := s.latestErrByUser[userID]; ok {
				return nil, err
			}
		}
		if s.latestSaveByUser != nil {
			if save, ok := s.latestSaveByUser[userID]; ok {
				return save, nil
			}
		}
		return nil, sql.ErrNoRows
	}
	return s.latestSave, s.latestErr
}

func (s *stubRepo) GetSaveHistory(ctx context.Context, userID string, limit int, before *time.Time) ([]models.SaveHistoryEntry, bool, error) {
	s.saveHistoryUserID = userID
	s.saveHistoryLimit = limit
	s.saveHistoryBefore = before
	return s.saveHistory, s.saveHistoryHasMore, s.saveHistoryErr
}

func (s *stubRepo) GetAchievementUnlockHistory(ctx context.Context, userID string, limit int) ([]models.AchievementUnlockEntry, int, error) {
	s.achievementsUser = userID
	s.achievementsLimit = limit
	return s.achievements, s.achievementsTotal, s.achievementsErr
}

type flowRepo struct {
	saves []*domain.SaveData
}

func (r *flowRepo) GetRankings(ctx context.Context, sortBy string, limit int) ([]models.GameData, error) {
	return nil, nil
}

func (r *flowRepo) GetTotalMedals(ctx context.Context) (int, error) {
	return 0, nil
}

func (r *flowRepo) GetStatisticsV3(ctx context.Context) (*models.StatisticsV3, error) {
	return nil, nil
}

func (r *flowRepo) GetStatisticsV4(ctx context.Context) (*models.StatisticsV4, error) {
	latestByUser := make(map[string]*domain.SaveData)
	for _, sd := range r.saves {
		existing, ok := latestByUser[sd.UserId]
		if !ok || sd.Playtime > existing.Playtime {
			latestByUser[sd.UserId] = sd
		}
	}
	total := 0
	for _, sd := range latestByUser {
		total += int(sd.CreditAll)
	}
	return &models.StatisticsV4{
		TotalMedals: &total,
	}, nil
}

func (r *flowRepo) GetAchievementRates(ctx context.Context) (*models.AchievementRates, error) {
	return &models.AchievementRates{}, nil
}

func (r *flowRepo) GetMedalTimeseries(ctx context.Context, days int) (*models.MedalTimeseriesResponse, error) {
	return &models.MedalTimeseriesResponse{}, nil
}

func (r *flowRepo) GetSaveActivity(ctx context.Context, hours int) (*models.SaveActivityResponse, error) {
	return &models.SaveActivityResponse{}, nil
}

func (r *flowRepo) GetCreditAllDistribution(ctx context.Context) (*models.CreditAllDistributionResponse, error) {
	return &models.CreditAllDistributionResponse{}, nil
}

func (r *flowRepo) ExistsSameSave(ctx context.Context, userID string, playtime int64) (bool, error) {
	for _, sd := range r.saves {
		if sd.UserId == userID && sd.Playtime == playtime {
			return true, nil
		}
	}
	return false, nil
}

func (r *flowRepo) InsertSaveV4(ctx context.Context, sd *domain.SaveData) error {
	r.saves = append(r.saves, sd)
	return nil
}

func (r *flowRepo) GetLatestSave(ctx context.Context, userID string) (*domain.SaveData, error) {
	var latest *domain.SaveData
	for _, sd := range r.saves {
		if sd.UserId != userID {
			continue
		}
		if latest == nil || sd.Playtime > latest.Playtime {
			latest = sd
		}
	}
	if latest == nil {
		return nil, errors.New("not found")
	}
	return latest, nil
}

func (r *flowRepo) GetSaveHistory(ctx context.Context, userID string, limit int, before *time.Time) ([]models.SaveHistoryEntry, bool, error) {
	return nil, false, nil
}

func (r *flowRepo) GetAchievementUnlockHistory(ctx context.Context, userID string, limit int) ([]models.AchievementUnlockEntry, int, error) {
	return nil, 0, nil
}

func setTestSecrets(t *testing.T) {
	t.Setenv("SAVE", testSaveSecret)
	t.Setenv("LOAD", testLoadSecret)
}

func newTestServer(t *testing.T, repo Repository) *echo.Echo {
	t.Helper()
	e := echo.New()
	openapi.RegisterHandlers(e, New(repo))
	return e
}

func makeV4SaveSig(rawUserID, decodedUserID, data string) string {
	dataEncoded := strings.ReplaceAll(url.QueryEscape(data), "+", "%20")
	userIDEncoded := strings.ReplaceAll(url.QueryEscape(rawUserID), "+", "%20")
	signingStr := "data=" + dataEncoded + "&user_id=" + userIDEncoded
	return signPayload(signingStr, generateUserSecretV4(decodedUserID))
}

func makeLoadSig(userID string) string {
	return signPayload(userID, []byte(testLoadSecret))
}

func intPtr(v int) *int {
	return &v
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func TestGetV4Data_Success(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{}
	e := newTestServer(t, repo)

	userID := "user-1"
	payload := `{"playtime":100,"credit_all":500,"version":4}`
	data := base64.RawURLEncoding.EncodeToString([]byte(payload))
	sig := makeV4SaveSig(userID, userID, data)

	q := url.Values{}
	q.Set("data", data)
	q.Set("user_id", userID)
	q.Set("sig", sig)
	req := httptest.NewRequest(http.MethodGet, "/v4/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if repo.insertedSave == nil {
		t.Fatalf("expected InsertSaveV4 to be called")
	}
	if repo.insertedSave.UserId != userID {
		t.Fatalf("inserted user_id: got %q", repo.insertedSave.UserId)
	}
	if repo.insertedSave.Playtime != 100 {
		t.Fatalf("inserted playtime: got %d", repo.insertedSave.Playtime)
	}
	if repo.insertedSave.CreditAll != 500 {
		t.Fatalf("inserted credit_all: got %d", repo.insertedSave.CreditAll)
	}
}

func TestGetV4Data_Duplicate(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{existsSameSave: true}
	e := newTestServer(t, repo)

	userID := "user-1"
	payload := `{"playtime":100}`
	data := base64.RawURLEncoding.EncodeToString([]byte(payload))
	sig := makeV4SaveSig(userID, userID, data)

	q := url.Values{}
	q.Set("data", data)
	q.Set("user_id", userID)
	q.Set("sig", sig)
	req := httptest.NewRequest(http.MethodGet, "/v4/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if repo.insertedSave != nil {
		t.Fatalf("InsertSaveV4 should not be called on duplicate")
	}
}

func TestGetV4Data_InvalidSignature(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{}
	e := newTestServer(t, repo)

	q := url.Values{}
	q.Set("data", "abc")
	q.Set("user_id", "user-1")
	q.Set("sig", "bad")
	req := httptest.NewRequest(http.MethodGet, "/v4/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGetV4Data_ParseError(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{}
	e := newTestServer(t, repo)

	userID := "user-1"
	data := "not-json"
	sig := makeV4SaveSig(userID, userID, data)

	q := url.Values{}
	q.Set("data", data)
	q.Set("user_id", userID)
	q.Set("sig", sig)
	req := httptest.NewRequest(http.MethodGet, "/v4/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGetV4DataVerify(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{}
	e := newTestServer(t, repo)

	userID := "user-1"
	payload := `{"playtime":10}`
	data := base64.RawURLEncoding.EncodeToString([]byte(payload))
	sig := makeV4SaveSig(userID, userID, data)

	q := url.Values{}
	q.Set("data", data)
	q.Set("user_id", userID)
	q.Set("sig", sig)
	req := httptest.NewRequest(http.MethodGet, "/v4/data/verify?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	var resp models.SignatureVerifyResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !resp.Valid {
		t.Fatalf("expected valid signature")
	}
}

func TestGetV4UsersUserIdData_Success(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{
		latestSave: &domain.SaveData{
			UserId:     "user-1",
			Legacy:     1,
			Version:    4,
			Credit:     10,
			CreditAll:  1000,
			Playtime:   123,
			LAchieve:   []string{"a"},
			DCMedalGet: map[string]int{"1": 2},
		},
	}
	e := newTestServer(t, repo)

	rawUserID := base64.RawURLEncoding.EncodeToString([]byte("user-1"))
	sig := makeLoadSig("user-1")

	q := url.Values{}
	q.Set("sig", sig)
	req := httptest.NewRequest(http.MethodGet, "/v4/users/"+rawUserID+"/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}

	var resp models.SignedSaveData
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.Sig != signPayload(resp.Data, []byte(testLoadSecret)) {
		t.Fatalf("signed response has invalid signature")
	}
	decoded, err := base64.StdEncoding.DecodeString(resp.Data)
	if err != nil {
		t.Fatalf("decode data: %v", err)
	}
	var model models.SaveDataV2
	if err := json.Unmarshal(decoded, &model); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if model.Playtime == nil || *model.Playtime != 123 {
		t.Fatalf("playtime: got %#v", model.Playtime)
	}
	if model.CreditAll == nil || *model.CreditAll != "1000" {
		t.Fatalf("credit_all: got %#v", model.CreditAll)
	}
}

func TestGetV4UsersUserIdData_FallbackToRawUserID(t *testing.T) {
	setTestSecrets(t)
	rawUserID := "dekapu_debug"
	decodedUserID, err := decodeUserIDParam(rawUserID)
	if err != nil {
		t.Fatalf("decodeUserIDParam: %v", err)
	}
	if decodedUserID == rawUserID {
		t.Fatalf("expected base64 decode to differ for %q", rawUserID)
	}

	repo := &stubRepo{
		latestSaveByUser: map[string]*domain.SaveData{
			rawUserID: {
				UserId:    rawUserID,
				Legacy:    1,
				Version:   4,
				CreditAll: 1000,
				Playtime:  123,
			},
		},
	}
	e := newTestServer(t, repo)

	q := url.Values{}
	q.Set("sig", makeLoadSig(rawUserID))
	req := httptest.NewRequest(http.MethodGet, "/v4/users/"+rawUserID+"/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if len(repo.latestSaveUserIDs) != 2 || repo.latestSaveUserIDs[0] != decodedUserID || repo.latestSaveUserIDs[1] != rawUserID {
		t.Fatalf("fallback order: got %#v", repo.latestSaveUserIDs)
	}
}

func TestGetV4UsersUserIdData_MissingSignature(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{}
	e := newTestServer(t, repo)

	req := httptest.NewRequest(http.MethodGet, "/v4/users/user-1/data?sig=", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGetV4UsersUserIdData_InvalidSignature(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{latestSave: &domain.SaveData{}}
	e := newTestServer(t, repo)

	req := httptest.NewRequest(http.MethodGet, "/v4/users/user-1/data?sig=bad", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGetV4UsersUserIdDataVerify(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{}
	e := newTestServer(t, repo)

	sig := makeLoadSig("user-1")
	req := httptest.NewRequest(http.MethodGet, "/v4/users/user-1/data/verify?sig="+url.QueryEscape(sig), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	var resp models.SignatureVerifyResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !resp.Valid {
		t.Fatalf("expected valid signature")
	}
}

func TestGetV4Statistics_Cache(t *testing.T) {
	setTestSecrets(t)
	total := 42
	repo := &stubRepo{statsV4: &models.StatisticsV4{TotalMedals: &total}}
	e := newTestServer(t, repo)

	req := httptest.NewRequest(http.MethodGet, "/v4/statistics", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}

	rec2 := httptest.NewRecorder()
	e.ServeHTTP(rec2, req)
	if repo.statsV4Calls != 1 {
		t.Fatalf("expected cache hit, calls=%d", repo.statsV4Calls)
	}
}

func TestGetV4AchievementsRates(t *testing.T) {
	setTestSecrets(t)
	count := 2
	rate := float32(0.5)
	rates := map[string]struct {
		Count *int     `json:"count,omitempty"`
		Rate  *float32 `json:"rate,omitempty"`
	}{
		"ach-1": {Count: &count, Rate: &rate},
	}
	repo := &stubRepo{
		achievementRates: &models.AchievementRates{
			AchievementRates: &rates,
			TotalUsers:       &count,
		},
	}
	e := newTestServer(t, repo)

	req := httptest.NewRequest(http.MethodGet, "/v4/achievements/rates", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGetV4UsersUserIdSaves_LimitAndNextBefore(t *testing.T) {
	setTestSecrets(t)
	updatedAt := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	repo := &stubRepo{
		saveHistory: []models.SaveHistoryEntry{
			{UpdatedAt: timePtr(updatedAt)},
		},
		saveHistoryHasMore: true,
	}
	e := newTestServer(t, repo)

	before := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	q := url.Values{}
	q.Set("sig", makeLoadSig("user-1"))
	q.Set("limit", "200")
	q.Set("before", before.Format(time.RFC3339))
	req := httptest.NewRequest(http.MethodGet, "/v4/users/user-1/saves?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if repo.saveHistoryLimit != 100 {
		t.Fatalf("limit clamp: got %d", repo.saveHistoryLimit)
	}
	if repo.saveHistoryBefore == nil || !repo.saveHistoryBefore.Equal(before) {
		t.Fatalf("before: got %#v", repo.saveHistoryBefore)
	}
	var resp models.SaveHistoryResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.NextBefore == nil || !resp.NextBefore.Equal(updatedAt) {
		t.Fatalf("next_before: got %#v", resp.NextBefore)
	}
}

func TestGetV4UsersUserIdAchievementsHistory_LimitClamp(t *testing.T) {
	setTestSecrets(t)
	repo := &stubRepo{
		achievements:      []models.AchievementUnlockEntry{},
		achievementsTotal: 10,
	}
	e := newTestServer(t, repo)

	q := url.Values{}
	q.Set("sig", makeLoadSig("user-1"))
	q.Set("limit", "5000")
	req := httptest.NewRequest(http.MethodGet, "/v4/users/user-1/achievements/history?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if repo.achievementsLimit != 2000 {
		t.Fatalf("limit clamp: got %d", repo.achievementsLimit)
	}
}

func TestGetV4StatisticsMedalsTimeseries_DaysClamp(t *testing.T) {
	setTestSecrets(t)
	date := openapi_types.Date{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
	repo := &stubRepo{
		medalTimeseries: &models.MedalTimeseriesResponse{
			Buckets: &[]models.MedalTimeseriesBucket{
				{
					Date:        &date,
					TotalMedals: intPtr(100),
					ActiveUsers: intPtr(2),
					AvgPlaytime: intPtr(50),
				},
			},
		},
	}
	e := newTestServer(t, repo)

	req := httptest.NewRequest(http.MethodGet, "/v4/statistics/medals/timeseries?days=0", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if len(repo.medalTimeseriesCalls) != 1 || repo.medalTimeseriesCalls[0] != 1 {
		t.Fatalf("days clamp: got %#v", repo.medalTimeseriesCalls)
	}
}

func TestGetV4StatisticsSavesActivity_HoursClamp(t *testing.T) {
	setTestSecrets(t)
	hour := time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)
	repo := &stubRepo{
		saveActivity: &models.SaveActivityResponse{
			Buckets: &[]models.SaveActivityBucket{
				{HourStart: &hour, Saves: intPtr(1), UniqueUsers: intPtr(1)},
			},
		},
	}
	e := newTestServer(t, repo)

	req := httptest.NewRequest(http.MethodGet, "/v4/statistics/saves/activity?hours=999", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: got %d body=%s", rec.Code, rec.Body.String())
	}
	if len(repo.saveActivityCalls) != 1 || repo.saveActivityCalls[0] != 720 {
		t.Fatalf("hours clamp: got %#v", repo.saveActivityCalls)
	}
}

func TestV4Flow_SaveThenLoadThenStatistics(t *testing.T) {
	setTestSecrets(t)
	repo := &flowRepo{}
	e := newTestServer(t, repo)

	userID := "user-1"
	payloads := []struct {
		playtime  int
		creditAll int
	}{
		{playtime: 10, creditAll: 100},
		{playtime: 20, creditAll: 200},
	}

	for _, p := range payloads {
		payload := `{"playtime":` + strconv.Itoa(p.playtime) + `,"credit_all":` + strconv.Itoa(p.creditAll) + `,"version":4}`
		data := base64.RawURLEncoding.EncodeToString([]byte(payload))
		sig := makeV4SaveSig(userID, userID, data)
		q := url.Values{}
		q.Set("data", data)
		q.Set("user_id", userID)
		q.Set("sig", sig)
		req := httptest.NewRequest(http.MethodGet, "/v4/data?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("save status: got %d body=%s", rec.Code, rec.Body.String())
		}
	}

	q := url.Values{}
	q.Set("sig", makeLoadSig(userID))
	req := httptest.NewRequest(http.MethodGet, "/v4/users/"+userID+"/data?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("load status: got %d body=%s", rec.Code, rec.Body.String())
	}
	var signed models.SignedSaveData
	if err := json.NewDecoder(rec.Body).Decode(&signed); err != nil {
		t.Fatalf("decode: %v", err)
	}
	decoded, err := base64.StdEncoding.DecodeString(signed.Data)
	if err != nil {
		t.Fatalf("decode data: %v", err)
	}
	var model models.SaveDataV2
	if err := json.Unmarshal(decoded, &model); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if model.Playtime == nil || *model.Playtime != 20 {
		t.Fatalf("latest playtime: got %#v", model.Playtime)
	}
	if model.CreditAll == nil || *model.CreditAll != "200" {
		t.Fatalf("latest credit_all: got %#v", model.CreditAll)
	}

	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/v4/statistics", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("stats status: got %d body=%s", rec.Code, rec.Body.String())
	}
	var stats models.StatisticsV4
	if err := json.NewDecoder(rec.Body).Decode(&stats); err != nil {
		t.Fatalf("decode stats: %v", err)
	}
	if stats.TotalMedals == nil || *stats.TotalMedals != 200 {
		t.Fatalf("stats total_medals: got %#v", stats.TotalMedals)
	}
}
