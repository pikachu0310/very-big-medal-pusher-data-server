package repository

import (
	"context"
	"math"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// GetSaveHistory returns recent save snapshots for a user ordered by updated_at desc.
func (r *Repository) GetSaveHistory(ctx context.Context, userID string, limit int, before *time.Time) ([]models.SaveHistoryEntry, bool, error) {
	query := `
SELECT
  id AS save_id,
  version,
  playtime,
  credit_all,
  medal_get,
  ball_get,
  jack_totalmax_v2,
  ult_totalmax_v2,
  cpm_max,
  blackbox_total,
  sp_use,
  created_at,
  updated_at
FROM v2_save_data
WHERE user_id = ?
`
	args := []any{userID}
	if before != nil {
		query += " AND updated_at < ?"
		args = append(args, *before)
	}
	query += " ORDER BY updated_at DESC LIMIT ?"
	args = append(args, limit+1)

	var rows []struct {
		SaveID        int       `db:"save_id"`
		Version       int       `db:"version"`
		Playtime      int64     `db:"playtime"`
		CreditAll     int64     `db:"credit_all"`
		MedalGet      int64     `db:"medal_get"`
		BallGet       int64     `db:"ball_get"`
		JackTotalMax  int64     `db:"jack_totalmax_v2"`
		UltTotalMax   int64     `db:"ult_totalmax_v2"`
		CpMMax        float64   `db:"cpm_max"`
		BlackboxTotal int64     `db:"blackbox_total"`
		SpUse         int64     `db:"sp_use"`
		CreatedAt     time.Time `db:"created_at"`
		UpdatedAt     time.Time `db:"updated_at"`
	}

	if err := r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, false, err
	}

	hasMore := len(rows) > limit
	if hasMore {
		rows = rows[:limit]
	}

	entries := make([]models.SaveHistoryEntry, 0, len(rows))
	for _, row := range rows {
		playtime := int(row.Playtime)
		creditAll := int(row.CreditAll)
		medalGet := int(row.MedalGet)
		ballGet := int(row.BallGet)
		jackTotal := int(row.JackTotalMax)
		ultTotal := int(row.UltTotalMax)
		blackboxTotal := int(row.BlackboxTotal)
		spUse := int(row.SpUse)

		entry := models.SaveHistoryEntry{
			SaveId:         intPtr(row.SaveID),
			Version:        intPtr(row.Version),
			Playtime:       intPtr(playtime),
			CreditAll:      intPtr(creditAll),
			MedalGet:       intPtr(medalGet),
			BallGet:        intPtr(ballGet),
			JackTotalmaxV2: intPtr(jackTotal),
			UltTotalmaxV2:  intPtr(ultTotal),
			CpmMax:         float64Ptr(row.CpMMax),
			BlackboxTotal:  intPtr(blackboxTotal),
			SpUse:          intPtr(spUse),
			CreatedAt:      timePtr(row.CreatedAt),
			UpdatedAt:      timePtr(row.UpdatedAt),
		}
		entries = append(entries, entry)
	}

	return entries, hasMore, nil
}

// GetAchievementUnlockHistory returns unlock moments for a user.
func (r *Repository) GetAchievementUnlockHistory(ctx context.Context, userID string, limit int) ([]models.AchievementUnlockEntry, int, error) {
	const totalQuery = `
SELECT COUNT(*) 
FROM v2_save_data_achievements a
JOIN v2_save_data s ON a.save_id = s.id
WHERE s.user_id = ?
`
	var total int
	if err := r.db.GetContext(ctx, &total, totalQuery, userID); err != nil {
		return nil, 0, err
	}

	const dataQuery = `
SELECT
  a.achievement_id,
  s.updated_at AS unlocked_at,
  s.playtime,
  a.save_id
FROM v2_save_data_achievements a
JOIN v2_save_data s ON a.save_id = s.id
WHERE s.user_id = ?
ORDER BY s.updated_at ASC
LIMIT ?
`

	var rows []struct {
		AchievementID string    `db:"achievement_id"`
		UnlockedAt    time.Time `db:"unlocked_at"`
		Playtime      int64     `db:"playtime"`
		SaveID        int64     `db:"save_id"`
	}

	if err := r.db.SelectContext(ctx, &rows, dataQuery, userID, limit); err != nil {
		return nil, 0, err
	}

	entries := make([]models.AchievementUnlockEntry, 0, len(rows))
	for _, row := range rows {
		playtime := int(row.Playtime)
		saveID := int(row.SaveID)
		entry := models.AchievementUnlockEntry{
			AchievementId: stringPtr(row.AchievementID),
			UnlockedAt:    timePtr(row.UnlockedAt),
			Playtime:      intPtr(playtime),
			SaveId:        intPtr(saveID),
		}
		entries = append(entries, entry)
	}

	return entries, total, nil
}

