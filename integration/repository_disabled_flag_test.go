//go:build integration

package integration

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
)

// TestDisabledFlag_DefaultZero は INSERT 時に disabled が DEFAULT 0 で入ることを確認する。
func TestDisabledFlag_DefaultZero(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-default"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert save: %v", err)
	}

	var disabled int
	if err := db.Get(&disabled, "SELECT disabled FROM v2_save_data WHERE user_id = ?", userID); err != nil {
		t.Fatalf("select disabled: %v", err)
	}
	if disabled != 0 {
		t.Fatalf("expected disabled=0, got %d", disabled)
	}
}

// TestDisabledFlag_GetLatestSave は disabled=1 のセーブが LOAD API から除外され、
// 直前の有効なセーブが返ることを確認する。
func TestDisabledFlag_GetLatestSave(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-latest"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert save1: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 20, 200, []string{"ach-2"})); err != nil {
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

	// 通常: 最新 (playtime=20) が返る
	latest, err := repo.GetLatestSave(ctx, userID)
	if err != nil {
		t.Fatalf("latest: %v", err)
	}
	if latest.Playtime != 20 {
		t.Fatalf("expected playtime=20, got %d", latest.Playtime)
	}

	// 最新セーブを disabled に
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE user_id = ? AND playtime = ?", userID, 20); err != nil {
		t.Fatalf("disable save2: %v", err)
	}

	// 直前の有効セーブ (playtime=10) にフォールバック
	latest, err = repo.GetLatestSave(ctx, userID)
	if err != nil {
		t.Fatalf("latest after disable: %v", err)
	}
	if latest.Playtime != 10 {
		t.Fatalf("expected playtime=10 (fallback), got %d", latest.Playtime)
	}

	// 全セーブを disabled にした場合は sql.ErrNoRows
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE user_id = ?", userID); err != nil {
		t.Fatalf("disable all: %v", err)
	}
	if _, err := repo.GetLatestSave(ctx, userID); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected sql.ErrNoRows when all disabled, got %v", err)
	}
}

// TestDisabledFlag_GetSaveHistory は disabled=1 のセーブが履歴 API から除外されることを確認する。
func TestDisabledFlag_GetSaveHistory(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-history"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert save1: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 20, 200, []string{"ach-2"})); err != nil {
		t.Fatalf("insert save2: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 30, 300, []string{"ach-3"})); err != nil {
		t.Fatalf("insert save3: %v", err)
	}

	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC)
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE user_id = ? AND playtime = ?", t1, userID, 10); err != nil {
		t.Fatalf("update save1 timestamp: %v", err)
	}
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE user_id = ? AND playtime = ?", t2, userID, 20); err != nil {
		t.Fatalf("update save2 timestamp: %v", err)
	}
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ? WHERE user_id = ? AND playtime = ?", t3, userID, 30); err != nil {
		t.Fatalf("update save3 timestamp: %v", err)
	}

	// 真ん中 (playtime=20) を disable
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE user_id = ? AND playtime = ?", userID, 20); err != nil {
		t.Fatalf("disable save2: %v", err)
	}

	entries, hasMore, err := repo.GetSaveHistory(ctx, userID, 10, nil)
	if err != nil {
		t.Fatalf("save history: %v", err)
	}
	if hasMore {
		t.Fatalf("expected hasMore=false, got true")
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries (disabled excluded), got %d", len(entries))
	}
	if entries[0].Playtime == nil || *entries[0].Playtime != 30 {
		t.Fatalf("expected first=30, got %#v", entries[0].Playtime)
	}
	if entries[1].Playtime == nil || *entries[1].Playtime != 10 {
		t.Fatalf("expected second=10 (skipping disabled 20), got %#v", entries[1].Playtime)
	}
}

// TestDisabledFlag_GetAchievementUnlockHistory は disabled=1 セーブで解除した実績が履歴に出ないことを確認する。
func TestDisabledFlag_GetAchievementUnlockHistory(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-ach-history"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert save1: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 20, 200, []string{"ach-2"})); err != nil {
		t.Fatalf("insert save2: %v", err)
	}

	// save2 を disable
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE user_id = ? AND playtime = ?", userID, 20); err != nil {
		t.Fatalf("disable save2: %v", err)
	}

	entries, total, err := repo.GetAchievementUnlockHistory(ctx, userID, 10)
	if err != nil {
		t.Fatalf("achievement history: %v", err)
	}
	if total != 1 {
		t.Fatalf("expected total=1 (disabled save's achievement excluded), got %d", total)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].AchievementId == nil || *entries[0].AchievementId != "ach-1" {
		t.Fatalf("expected ach-1, got %#v", entries[0].AchievementId)
	}
}

