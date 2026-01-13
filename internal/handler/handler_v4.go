package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

const (
	// 統計データのキャッシュキー（v4用）
	statisticsCacheV4Key = "v4_stats"

	// 実績取得率のキャッシュキー
	achievementRatesCacheKey = "v4_achievements_rates"
)

// GetV4Data は v4 エンドポイントでセーブデータを保存する
func (h *Handler) GetV4Data(ctx echo.Context, params models.GetV4DataParams) error {
	userID, err := decodeUserIDParam(params.UserId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id")
	}
	if userID == "" {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}

	// クエリパラメータを正しい順序で構築して署名検証を行う
	dataEncoded := strings.ReplaceAll(url.QueryEscape(params.Data), "+", "%20")
	userIdEncoded := strings.ReplaceAll(url.QueryEscape(params.UserId), "+", "%20")
	signingStr := "data=" + dataEncoded + "&user_id=" + userIdEncoded

	// 署名検証
	if !verifySignature(signingStr, params.Sig, generateUserSecretV4(userID)) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// JSON 部分をパース
	sd, err := domain.ParseSaveData(params.Data)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// 重複チェック
	exists, err := h.repo.ExistsSameSave(ctx.Request().Context(), userID, sd.Playtime)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if exists {
		return ctx.String(http.StatusConflict, "duplicate save data")
	}

	// 保存（v2_save_data に保存し、v3_user_latest_save_data を更新）
	sd.UserId = userID
	if err := h.repo.InsertSaveV4(ctx.Request().Context(), sd); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}

// GetV4DataVerify は v4 セーブデータの署名検証のみを行う
func (h *Handler) GetV4DataVerify(ctx echo.Context, params models.GetV4DataVerifyParams) error {
	userID, err := decodeUserIDParam(params.UserId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id")
	}
	if userID == "" {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}
	if params.Data == "" {
		return ctx.String(http.StatusBadRequest, "missing data")
	}
	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}

	dataEncoded := strings.ReplaceAll(url.QueryEscape(params.Data), "+", "%20")
	userIdEncoded := strings.ReplaceAll(url.QueryEscape(params.UserId), "+", "%20")
	signingStr := "data=" + dataEncoded + "&user_id=" + userIdEncoded

	if !verifySignature(signingStr, params.Sig, generateUserSecretV4(userID)) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	return ctx.JSON(http.StatusOK, models.SignatureVerifyResponse{Valid: true})
}

// GetV4UsersUserIdData は v4 エンドポイントでユーザーの最新セーブデータを取得する
func (h *Handler) GetV4UsersUserIdData(
	ctx echo.Context,
	userId string,
	params models.GetV4UsersUserIdDataParams,
) error {
	decodedUserID, err := decodeUserIDParam(userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id")
	}
	if decodedUserID == "" {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}

	// 署名必須
	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	if !verifyUserSignatureV4(userId, decodedUserID, params.Sig) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// 最新セーブデータ取得（v2_save_data を参照）
	sd, err := h.repo.GetLatestSave(ctx.Request().Context(), decodedUserID)
	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	// OpenAPI モデル化 & 返却
	model := sd.ToModel()
	signed, err := buildSignedSaveData(model)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, signed)
}

// GetV4UsersUserIdDataVerify はロード用の署名検証のみを行う
func (h *Handler) GetV4UsersUserIdDataVerify(
	ctx echo.Context,
	userId string,
	params models.GetV4UsersUserIdDataVerifyParams,
) error {
	decodedUserID, err := decodeUserIDParam(userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id")
	}
	if decodedUserID == "" {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}

	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	if !verifyUserSignatureV4(userId, decodedUserID, params.Sig) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	return ctx.JSON(http.StatusOK, models.SignatureVerifyResponse{Valid: true})
}

// GetV4Statistics は v4 エンドポイントで最適化された統計データを返す
func (h *Handler) GetV4Statistics(ctx echo.Context) error {
	// キャッシュから取得
	stats, err := h.statisticsCacheV4.Get(ctx.Request().Context(), statisticsCacheV4Key)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, stats)
}

// GetV4AchievementsRates は v4 エンドポイントで実績取得率を返す
func (h *Handler) GetV4AchievementsRates(ctx echo.Context) error {
	rates, err := h.achievementRatesCache.Get(ctx.Request().Context(), achievementRatesCacheKey)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, rates)
}

