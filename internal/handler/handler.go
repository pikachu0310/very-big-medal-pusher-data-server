package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

var GlobalSecret = "your_global_secret_here"

// キャッシュ用エントリ数・TTL は必要に応じて調整
const rankingsCacheTTL = time.Minute

type Handler struct {
	repo *repository.Repository

	// キャッシュ用ミューテックス
	cacheMu sync.RWMutex

	// 200件キャッシュ
	cacheMaxChainOrange200    []models.GameData
	cacheMaxChainOrange200At  time.Time
	cacheMaxChainRainbow200   []models.GameData
	cacheMaxChainRainbow200At time.Time

	// 500件キャッシュ
	cacheMaxChainOrange500    []models.GameData
	cacheMaxChainOrange500At  time.Time
	cacheMaxChainRainbow500   []models.GameData
	cacheMaxChainRainbow500At time.Time

	// ← ここにメダル総量キャッシュ用フィールドを追加
	cacheTotalMedals   int
	cacheTotalMedalsAt time.Time
}

func New(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (h *Handler) GetData(ctx echo.Context, params models.GetDataParams) error {
	log.Printf("Received params: %+v", params)
	userSecret := generateUserSecret(params.UserId)
	paramStr := createSortedParamString(params)
	log.Printf("Generated param string: %s", paramStr)

	if !verifySignature(paramStr, params.Sig, userSecret) {
		return ctx.String(http.StatusBadRequest, "invalid signature")
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

		// ポインタ型かつ要素型が int で、かつ現在 nil なら…
		if fv.Kind() == reflect.Ptr &&
			fv.Type().Elem().Kind() == reflect.Int &&
			fv.IsNil() {

			// reflect.New(要素の型) で *int の Value を作り、
			// そのアドレスはすべて 0 初期化されている
			newPtr := reflect.New(fv.Type().Elem())
			fv.Set(newPtr)
		}
	}
}

func (h *Handler) GetUsersUserIdData(ctx echo.Context, userId string) error {
	data, err := h.repo.GetUserGameData(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h *Handler) GetRankings(ctx echo.Context, params models.GetRankingsParams) error {
	sortBy := "have_medal"
	if params.Sort != nil {
		sortBy = string(*params.Sort)
	}
	limit := 50
	if params.Limit != nil {
		limit = *params.Limit
	}

	// キャッシュ判定
	isOrange200 := sortBy == "max_chain_orange" && limit == 200
	isOrange500 := sortBy == "max_chain_orange" && limit == 500
	isRainbow200 := sortBy == "max_chain_rainbow" && limit == 200
	isRainbow500 := sortBy == "max_chain_rainbow" && limit == 500

	// 200件キャッシュ応答
	if isOrange200 {
		h.cacheMu.RLock()
		if time.Since(h.cacheMaxChainOrange200At) < rankingsCacheTTL {
			data := h.cacheMaxChainOrange200
			h.cacheMu.RUnlock()
			return ctx.JSON(http.StatusOK, data)
		}
		h.cacheMu.RUnlock()
	}
	if isRainbow200 {
		h.cacheMu.RLock()
		if time.Since(h.cacheMaxChainRainbow200At) < rankingsCacheTTL {
			data := h.cacheMaxChainRainbow200
			h.cacheMu.RUnlock()
			return ctx.JSON(http.StatusOK, data)
		}
		h.cacheMu.RUnlock()
	}

	// 500件キャッシュ応答
	if isOrange500 {
		h.cacheMu.RLock()
		if time.Since(h.cacheMaxChainOrange500At) < rankingsCacheTTL {
			data := h.cacheMaxChainOrange500
			h.cacheMu.RUnlock()
			return ctx.JSON(http.StatusOK, data)
		}
		h.cacheMu.RUnlock()
	}
	if isRainbow500 {
		h.cacheMu.RLock()
		if time.Since(h.cacheMaxChainRainbow500At) < rankingsCacheTTL {
			data := h.cacheMaxChainRainbow500
			h.cacheMu.RUnlock()
			return ctx.JSON(http.StatusOK, data)
		}
		h.cacheMu.RUnlock()
	}

	// DB取得
	rankings, err := h.repo.GetRankings(ctx.Request().Context(), sortBy, limit)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// キャッシュ更新
	if isOrange200 {
		h.cacheMu.Lock()
		h.cacheMaxChainOrange200 = rankings
		h.cacheMaxChainOrange200At = time.Now()
		h.cacheMu.Unlock()
	}
	if isRainbow200 {
		h.cacheMu.Lock()
		h.cacheMaxChainRainbow200 = rankings
		h.cacheMaxChainRainbow200At = time.Now()
		h.cacheMu.Unlock()
	}
	if isOrange500 {
		h.cacheMu.Lock()
		h.cacheMaxChainOrange500 = rankings
		h.cacheMaxChainOrange500At = time.Now()
		h.cacheMu.Unlock()
	}
	if isRainbow500 {
		h.cacheMu.Lock()
		h.cacheMaxChainRainbow500 = rankings
		h.cacheMaxChainRainbow500At = time.Now()
		h.cacheMu.Unlock()
	}

	return ctx.JSON(http.StatusOK, rankings)
}

// GetTotalMedals は全ユーザーの最新 have_medal 合計を返すエンドポイント
func (h *Handler) GetTotalMedals(ctx echo.Context) error {
	// キャッシュ判定
	h.cacheMu.RLock()
	if time.Since(h.cacheTotalMedalsAt) < rankingsCacheTTL {
		total := h.cacheTotalMedals
		h.cacheMu.RUnlock()
		return ctx.JSON(http.StatusOK, map[string]int{"total_medals": total})
	}
	h.cacheMu.RUnlock()

	// キャッシュ切れ or 未初期化
	total, err := h.repo.GetTotalMedals(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// キャッシュ更新
	h.cacheMu.Lock()
	h.cacheTotalMedals = total
	h.cacheTotalMedalsAt = time.Now()
	h.cacheMu.Unlock()

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
			// ここで必要な型を追加
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
	//return sig == "test"

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedMAC), []byte(sig))
}
