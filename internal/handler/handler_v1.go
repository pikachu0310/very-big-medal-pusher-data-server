package handler

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/motoki317/sc"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

var GlobalSecret = "your_global_secret_here"

// キャッシュ用エントリ数・TTL は必要に応じて調整
const rankingsCacheTTL = time.Minute

type Handler struct {
	repo              *repository.Repository
	rankingCache      *sc.Cache[string, []models.GameData]
	totalMedalsCache  *sc.Cache[string, int]
	statisticsCacheV3 *sc.Cache[string, *models.StatisticsV3]
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

	return h
}

func (h *Handler) GetPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (h *Handler) GetData(ctx echo.Context, params models.GetDataParams) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v3 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})

	log.Printf("Received params: %+v", params)
	userSecret := generateUserSecret(params.UserId)
	paramStr := createSortedParamString(params)
	log.Printf("Generated param string: %s", paramStr)

	if !verifySignature(paramStr, params.Sig, userSecret) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	exist, err := h.repo.ExistsSameGameData(ctx.Request().Context(), params.UserId, params.TotalPlayTime)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if exist {
		return ctx.String(http.StatusConflict, "Same data already exists! (You need to replace new Save URL)")
	}

	nullifyNullValues(&params)
	data := domain.GetDataParamsToGameData(params)

	if err := h.repo.InsertGameData(ctx.Request().Context(), data); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}

// nullifyNullValues は params の中で `*int` 型のフィールドをすべて見て
// nil のものを新しい int ポインタ(&0)で埋めます。
func nullifyNullValues(params *models.GetDataParams) {
	// params はポインタなので Elem() で構造体本体へ
	v := reflect.ValueOf(params).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		if fv.Kind() == reflect.Ptr && fv.Elem().Kind() == reflect.Int && fv.IsNil() {
			newPtr := reflect.New(fv.Type().Elem())
			fv.Set(newPtr)
		}
	}
}

func (h *Handler) GetUsersUserIdData(ctx echo.Context, userId string) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v3 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})

	data, err := h.repo.GetUserGameData(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h *Handler) GetRankings(ctx echo.Context, params models.GetRankingsParams) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v3 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})

	sortBy := "have_medal"
	if params.Sort != nil {
		sortBy = string(*params.Sort)
	}
	limit := 50
	if params.Limit != nil {
		limit = *params.Limit
	}

	// キャッシュキーを生成
	key := fmt.Sprintf("%s:%d", sortBy, limit)
	raw, err := h.rankingCache.Get(ctx.Request().Context(), key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// ソート種別に応じて変換
	switch sortBy {
	case "max_chain_orange":
		resp := domain.GetDatasToRankingResponseMaxChainOrange(raw)
		return ctx.JSON(http.StatusOK, resp)
	case "max_chain_rainbow":
		resp := domain.GetDatasToRankingResponseMaxChainRainbow(raw)
		return ctx.JSON(http.StatusOK, resp)
	case "max_total_jackpot":
		resp := domain.GetDatasToRankingResponseMaxTotalJackpot(raw)
		return ctx.JSON(http.StatusOK, resp)
	default:
		// その他は生の GameData
		return ctx.JSON(http.StatusOK, raw)
	}
}

// GetTotalMedals は全ユーザーの最新 have_medal 合計を返すエンドポイント
func (h *Handler) GetTotalMedals(ctx echo.Context) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v3 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})

	total, err := h.totalMedalsCache.Get(ctx.Request().Context(), "total")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]int{"total_medals": total})
}

func generateUserSecret(userID string) []byte {
	h := hmac.New(sha256.New, []byte(GlobalSecret))
	h.Write([]byte(userID))
	return h.Sum(nil)
}

// createSortedParamString は GetDataParams の struct タグ (form) を見て
// nil ポインタはスキップし、それ以外を key=value 形式でソート結合します。
func createSortedParamString(params models.GetDataParams) string {
	v := reflect.ValueOf(params)
	t := v.Type()

	// form タグ名 → 値 のマップ
	paramMap := make(map[string]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// form タグを取得し、",omitempty" 等を取り除く
		tag := field.Tag.Get("form")
		if tag == "" || tag == "-" {
			continue
		}
		key := strings.Split(tag, ",")[0]
		if key == "sig" {
			// signature は含めない
			continue
		}
		fv := v.Field(i)
		// ポインタは nil チェックして、中身を取り出す
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}
		var strVal string
		switch fv.Kind() {
		case reflect.String:
			strVal = fv.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strVal = strconv.FormatInt(fv.Int(), 10)
		default:
			continue
		}
		paramMap[key] = strVal
	}

	// キーをソート（大文字小文字はタグで揃っている前提なので ToLower で十分）
	keys := make([]string, 0, len(paramMap))
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	// key=value&key2=value2 ... に結合
	var sb strings.Builder
	for i, k := range keys {
		sb.WriteString(k)
		sb.WriteByte('=')
		sb.WriteString(paramMap[k])
		if i < len(keys)-1 {
			sb.WriteByte('&')
		}
	}
	return sb.String()
}

func verifySignature(data, sig string, secret []byte) bool {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(sig))
}
