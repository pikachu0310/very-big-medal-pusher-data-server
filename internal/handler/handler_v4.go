package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

const (
	// 統計データのキャッシュキー（v4用）
	statisticsCacheV4Key = "v4_stats"
)

// GetV4Data は v4 エンドポイントでセーブデータを保存する
func (h *Handler) GetV4Data(ctx echo.Context, params models.GetV4DataParams) error {
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
func (h *Handler) GetV4UsersUserIdData(
	ctx echo.Context,
	userId string,
	params models.GetV4UsersUserIdDataParams,
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
	rates, err := h.repo.GetAchievementRates(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, rates)
}

// generateUserSecretV4 は v4 用のユーザーシークレットを生成する
func generateUserSecretV4(userID string) []byte {
	h := hmac.New(sha256.New, []byte(config.GetSecretKeySaveV2()))
	h.Write([]byte(userID))
	return h.Sum(nil)
}
