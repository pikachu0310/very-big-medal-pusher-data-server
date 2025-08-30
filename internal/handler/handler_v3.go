package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

const (
	// 統計データのキャッシュキー（単一なので何でも良い）
	statisticsCacheV3Key = "v3_stats"
	// キャッシュ TTL を 5 分に設定
	statisticsCacheV3TTL = 5 * time.Minute
)

func (h *Handler) GetV3Data(ctx echo.Context, params models.GetV3DataParams) error {
	// クエリパラメータを正しい順序で構築して署名検証を行う
	// クエリパラメータの順序は重要: data, user_id, sig の順序で署名を生成する必要がある

	// 1. 署名対象となるクエリ文字列を構築（sigパラメータを除く）
	// クライアント側と同じエンコーディングを使用するため、手動でエンコーディング
	dataEncoded := strings.ReplaceAll(url.QueryEscape(params.Data), "+", "%20")
	userIdEncoded := strings.ReplaceAll(url.QueryEscape(params.UserId), "+", "%20")
	signingStr := "data=" + dataEncoded + "&user_id=" + userIdEncoded

	// 3. 署名検証
	if !verifySignature(signingStr, params.Sig, generateUserSecretV3(params.UserId)) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// JSON 部分をパース
	sd, err := domain.ParseSaveData(params.Data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// 重複チェック
	exists, err := h.repo.ExistsSameSave(ctx.Request().Context(), params.UserId, sd.Playtime)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if exists {
		return ctx.String(http.StatusConflict, "duplicate save data")
	}

	// 保存
	sd.UserId = params.UserId
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

// GetV3AchievementsRates returns achievement acquisition rates
func (h *Handler) GetV3AchievementsRates(ctx echo.Context) error {
	rates, err := h.repo.GetAchievementRates(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, rates)
}
