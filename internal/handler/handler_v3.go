package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

const (
	// 統計データのキャッシュキー（単一なので何でも良い）
	statisticsCacheV3Key = "v3_stats"
)

// func (h *Handler) GetV3Data(ctx echo.Context, params models.GetV3DataParams) error {
// 	fmt.Printf("[V3-DEBUG] GetV3Data START - user_id=%s, data_len=%d\n", params.UserId, len(params.Data))

// 	// クエリパラメータを正しい順序で構築して署名検証を行う
// 	// クエリパラメータの順序は重要: data, user_id, sig の順序で署名を生成する必要がある

// 	// 1. 署名対象となるクエリ文字列を構築（sigパラメータを除く）
// 	// クライアント側と同じエンコーディングを使用するため、手動でエンコーディング
// 	dataEncoded := strings.ReplaceAll(url.QueryEscape(params.Data), "+", "%20")
// 	userIdEncoded := strings.ReplaceAll(url.QueryEscape(params.UserId), "+", "%20")
// 	signingStr := "data=" + dataEncoded + "&user_id=" + userIdEncoded

// 	// 3. 署名検証
// 	if !verifySignature(signingStr, params.Sig, generateUserSecretV3(params.UserId)) {
// 		fmt.Printf("[V3-DEBUG] GetV3Data SIGNATURE_INVALID - user_id=%s\n", params.UserId)
// 		return ctx.String(http.StatusUnauthorized, "invalid signature")
// 	}

// 	// JSON 部分をパース
// 	fmt.Printf("[V3-DEBUG] GetV3Data PARSING_DATA - user_id=%s\n", params.UserId)
// 	sd, err := domain.ParseSaveData(params.Data)
// 	if err != nil {
// 		fmt.Printf("[V3-DEBUG] GetV3Data PARSE_ERROR - user_id=%s, error=%v\n", params.UserId, err)
// 		return ctx.String(http.StatusBadRequest, err.Error())
// 	}

// 	// 重複チェック
// 	fmt.Printf("[V3-DEBUG] GetV3Data CHECKING_DUPLICATE - user_id=%s\n", params.UserId)
// 	exists, err := h.repo.ExistsSameSave(ctx.Request().Context(), params.UserId, sd.Playtime)
// 	if err != nil {
// 		fmt.Printf("[V3-DEBUG] GetV3Data DUPLICATE_CHECK_ERROR - user_id=%s, error=%v\n", params.UserId, err)
// 		return ctx.String(http.StatusInternalServerError, err.Error())
// 	}
// 	if exists {
// 		fmt.Printf("[V3-DEBUG] GetV3Data DUPLICATE_FOUND - user_id=%s\n", params.UserId)
// 		return ctx.String(http.StatusConflict, "duplicate save data")
// 	}

// 	// 保存
// 	fmt.Printf("[V3-DEBUG] GetV3Data INSERTING_DATA - user_id=%s\n", params.UserId)
// 	sd.UserId = params.UserId
// 	if err := h.repo.InsertSave(ctx.Request().Context(), sd); err != nil {
// 		fmt.Printf("[V3-DEBUG] GetV3Data INSERT_ERROR - user_id=%s, error=%v\n", params.UserId, err)
// 		return ctx.String(http.StatusInternalServerError, err.Error())
// 	}

// 	fmt.Printf("[V3-DEBUG] GetV3Data SUCCESS - user_id=%s\n", params.UserId)
// 	return ctx.JSON(http.StatusOK, "success")
// }

// // -----------------------------------------------------------------
// // /v3/users/{id}/data  ハンドラ
// // -----------------------------------------------------------------
// func (h *Handler) GetV3UsersUserIdData(
// 	ctx echo.Context,
// 	userId string,
// 	params models.GetV3UsersUserIdDataParams, // ← Sig はここに入ってくる
// ) error {
// 	fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData START - user_id=%s\n", userId)

// 	// 1) 署名必須
// 	if params.Sig == "" {
// 		fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData MISSING_SIGNATURE - user_id=%s\n", userId)
// 		return ctx.String(http.StatusBadRequest, "missing signature")
// 	}
// 	if !verifyUserSignature(userId, params.Sig) {
// 		fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData INVALID_SIGNATURE - user_id=%s\n", userId)
// 		return ctx.String(http.StatusUnauthorized, "invalid signature")
// 	}

// 	// 2) 最新セーブデータ取得
// 	fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData FETCHING_DATA - user_id=%s\n", userId)
// 	sd, err := h.repo.GetLatestSave(ctx.Request().Context(), userId)
// 	if err != nil {
// 		fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData FETCH_ERROR - user_id=%s, error=%v\n", userId, err)
// 		return ctx.String(http.StatusNotFound, err.Error())
// 	}