// GetMedalTimeseries aggregates total medals by day using the latest save per user per day.
func (r *Repository) GetMedalTimeseries(ctx context.Context, days int) (*models.MedalTimeseriesResponse, error) {
	query := `
WITH latest_per_user_day AS (
  SELECT 
    user_id,
    DATE(updated_at) AS day,
    MAX(updated_at) AS latest_updated_at
  FROM v2_save_data
  WHERE updated_at >= DATE_SUB(CURRENT_DATE, INTERVAL ? DAY)
  GROUP BY user_id, DATE(updated_at)
)
SELECT
  day,
  COALESCE(SUM(sd.credit_all), 0) AS total_medals,
  COUNT(*) AS active_users,
  COALESCE(AVG(sd.playtime), 0) AS avg_playtime
FROM latest_per_user_day l
JOIN v2_save_data sd
  ON sd.user_id = l.user_id
  AND DATE(sd.updated_at) = l.day
  AND sd.updated_at = l.latest_updated_at
GROUP BY day
ORDER BY day ASC
`

	var rows []struct {
		Day         time.Time `db:"day"`
		TotalMedals int64     `db:"total_medals"`
		ActiveUsers int64     `db:"active_users"`
		AvgPlaytime float64   `db:"avg_playtime"`
	}

	if err := r.db.SelectContext(ctx, &rows, query, days); err != nil {
		return nil, err
	}

	buckets := make([]models.MedalTimeseriesBucket, 0, len(rows))
	for _, row := range rows {
		total := int(row.TotalMedals)
		users := int(row.ActiveUsers)
		avg := int(math.Round(row.AvgPlaytime))
		date := openapi_types.Date{Time: row.Day}

		buckets = append(buckets, models.MedalTimeseriesBucket{
			Date:        &date,
			TotalMedals: intPtr(total),
			ActiveUsers: intPtr(users),
			AvgPlaytime: intPtr(avg),
		})
	}

	return &models.MedalTimeseriesResponse{
		Buckets: &buckets,
	}, nil
}

// GetSaveActivity aggregates hourly save counts for a time window.
func (r *Repository) GetSaveActivity(ctx context.Context, hours int) (*models.SaveActivityResponse, error) {
	query := `
SELECT
  TIMESTAMP(DATE_FORMAT(updated_at, '%Y-%m-%d %H:00:00')) AS hour_start,
  COUNT(*) AS saves,
  COUNT(DISTINCT user_id) AS unique_users
FROM v2_save_data
WHERE updated_at >= DATE_SUB(UTC_TIMESTAMP(), INTERVAL ? HOUR)
GROUP BY hour_start
ORDER BY hour_start ASC
`

	var rows []struct {
		HourStart   time.Time `db:"hour_start"`
		Saves       int       `db:"saves"`
		UniqueUsers int       `db:"unique_users"`
	}

	if err := r.db.SelectContext(ctx, &rows, query, hours); err != nil {
		return nil, err
	}

	buckets := make([]models.SaveActivityBucket, 0, len(rows))
	for _, row := range rows {
		saves := row.Saves
		users := row.UniqueUsers
		buckets = append(buckets, models.SaveActivityBucket{
			HourStart:   timePtr(row.HourStart),
			Saves:       intPtr(saves),
			UniqueUsers: intPtr(users),
		})
	}

	return &models.SaveActivityResponse{
		Buckets: &buckets,
	}, nil
}

func intPtr(v int) *int {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}

func stringPtr(v string) *string {
	return &v
}

func timePtr(t time.Time) *time.Time {
	return &t
}
