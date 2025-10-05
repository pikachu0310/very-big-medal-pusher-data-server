package handler

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/motoki317/sc"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// キャッシュ用エントリ数・TTL は必要に応じて調整
const rankingsCacheTTL = time.Minute

// v3統計データのキャッシュTTL
const statisticsCacheV3TTL = 30 * time.Minute

// v4統計データのキャッシュTTL
const statisticsCacheV4TTL = 5 * time.Minute

type Handler struct {
	repo              *repository.Repository
	rankingCache      *sc.Cache[string, []models.GameData]
	totalMedalsCache  *sc.Cache[string, int]
	statisticsCacheV3 *sc.Cache[string, *models.StatisticsV3]
	statisticsCacheV4 *sc.Cache[string, *models.StatisticsV4]
}

func New(repo *repository.Repository) *Handler {
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

	// v3 統計データキャッシュの初期化
	statsCacheV3, err := sc.New(
		func(ctx context.Context, key string) (*models.StatisticsV3, error) {
			// key は使わないので無視
			return repo.GetStatisticsV3(ctx)
		},
		statisticsCacheV3TTL, // freshFor: 5分
		statisticsCacheV3TTL, // ttl:      5分
		// 単一キーなのでバックエンドはデフォルトの map で十分
	)
	if err != nil {
		log.Fatalf("failed to create statistics v3 cache: %v", err)
	}
	h.statisticsCacheV3 = statsCacheV3

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

	return h
}