// 	// 3) OpenAPI モデル化 & 返却
// 	fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData CONVERTING_TO_MODEL - user_id=%s\n", userId)
// 	model := sd.ToModel()
// 	fmt.Printf("[V3-DEBUG] GetV3UsersUserIdData SUCCESS - user_id=%s\n", userId)
// 	return ctx.JSON(http.StatusOK, model)
// }

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

// // GetV3Statistics returns combined rankings and total medals, with cache.
// func (h *Handler) GetV3Statistics(ctx echo.Context) error {
// 	fmt.Printf("[V3-DEBUG] GetV3Statistics START\n")

// 	// キャッシュから取得。キーは statisticsCacheKey を常に使用
// 	fmt.Printf("[V3-DEBUG] GetV3Statistics FETCHING_FROM_CACHE\n")
// 	stats, err := h.statisticsCacheV3.Get(ctx.Request().Context(), statisticsCacheV3Key)
// 	if err != nil {
// 		fmt.Printf("[V3-DEBUG] GetV3Statistics CACHE_ERROR - error=%v\n", err)
// 		return ctx.String(http.StatusInternalServerError, err.Error())
// 	}

// 	fmt.Printf("[V3-DEBUG] GetV3Statistics SUCCESS\n")
// 	return ctx.JSON(http.StatusOK, stats)
// }

// func generateUserSecretV3(userID string) []byte {
// 	h := hmac.New(sha256.New, []byte(config.GetSecretKeySaveV2()))
// 	h.Write([]byte(userID))
// 	return h.Sum(nil)
// }

// // GetV3AchievementsRates returns achievement acquisition rates
// func (h *Handler) GetV3AchievementsRates(ctx echo.Context) error {
// 	fmt.Printf("[V3-DEBUG] GetV3AchievementsRates START\n")

// 	fmt.Printf("[V3-DEBUG] GetV3AchievementsRates FETCHING_RATES\n")
// 	rates, err := h.repo.GetAchievementRates(ctx.Request().Context())
// 	if err != nil {
// 		fmt.Printf("[V3-DEBUG] GetV3AchievementsRates FETCH_ERROR - error=%v\n", err)
// 		return ctx.String(http.StatusInternalServerError, err.Error())
// 	}

// 	fmt.Printf("[V3-DEBUG] GetV3AchievementsRates SUCCESS\n")
// 	return ctx.JSON(http.StatusOK, rates)
// }

// GetV4Data は v4 エンドポイントでセーブデータを保存する
func (h *Handler) GetV3Data(ctx echo.Context, params models.GetV3DataParams) error {
	// クエリパラメータを正しい順序で構築して署名検証を行う
	dataEncoded := strings.ReplaceAll(url.QueryEscape(params.Data), "+", "%20")
	userIdEncoded := strings.ReplaceAll(url.QueryEscape(params.UserId), "+", "%20")
	signingStr := "data=" + dataEncoded + "&user_id=" + userIdEncoded

	// 署名検証
	if !verifySignature(signingStr, params.Sig, generateUserSecretV4(params.UserId)) {
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

	// 保存（v2_save_data に保存し、v3_user_latest_save_data を更新）
	sd.UserId = params.UserId
	if err := h.repo.InsertSaveV4(ctx.Request().Context(), sd); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}

// GetV4UsersUserIdData は v4 エンドポイントでユーザーの最新セーブデータを取得する
func (h *Handler) GetV3UsersUserIdData(
	ctx echo.Context,
	userId string,
	params models.GetV3UsersUserIdDataParams,
) error {
	// 署名必須
	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	if !verifyUserSignature(userId, params.Sig) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// 最新セーブデータ取得（v2_save_data を参照）
	sd, err := h.repo.GetLatestSave(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	// OpenAPI モデル化 & 返却
	model := sd.ToModel()
	return ctx.JSON(http.StatusOK, model)
}

// GetV4Statistics は v4 エンドポイントで最適化された統計データを返す
func (h *Handler) GetV3Statistics(ctx echo.Context) error {
	// キャッシュから取得
	stats, err := h.statisticsCacheV4.Get(ctx.Request().Context(), statisticsCacheV4Key)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, stats)
}

// GetV4AchievementsRates は v4 エンドポイントで実績取得率を返す
func (h *Handler) GetV3AchievementsRates(ctx echo.Context) error {
	rates, err := h.repo.GetAchievementRates(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, rates)
}