// GetV4UsersUserIdSaves は v4 エンドポイントでユーザーのセーブ履歴を返す
func (h *Handler) GetV4UsersUserIdSaves(
	ctx echo.Context,
	userId string,
	params models.GetV4UsersUserIdSavesParams,
) error {
	decodedUserID, err := decodeUserIDParam(userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id")
	}
	if decodedUserID == "" {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}

	// 署名必須
	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	if !verifyUserSignatureV4(userId, decodedUserID, params.Sig) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// limit/before の整形
	limit := 20
	if params.Limit != nil {
		limit = *params.Limit
	}
	if limit < 1 {
		limit = 1
	}
	if limit > 100 {
		limit = 100
	}

	entries, hasMore, err := h.repo.GetSaveHistory(ctx.Request().Context(), decodedUserID, limit, params.Before)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp := models.SaveHistoryResponse{
		Items: &entries,
	}
	if hasMore && len(entries) > 0 && entries[len(entries)-1].UpdatedAt != nil {
		resp.NextBefore = entries[len(entries)-1].UpdatedAt
	}

	return ctx.JSON(http.StatusOK, resp)
}

// GetV4UsersUserIdAchievementsHistory は v4 エンドポイントでアチーブメント履歴を返す
func (h *Handler) GetV4UsersUserIdAchievementsHistory(
	ctx echo.Context,
	userId string,
	params models.GetV4UsersUserIdAchievementsHistoryParams,
) error {
	decodedUserID, err := decodeUserIDParam(userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "invalid user_id")
	}
	if decodedUserID == "" {
		return ctx.String(http.StatusBadRequest, "missing user_id")
	}

	if params.Sig == "" {
		return ctx.String(http.StatusBadRequest, "missing signature")
	}
	if !verifyUserSignatureV4(userId, decodedUserID, params.Sig) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	limit := 500
	if params.Limit != nil {
		limit = *params.Limit
	}
	if limit < 1 {
		limit = 1
	}
	if limit > 2000 {
		limit = 2000
	}

	entries, total, err := h.repo.GetAchievementUnlockHistory(ctx.Request().Context(), decodedUserID, limit)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp := models.AchievementUnlockHistoryResponse{
		Items: &entries,
	}
	if total >= 0 {
		resp.Total = &total
	}

	return ctx.JSON(http.StatusOK, resp)
}

// GetV4StatisticsMedalsTimeseries は世界のメダル総量推移を返す
func (h *Handler) GetV4StatisticsMedalsTimeseries(ctx echo.Context, params models.GetV4StatisticsMedalsTimeseriesParams) error {
	days := 30
	if params.Days != nil {
		days = *params.Days
	}
	if days < 1 {
		days = 1
	}
	if days > 180 {
		days = 180
	}

	resp, err := h.medalTimeseriesCache.Get(ctx.Request().Context(), strconv.Itoa(days))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

// GetV4StatisticsSavesActivity はセーブ投稿の時間別推移を返す
func (h *Handler) GetV4StatisticsSavesActivity(ctx echo.Context, params models.GetV4StatisticsSavesActivityParams) error {
	hours := 168
	if params.Hours != nil {
		hours = *params.Hours
	}
	if hours < 1 {
		hours = 1
	}
	if hours > 720 {
		hours = 720
	}

	resp, err := h.saveActivityCache.Get(ctx.Request().Context(), strconv.Itoa(hours))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

// generateUserSecretV4 は v4 用のユーザーシークレットを生成する
func generateUserSecretV4(userID string) []byte {
	h := hmac.New(sha256.New, []byte(config.GetSecretKeySaveV2()))
	h.Write([]byte(userID))
	return h.Sum(nil)
}

// verifyUserSignatureV4 tries both decoded and raw user_id to keep compatibility with
// clients that sign either representation.
func verifyUserSignatureV4(rawUserID, decodedUserID, sig string) bool {
	if decodedUserID != "" && verifyUserSignature(decodedUserID, sig) {
		return true
	}
	if rawUserID != "" && rawUserID != decodedUserID {
		return verifyUserSignature(rawUserID, sig)
	}
	return false
}
