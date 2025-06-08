package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// 統計データのキャッシュキー（単一なので何でも良い）
	statisticsCacheV3Key = "v3_stats"
	// キャッシュ TTL を 5 分に設定
	statisticsCacheV3TTL = 5 * time.Minute
)

func (h *Handler) GetV3Data(ctx echo.Context, params models.GetV3DataParams) error {
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
	if !verifySignature(signingStr, params.Sig, generateUserSecretV3(uid)) {
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

// -----------------------------------------------------------------
// /v3/users/{id}/data  ハンドラ
// -----------------------------------------------------------------
func (h *Handler) GetV3UsersUserIdData(
	ctx echo.Context,
	userId string,
	params models.GetV3UsersUserIdDataParams, // ← Sig はここに入ってくる
) error {
	// 1) 署名必須
	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	if !verifyUserSignature(userId, params.Sig) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// 2) 最新セーブデータ取得
	sd, err := h.repo.GetLatestSave(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	// 3) OpenAPI モデル化 & 返却
	return ctx.JSON(http.StatusOK, sd.ToModel())
}

// -----------------------------------------------------------------
// 署名検証：sig == HMAC-SHA256( key=<LoadSecret>, msg=userID )
// -----------------------------------------------------------------
func verifyUserSignature(userID, sig string) bool {
	secret := []byte(config.GetSecretKeyLoadV2())

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(userID))
	expected := hex.EncodeToString(mac.Sum(nil))

	return strings.EqualFold(sig, expected)
}

// GetV3Statistics returns combined rankings and total medals, with cache.
func (h *Handler) GetV3Statistics(ctx echo.Context) error {
	// キャッシュから取得。キーは statisticsCacheKey を常に使用
	stats, err := h.statisticsCacheV3.Get(ctx.Request().Context(), statisticsCacheV3Key)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, stats)
}

func generateUserSecretV3(userID string) []byte {
	h := hmac.New(sha256.New, []byte(config.GetSecretKeySaveV2()))
	h.Write([]byte(userID))
	return h.Sum(nil)
}
