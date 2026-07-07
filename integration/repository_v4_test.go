//go:build integration

package integration

import (
	"context"
	"database/sql"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/migration"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
)

var migrateOnce sync.Once

func setupDB(t *testing.T) *sqlx.DB {
	t.Helper()

	cfg := config.MySQL()
	ensureDatabase(t, cfg)

	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		t.Fatalf("connect db: %v", err)
	}

	var migrateErr error
	migrateOnce.Do(func() {
		migrateErr = migration.MigrateTables(db.DB)
	})
	if migrateErr != nil {
		t.Fatalf("migrate: %v", migrateErr)
	}

	cleanupTables(t, db)
	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}

func ensureDatabase(t *testing.T, cfg *mysql.Config) {
	t.Helper()

	if cfg.DBName == "" {
		t.Fatalf("DB_NAME is required for integration tests")
	}

	dbName := cfg.DBName
	cfgCopy := *cfg
	cfgCopy.DBName = ""

	db, err := sql.Open("mysql", cfgCopy.FormatDSN())
	if err != nil {
		t.Fatalf("connect server: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "`"); err != nil {
		t.Fatalf("create database: %v", err)
	}
}

func cleanupTables(t *testing.T, db *sqlx.DB) {
	t.Helper()

	tables := []string{
		"v2_save_data_medal_get",
		"v2_save_data_ball_get",
		"v2_save_data_ball_chain",
		"v2_save_data_achievements",
		"v2_save_data_palball_get",
		"v2_save_data_palball_jp",
		"v2_save_data_bbox_shop",
		"v2_save_data_ferlot_item",
		"v2_save_data_ferlot_useitem",
		"v2_save_data_perks",
		"v2_save_data_perks_credit",
		"v2_save_data_totems",
		"v2_save_data_totems_credit",
		"v2_save_data_totems_placement",
		"v3_user_latest_save_data_achievements",
		"v3_user_latest_save_data",
		"v2_save_data",
		"user_banned",
	}

	if _, err := db.Exec("SET FOREIGN_KEY_CHECKS=0"); err != nil {
		t.Fatalf("disable foreign keys: %v", err)
	}
	for _, table := range tables {
		if _, err := db.Exec("TRUNCATE TABLE " + table); err != nil {
			t.Fatalf("truncate %s: %v", table, err)
		}
	}
	if _, err := db.Exec("SET FOREIGN_KEY_CHECKS=1"); err != nil {
		t.Fatalf("enable foreign keys: %v", err)
	}
}

func newSaveData(userID string, playtime int64, creditAll int64, achievements []string) *domain.SaveData {
	return &domain.SaveData{
		UserId:                      userID,
		Legacy:                      1,
		Version:                     4,
		Credit:                      10,
		CreditAll:                   creditAll,
		MedalIn:                     1,
		MedalGet:                    2,
		BallGet:                     3,
		BallChain:                   4,
		SlotStart:                   5,
		SlotStartFev:                6,
		SlotHit:                     7,
		SlotGetFev:                  8,
		SqrGet:                      9,
		SqrStep:                     10,
		JackGet:                     11,
		JackStartMax:                12,
		JackTotalMax:                13,
		UltGet:                      14,
		UltComboMax:                 15,
		UltTotalMax:                 16,
		RmShbiGet:                   17,
		BuyShbi:                     18,
		FirstBoot:                   100,
		LastSave:                    200,
		Playtime:                    playtime,
		BstpStep:                    1,
		BstpRwd:                     2,
		BuyTotal:                    3,
		SkillPoint:                  4,
		BlackBox:                    5,
		BlackBoxTotal:               6,
		SpUse:                       7,
		HideRecord:                  0,
		CpMMax:                      1.2,
		JackTotalMaxV2:              33,
		UltimateTotalMaxV2:          44,
		PalettaBallGet:              2,
		PalettaLotteryAttemptTier0:  1,
		PalettaLotteryAttemptTier1:  1,
		PalettaLotteryAttemptTier2:  1,
		PalettaLotteryAttemptTier3:  1,
		PalettaLotteryAttemptTier4:  1,
		JackpotSuperGetTotal:        1,
		JackpotSuperGetTier0:        1,
		JackpotSuperGetTier1:        1,
		JackpotSuperGetTier2:        1,
		JackpotSuperGetTier3:        1,
		JackpotSuperGetTier4:        1,
		JackpotSuperStartMax:        50,
		JackpotSuperTotalMax:        60,
		FerrettaBallGet:             7,
		FerrettaLotteryAttempt:      8,
		JackpotFerrettaGetTotal:     9,
		JackpotFerrettaGetTier0:     1,
		JackpotFerrettaGetTier1:     2,
		JackpotFerrettaGetTier2:     3,
		JackpotFerrettaGetTier3:     4,
		JackpotFerrettaGetTier4:     5,
		JackpotFerrettaStartMax:     playtime * 10,
		JackpotFerrettaTotalMax:     playtime * 20,
		FerrettaLotteryHit:          11,
		FerrettaLotteryLose:         12,
		FerrettaLotteryChance:       13,
		FerrettaLotteryActives:      14,
		FerrettaLotteryLines:        int(playtime),
		BlackBoxShopUsed:            15,
		FerrettaLotteryMaxLines:     int(playtime) + 1,
		BlackBoxUsedFerrettaItem:    16,
		GetMedalTower:               int(playtime) + 2,
		TaskCompleteCount:           2,
		TotemAltarUnlockCount:       1,
		TotemAltarUnlockUsedCredits: 20,
		DCMedalGet:                  map[string]int{"1": 2},
		DCBallGet:                   map[string]int64{"1": 3},
		DCBallChain:                 map[string]int{"3": 5},
		LAchieve:                    achievements,
		DCPalettaBallGet:            map[string]int{"100": 4},
		DCBlackBoxShopUsed:          map[string]int{"item-1": 2},
		DCFerrettaLotteryItem:       map[string]int{"item-2": 3},
		DCFerrettaLotteryItemUsed:   map[string]int{"item-3": 4},
	}
}

func TestRepositoryV4_InsertLatestAndExists(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-1"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1", "ach-2"})); err != nil {
		t.Fatalf("insert save1: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 20, 200, []string{"ach-2", "ach-3"})); err != nil {
		t.Fatalf("insert save2: %v", err)
	}

	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE user_id = ? AND playtime = ?", t1, userID, 10); err != nil {
		t.Fatalf("update save1 timestamp: %v", err)
	}
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE user_id = ? AND playtime = ?", t2, userID, 20); err != nil {
		t.Fatalf("update save2 timestamp: %v", err)
	}

	exists, err := repo.ExistsSameSave(ctx, userID, 10)
	if err != nil {
		t.Fatalf("exists: %v", err)
	}
	if !exists {
		t.Fatalf("expected ExistsSameSave true")
	}
	exists, err = repo.ExistsSameSave(ctx, userID, 999)
	if err != nil {
		t.Fatalf("exists: %v", err)
	}
	if exists {
		t.Fatalf("expected ExistsSameSave false")
	}

	latest, err := repo.GetLatestSave(ctx, userID)
	if err != nil {
		t.Fatalf("latest save: %v", err)
	}
	if latest.Playtime != 20 {
		t.Fatalf("latest playtime: got %d", latest.Playtime)
	}
	if latest.CreditAll != 200 {
		t.Fatalf("latest credit_all: got %d", latest.CreditAll)
	}
	if latest.JackpotFerrettaStartMax != 200 {
		t.Fatalf("latest jackfr_startmax: got %d", latest.JackpotFerrettaStartMax)
	}
	if latest.JackpotFerrettaTotalMax != 400 {
		t.Fatalf("latest jackfr_totalmax: got %d", latest.JackpotFerrettaTotalMax)
	}
	if latest.FerrettaLotteryLines != 20 {
		t.Fatalf("latest ferlot_lines: got %d", latest.FerrettaLotteryLines)
	}
	if latest.DCBlackBoxShopUsed["item-1"] != 2 {
		t.Fatalf("latest dc_bbox_shop: got %#v", latest.DCBlackBoxShopUsed)
	}
	if latest.DCFerrettaLotteryItem["item-2"] != 3 {
		t.Fatalf("latest dc_ferlot_item: got %#v", latest.DCFerrettaLotteryItem)
	}
	if latest.FerrettaLotteryMaxLines != 21 {
		t.Fatalf("latest ferlot_maxln: got %d", latest.FerrettaLotteryMaxLines)
	}
	if latest.BlackBoxUsedFerrettaItem != 16 {
		t.Fatalf("latest bbox_used_ferlot: got %d", latest.BlackBoxUsedFerrettaItem)
	}
	if latest.GetMedalTower != 22 {
		t.Fatalf("latest get_medaltower: got %d", latest.GetMedalTower)
	}
	if latest.DCFerrettaLotteryItemUsed["item-3"] != 4 {
		t.Fatalf("latest dc_ferlot_useitem: got %#v", latest.DCFerrettaLotteryItemUsed)
	}

	wantAchievements := []string{"ach-1", "ach-2", "ach-3"}
	if !sameStringSet(latest.LAchieve, wantAchievements) {
		t.Fatalf("achievements: got %#v", latest.LAchieve)
	}
}

func TestRepositoryV4_SaveHistoryAndAchievementsHistory(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-1"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1", "ach-2"})); err != nil {
		t.Fatalf("insert save1: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 20, 200, []string{"ach-2", "ach-3"})); err != nil {
		t.Fatalf("insert save2: %v", err)
	}

	var saves []struct {
		ID       int64 `db:"id"`
		Playtime int64 `db:"playtime"`
	}
	if err := db.Select(&saves, "SELECT id, playtime FROM v2_save_data WHERE user_id = ? ORDER BY playtime ASC", userID); err != nil {
		t.Fatalf("select saves: %v", err)
	}
	if len(saves) != 2 {
		t.Fatalf("expected 2 saves, got %d", len(saves))
	}

	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE id = ?", t1, saves[0].ID); err != nil {
		t.Fatalf("update save1 timestamp: %v", err)
	}
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE id = ?", t2, saves[1].ID); err != nil {
		t.Fatalf("update save2 timestamp: %v", err)
	}

	entries, hasMore, err := repo.GetSaveHistory(ctx, userID, 1, nil)
	if err != nil {
		t.Fatalf("save history: %v", err)
	}
	if !hasMore {
		t.Fatalf("expected hasMore true")
	}
	if len(entries) != 1 || entries[0].Playtime == nil || *entries[0].Playtime != 20 {
		t.Fatalf("latest entry: got %#v", entries)
	}

	entries, hasMore, err = repo.GetSaveHistory(ctx, userID, 10, &t2)
	if err != nil {
		t.Fatalf("save history before: %v", err)
	}
	if hasMore {
		t.Fatalf("expected hasMore false")
	}
	if len(entries) != 1 || entries[0].Playtime == nil || *entries[0].Playtime != 10 {
		t.Fatalf("before entry: got %#v", entries)
	}

	achEntries, total, err := repo.GetAchievementUnlockHistory(ctx, userID, 10)
	if err != nil {
		t.Fatalf("achievement history: %v", err)
	}
	if total != 3 {
		t.Fatalf("achievement total: got %d", total)
	}
	if len(achEntries) != 3 {
		t.Fatalf("achievement entries: got %d", len(achEntries))
	}
	if achEntries[0].UnlockedAt == nil || achEntries[len(achEntries)-1].UnlockedAt == nil {
		t.Fatalf("missing unlock timestamps")
	}
	if !achEntries[0].UnlockedAt.Equal(t1) || !achEntries[len(achEntries)-1].UnlockedAt.Equal(t2) {
		t.Fatalf("unexpected unlock ordering: first=%v last=%v", achEntries[0].UnlockedAt, achEntries[len(achEntries)-1].UnlockedAt)
	}
}

func TestRepositoryV4_StatisticsAndAchievementRates(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()

	if err := repo.InsertSaveV4(ctx, newSaveData("user-1", 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert user1: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData("user-2", 15, 200, []string{"ach-1", "ach-2"})); err != nil {
		t.Fatalf("insert user2: %v", err)
	}

	stats, err := repo.GetStatisticsV4(ctx)
	if err != nil {
		t.Fatalf("statistics v4: %v", err)
	}
	if stats.TotalMedals == nil || *stats.TotalMedals != 300 {
		t.Fatalf("total medals: got %#v", stats.TotalMedals)
	}
	if stats.AchievementsCount == nil || len(*stats.AchievementsCount) != 2 {
		t.Fatalf("achievements count ranking: got %#v", stats.AchievementsCount)
	}
	if stats.JackfrStartmax == nil || len(*stats.JackfrStartmax) != 2 {
		t.Fatalf("jackfr_startmax ranking: got %#v", stats.JackfrStartmax)
	}
	if stats.JackfrTotalmax == nil || len(*stats.JackfrTotalmax) != 2 {
		t.Fatalf("jackfr_totalmax ranking: got %#v", stats.JackfrTotalmax)
	}
	if stats.FerlotLines == nil || len(*stats.FerlotLines) != 2 {
		t.Fatalf("ferlot_lines ranking: got %#v", stats.FerlotLines)
	}

	rates, err := repo.GetAchievementRates(ctx)
	if err != nil {
		t.Fatalf("achievement rates: %v", err)
	}
	if rates.TotalUsers == nil || *rates.TotalUsers != 2 {
		t.Fatalf("total users: got %#v", rates.TotalUsers)
	}
	if rates.AchievementRates == nil {
		t.Fatalf("achievement rates missing")
	}
	rate := (*rates.AchievementRates)["ach-1"]
	if rate.Count == nil || *rate.Count != 2 {
		t.Fatalf("ach-1 count: got %#v", rate.Count)
	}
}

func TestRepositoryV4_IsUserBanned(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-banned-1"

	banned, err := repo.IsUserBanned(ctx, userID)
	if err != nil {
		t.Fatalf("is banned (none): %v", err)
	}
	if banned {
		t.Fatalf("expected not banned for unknown user")
	}

	if _, err := db.Exec("INSERT INTO user_banned (user_id, note) VALUES (?, ?)", userID, "test"); err != nil {
		t.Fatalf("insert ban: %v", err)
	}
	banned, err = repo.IsUserBanned(ctx, userID)
	if err != nil {
		t.Fatalf("is banned: %v", err)
	}
	if !banned {
		t.Fatalf("expected banned after insert")
	}

	if _, err := db.Exec("UPDATE user_banned SET enabled = 0 WHERE user_id = ?", userID); err != nil {
		t.Fatalf("disable ban: %v", err)
	}
	banned, err = repo.IsUserBanned(ctx, userID)
	if err != nil {
		t.Fatalf("is banned after disable: %v", err)
	}
	if banned {
		t.Fatalf("expected not banned when enabled = 0")
	}
}

func TestRepositoryV4_BannedUserForcesHideRecord(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-banned-save"

	if _, err := db.Exec("INSERT INTO user_banned (user_id) VALUES (?)", userID); err != nil {
		t.Fatalf("insert ban: %v", err)
	}

	sd := newSaveData(userID, 10, 100, []string{"ach-1"})
	sd.HideRecord = 0 // クライアントが0を送ってきた想定

	banned, err := repo.IsUserBanned(ctx, userID)
	if err != nil {
		t.Fatalf("is banned: %v", err)
	}
	if banned {
		sd.HideRecord = 1
	}
	if err := repo.InsertSaveV4(ctx, sd); err != nil {
		t.Fatalf("insert save: %v", err)
	}

	var hide int
	if err := db.Get(&hide, "SELECT hide_record FROM v3_user_latest_save_data WHERE user_id = ?", userID); err != nil {
		t.Fatalf("select hide_record: %v", err)
	}
	if hide != 1 {
		t.Fatalf("hide_record after ban: got %d, want 1", hide)
	}

	// BAN無効化後の再セーブでhide_record=0が通ることを確認
	if _, err := db.Exec("UPDATE user_banned SET enabled = 0 WHERE user_id = ?", userID); err != nil {
		t.Fatalf("disable ban: %v", err)
	}
	sd2 := newSaveData(userID, 20, 200, []string{"ach-1"})
	sd2.HideRecord = 0
	banned, err = repo.IsUserBanned(ctx, userID)
	if err != nil {
		t.Fatalf("is banned after disable: %v", err)
	}
	if banned {
		sd2.HideRecord = 1
	}
	if err := repo.InsertSaveV4(ctx, sd2); err != nil {
		t.Fatalf("insert save2: %v", err)
	}
	if err := db.Get(&hide, "SELECT hide_record FROM v3_user_latest_save_data WHERE user_id = ?", userID); err != nil {
		t.Fatalf("select hide_record after unban: %v", err)
	}
	if hide != 0 {
		t.Fatalf("hide_record after unban: got %d, want 0", hide)
	}
}

func sameStringSet(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	gotCopy := append([]string(nil), got...)
	wantCopy := append([]string(nil), want...)
	sort.Strings(gotCopy)
	sort.Strings(wantCopy)
	for i := range gotCopy {
		if gotCopy[i] != wantCopy[i] {
			return false
		}
	}
	return true
}
