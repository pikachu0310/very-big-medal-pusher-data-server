package handler

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/motoki317/sc"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// キャッシュ用エントリ数・TTL は必要に応じて調整
const rankingsCacheTTL = time.Minute

// v3統計データのキャッシュTTL
const statisticsCacheV3TTL = 30 * time.Minute

// v4統計データのキャッシュTTL
const statisticsCacheV4TTL = 5 * time.Minute

// 実績取得率キャッシュTTL
const achievementRatesCacheTTL = time.Hour

// メダル推移キャッシュTTL
const medalTimeseriesCacheTTL = time.Hour

// セーブアクティビティキャッシュTTL
const saveActivityCacheTTL = 10 * time.Minute

type Handler struct {
	repo                  Repository
	rankingCache          *sc.Cache[string, []models.GameData]
	totalMedalsCache      *sc.Cache[string, int]
	statisticsCacheV3     *sc.Cache[string, *models.StatisticsV3]
	statisticsCacheV4     *sc.Cache[string, *models.StatisticsV4]
	achievementRatesCache *sc.Cache[string, *models.AchievementRates]
	medalTimeseriesCache  *sc.Cache[string, *models.MedalTimeseriesResponse]
	saveActivityCache     *sc.Cache[string, *models.SaveActivityResponse]
}

type Repository interface {
	GetRankings(ctx context.Context, sortBy string, limit int) ([]models.GameData, error)
	GetTotalMedals(ctx context.Context) (int, error)
	GetStatisticsV4(ctx context.Context) (*models.StatisticsV4, error)
	GetAchievementRates(ctx context.Context) (*models.AchievementRates, error)
	GetMedalTimeseries(ctx context.Context, days int) (*models.MedalTimeseriesResponse, error)
	GetSaveActivity(ctx context.Context, hours int) (*models.SaveActivityResponse, error)
	GetCreditAllDistribution(ctx context.Context) (*models.CreditAllDistributionResponse, error)

	ExistsSameSave(ctx context.Context, userID string, playtime int64) (bool, error)
	InsertSaveV4(ctx context.Context, sd *domain.SaveData) error
	GetLatestSave(ctx context.Context, userID string) (*domain.SaveData, error)
	GetSaveHistory(ctx context.Context, userID string, limit int, before *time.Time) ([]models.SaveHistoryEntry, bool, error)
	GetAchievementUnlockHistory(ctx context.Context, userID string, limit int) ([]models.AchievementUnlockEntry, int, error)
}

func New(repo Repository) *Handler {
	h := &Handler{repo: repo}

	// ランキング全般キャッシュ (キー: "sortBy:limit")
	rankCache, err := sc.New(
		func(ctx context.Context, key string) ([]models.GameData, error) {
			parts := strings.Split(key, ":")
			sortBy := parts[0]
			limit, _ := strconv.Atoi(parts[1])
			return h.repo.GetRankings(ctx, sortBy, limit)
		},
		rankingsCacheTTL, rankingsCacheTTL,
		sc.WithLRUBackend(500),
	)
	if err != nil {
		log.Fatalf("failed to create ranking cache: %v", err)
	}
	h.rankingCache = rankCache

	// 全ユーザーのメダル合計キャッシュ (固定キー)
	medalCache, err := sc.New(
		func(ctx context.Context, _ string) (int, error) {
			return h.repo.GetTotalMedals(ctx)
		},
		rankingsCacheTTL, rankingsCacheTTL,
	)
	if err != nil {
		log.Fatalf("failed to create total medals cache: %v", err)
	}
	h.totalMedalsCache = medalCache

	if v3Repo, ok := repo.(interface {
		GetStatisticsV3(ctx context.Context) (*models.StatisticsV3, error)
	}); ok {
		// v3 統計データキャッシュの初期化
		statsCacheV3, err := sc.New(
			func(ctx context.Context, key string) (*models.StatisticsV3, error) {
				// key は使わないので無視
				return v3Repo.GetStatisticsV3(ctx)
			},
			statisticsCacheV3TTL, // freshFor: 5分
			statisticsCacheV3TTL, // ttl:      5分
			// 単一キーなのでバックエンドはデフォルトの map で十分
		)
		if err != nil {
			log.Fatalf("failed to create statistics v3 cache: %v", err)
		}
		h.statisticsCacheV3 = statsCacheV3
	}

	// v4 統計データキャッシュの初期化
	statsCacheV4, err := sc.New(
		func(ctx context.Context, key string) (*models.StatisticsV4, error) {
			// key は使わないので無視
			return repo.GetStatisticsV4(ctx)
		},
		statisticsCacheV4TTL, // freshFor: 5分
		statisticsCacheV4TTL, // ttl:      5分
		// 単一キーなのでバックエンドはデフォルトの map で十分
	)
	if err != nil {
		log.Fatalf("failed to create statistics v4 cache: %v", err)
	}
	h.statisticsCacheV4 = statsCacheV4

	// achievements rate キャッシュ
	achievementsCache, err := sc.New(
		func(ctx context.Context, key string) (*models.AchievementRates, error) {
			return h.repo.GetAchievementRates(ctx)
		},
		achievementRatesCacheTTL,
		achievementRatesCacheTTL,
	)
	if err != nil {
		log.Fatalf("failed to create achievement rates cache: %v", err)
	}
	h.achievementRatesCache = achievementsCache

	// メダル推移キャッシュ（日単位）
	medalTimeseriesCache, err := sc.New(
		func(ctx context.Context, key string) (*models.MedalTimeseriesResponse, error) {
			days, _ := strconv.Atoi(key)
			if days <= 0 {
				days = 30
			}
			return h.repo.GetMedalTimeseries(ctx, days)
		},
		medalTimeseriesCacheTTL,
		medalTimeseriesCacheTTL,
		sc.WithLRUBackend(32),
	)
	if err != nil {
		log.Fatalf("failed to create medal timeseries cache: %v", err)
	}
	h.medalTimeseriesCache = medalTimeseriesCache

	// セーブアクティビティキャッシュ（時間単位）
	saveActivityCache, err := sc.New(
		func(ctx context.Context, key string) (*models.SaveActivityResponse, error) {
			hours, _ := strconv.Atoi(key)
			if hours <= 0 {
				hours = 168
			}
			return h.repo.GetSaveActivity(ctx, hours)
		},
		saveActivityCacheTTL,
		saveActivityCacheTTL,
		sc.WithLRUBackend(32),
	)
	if err != nil {
		log.Fatalf("failed to create save activity cache: %v", err)
	}
	h.saveActivityCache = saveActivityCache

	return h
}