// TestDisabledFlag_GetMedalTimeseries は disabled=1 セーブが時系列統計から除外されることを確認する。
func TestDisabledFlag_GetMedalTimeseries(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()

	if err := repo.InsertSaveV4(ctx, newSaveData("user-a", 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert user-a: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData("user-b", 15, 500, []string{"ach-1"})); err != nil {
		t.Fatalf("insert user-b: %v", err)
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ?", today); err != nil {
		t.Fatalf("set updated_at: %v", err)
	}

	// disable する前
	resp, err := repo.GetMedalTimeseries(ctx, 7)
	if err != nil {
		t.Fatalf("timeseries before: %v", err)
	}
	if resp.Buckets == nil || len(*resp.Buckets) != 1 {
		t.Fatalf("expected 1 bucket, got %#v", resp.Buckets)
	}
	bucket := (*resp.Buckets)[0]
	if bucket.TotalMedals == nil || *bucket.TotalMedals != 600 {
		t.Fatalf("expected total=600, got %#v", bucket.TotalMedals)
	}
	if bucket.ActiveUsers == nil || *bucket.ActiveUsers != 2 {
		t.Fatalf("expected active=2, got %#v", bucket.ActiveUsers)
	}

	// user-b を disable
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE user_id = 'user-b'"); err != nil {
		t.Fatalf("disable user-b: %v", err)
	}

	resp, err = repo.GetMedalTimeseries(ctx, 7)
	if err != nil {
		t.Fatalf("timeseries after: %v", err)
	}
	if resp.Buckets == nil || len(*resp.Buckets) != 1 {
		t.Fatalf("expected 1 bucket after disable, got %#v", resp.Buckets)
	}
	bucket = (*resp.Buckets)[0]
	if bucket.TotalMedals == nil || *bucket.TotalMedals != 100 {
		t.Fatalf("expected total=100 (user-b excluded), got %#v", bucket.TotalMedals)
	}
	if bucket.ActiveUsers == nil || *bucket.ActiveUsers != 1 {
		t.Fatalf("expected active=1 (user-b excluded), got %#v", bucket.ActiveUsers)
	}
}

// TestDisabledFlag_GetSaveActivity は disabled=1 セーブが活動統計から除外されることを確認する。
func TestDisabledFlag_GetSaveActivity(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()

	if err := repo.InsertSaveV4(ctx, newSaveData("user-a", 10, 100, []string{"ach-1"})); err != nil {
		t.Fatalf("insert user-a: %v", err)
	}
	if err := repo.InsertSaveV4(ctx, newSaveData("user-b", 15, 200, []string{"ach-1"})); err != nil {
		t.Fatalf("insert user-b: %v", err)
	}

	now := time.Now().UTC().Truncate(time.Hour)
	if _, err := db.Exec("UPDATE v2_save_data SET updated_at = ?", now); err != nil {
		t.Fatalf("set updated_at: %v", err)
	}

	// user-a を disable
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE user_id = 'user-a'"); err != nil {
		t.Fatalf("disable user-a: %v", err)
	}

	resp, err := repo.GetSaveActivity(ctx, 24)
	if err != nil {
		t.Fatalf("save activity: %v", err)
	}
	if resp.Buckets == nil || len(*resp.Buckets) != 1 {
		t.Fatalf("expected 1 bucket, got %#v", resp.Buckets)
	}
	bucket := (*resp.Buckets)[0]
	if bucket.Saves == nil || *bucket.Saves != 1 {
		t.Fatalf("expected saves=1 (user-a excluded), got %#v", bucket.Saves)
	}
	if bucket.UniqueUsers == nil || *bucket.UniqueUsers != 1 {
		t.Fatalf("expected unique=1, got %#v", bucket.UniqueUsers)
	}
}

// TestDisabledFlag_ChildTablesPreserved は disabled=1 にしても子テーブルの行が保全されていることを確認する。
// 不正データの事後分析を可能にするための要件。
func TestDisabledFlag_ChildTablesPreserved(t *testing.T) {
	db := setupDB(t)
	repo := repository.New(db)

	ctx := context.Background()
	userID := "user-child"

	if err := repo.InsertSaveV4(ctx, newSaveData(userID, 10, 100, []string{"ach-1", "ach-2"})); err != nil {
		t.Fatalf("insert save: %v", err)
	}

	var saveID int64
	if err := db.Get(&saveID, "SELECT id FROM v2_save_data WHERE user_id = ?", userID); err != nil {
		t.Fatalf("get save id: %v", err)
	}

	// disable する
	if _, err := db.Exec("UPDATE v2_save_data SET disabled = 1 WHERE id = ?", saveID); err != nil {
		t.Fatalf("disable: %v", err)
	}

	// v2_save_data 本体は残っている
	var count int
	if err := db.Get(&count, "SELECT COUNT(*) FROM v2_save_data WHERE id = ?", saveID); err != nil {
		t.Fatalf("count v2_save_data: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected v2_save_data row preserved, got %d", count)
	}

	// 子テーブルも残っている（事後分析用）
	childTables := []string{
		"v2_save_data_achievements",
		"v2_save_data_medal_get",
		"v2_save_data_ball_get",
		"v2_save_data_ball_chain",
		"v2_save_data_palball_get",
		"v2_save_data_bbox_shop",
		"v2_save_data_ferlot_item",
		"v2_save_data_ferlot_useitem",
	}
	for _, table := range childTables {
		var n int
		if err := db.Get(&n, "SELECT COUNT(*) FROM "+table+" WHERE save_id = ?", saveID); err != nil {
			t.Fatalf("count %s: %v", table, err)
		}
		if n == 0 {
			t.Fatalf("expected %s rows preserved (got 0) — newSaveData inserts into this table", table)
		}
	}
}
