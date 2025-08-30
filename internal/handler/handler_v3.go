package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
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
	log.Printf("[DEBUG] GetV3Data called with params: %+v", params)

	// クエリ文字列から signature 部分を切り出し
	rawQS := ctx.Request().URL.RawQuery
	log.Printf("[DEBUG] Raw query string: %s", rawQS)

	parts := strings.SplitN(rawQS, "&sig=", 2)
	if len(parts) != 2 {
		log.Printf("[DEBUG] Failed to split query string by &sig=, parts: %v", parts)
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	rawQueryPart := parts[0] // "data=…&user_id=%5B1%5D%20Local%20Player"
	log.Printf("[DEBUG] Raw query part (before sig): %s", rawQueryPart)

	// user_id をデコードして署名対象を再構築
	kv := strings.SplitN(rawQueryPart, "&user_id=", 2)
	if len(kv) != 2 {
		log.Printf("[DEBUG] Failed to split by &user_id=, kv: %v", kv)
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}
	dataPart, encodedUid := kv[0], kv[1]
	log.Printf("[DEBUG] Data part: %s, encoded user_id: %s", dataPart, encodedUid)

	uid, err := url.QueryUnescape(encodedUid)
	if err != nil {
		log.Printf("[DEBUG] Failed to unescape user_id: %v", err)
		return ctx.String(http.StatusBadRequest, "invalid user_id encoding")
	}
	log.Printf("[DEBUG] Decoded user_id: %s", uid)

	signingStr := dataPart + "&user_id=" + encodedUid
	log.Printf("[DEBUG] Signing string: %s", signingStr)

	// 署名検証
	userSecret := generateUserSecretV3(uid)
	log.Printf("[DEBUG] Generated user secret (hex): %x", userSecret)
	log.Printf("[DEBUG] Received signature: %s", params.Sig)

	// if !verifySignature(signingStr, params.Sig, generateUserSecretV3(uid)) {
	// 	return ctx.String(http.StatusUnauthorized, "invalid signature")
	// }

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
	log.Printf("[DEBUG] GetV3UsersUserIdData called with userId: %s, params: %+v", userId, params)

	// 1) 署名必須
	if params.Sig == "" {
		log.Printf("[DEBUG] Missing signature for userId: %s", userId)
		return ctx.String(http.StatusBadRequest, "missing signature")
	}

	log.Printf("[DEBUG] Verifying signature for userId: %s, sig: %s", userId, params.Sig)
	valid := verifyUserSignature(userId, params.Sig)
	log.Printf("[DEBUG] Signature verification result: %t", valid)

	if !valid {
		log.Printf("[DEBUG] Invalid signature for userId: %s", userId)
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
	log.Printf("[DEBUG] verifyUserSignature called with userID: %s, sig: %s", userID, sig)

	secret := []byte(config.GetSecretKeyLoadV2())
	log.Printf("[DEBUG] Load secret key (hex): %x", secret)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(userID))
	expected := hex.EncodeToString(mac.Sum(nil))

	log.Printf("[DEBUG] Expected signature: %s", expected)
	log.Printf("[DEBUG] Received signature: %s", sig)
	log.Printf("[DEBUG] Signatures match (case-insensitive): %t", strings.EqualFold(sig, expected))

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
	log.Printf("[DEBUG] generateUserSecretV3 called with userID: %s", userID)

	secretKey := config.GetSecretKeySaveV2()
	log.Printf("[DEBUG] Save secret key (hex): %x", []byte(secretKey))

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(userID))
	result := h.Sum(nil)

	log.Printf("[DEBUG] Generated user secret (hex): %x", result)
	return result
}

// GetV3AchievementsRates returns achievement acquisition rates
func (h *Handler) GetV3AchievementsRates(ctx echo.Context) error {
	rates, err := h.repo.GetAchievementRates(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, rates)
}
