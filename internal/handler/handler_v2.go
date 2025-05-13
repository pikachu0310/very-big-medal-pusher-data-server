package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// 統計データのキャッシュキー（単一なので何でも良い）
	statisticsCacheKey = "v2_stats"
	// キャッシュ TTL を 1 分に設定
	statisticsCacheTTL = time.Minute
)

func (h *Handler) GetV2Data(ctx echo.Context, params models.GetV2DataParams) error {
	// クエリ文字列から signature 部分を切り出し
	rawQS := ctx.Request().URL.RawQuery
	parts := strings.SplitN(rawQS, "&sig=", 2)
	if len(parts) != 2 {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	rawQueryPart := parts[0] // "data=…&user_id=%5B1%5D%20Local%20Player"

	// user_id をデコードして署名対象を再構築
	kv := strings.SplitN(rawQueryPart, "&user_id=", 2)
	if len(kv) != 2 {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}
	dataPart, encodedUid := kv[0], kv[1]
	uid, err := url.QueryUnescape(encodedUid)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id encoding")
	}
	signingStr := dataPart + "&user_id=" + encodedUid

	// 署名検証
	if !verifySignature(signingStr, params.Sig, generateUserSecret(uid)) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// JSON 部分をパース
	sd, err := domain.ParseSaveData(params.Data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// 重複チェック
	exists, err := h.repo.ExistsSameSave(ctx.Request().Context(), uid, sd.Playtime)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if exists {
		return ctx.String(http.StatusConflict, "duplicate save data")
	}

	// 保存
	sd.UserId = uid
	if err := h.repo.InsertSave(ctx.Request().Context(), sd); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}

func (h *Handler) GetV2UsersUserIdData(ctx echo.Context, userId string) error {
	// 1. 最新のセーブデータを取得
	sd, err := h.repo.GetLatestSave(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	// 2. domain→OpenAPIモデルに変換
	model := sd.ToModel()

	// 3. JSONで返却
	return ctx.JSON(http.StatusOK, model)
}

// GetV2Statistics returns combined rankings and total medals, with cache.
func (h *Handler) GetV2Statistics(ctx echo.Context) error {
	// キャッシュから取得。キーは statisticsCacheKey を常に使用
	stats, err := h.statisticsCache.Get(ctx.Request().Context(), statisticsCacheKey)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, stats)
}
