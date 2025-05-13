package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
	"net/http"
	"net/url"
)

func (h *Handler) GetV2Data(ctx echo.Context, params models.GetV2DataParams) error {
	// 1. パラメータ文字列を復元 (data と user_id)
	raw := params.Data
	userID := params.UserId
	sig := params.Sig
	// 2. 署名検証
	// ソートされたクエリ文字列: data=...&user_id=...
	paramStr := "data=" + raw + "&user_id=" + url.QueryEscape(userID)
	if !verifySignature(paramStr, sig, generateUserSecret(userID)) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}
	// 3. 重複チェック (playtime)
	sd, err := domain.ParseSaveData(raw)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	exists, err := h.repo.ExistsSameSave(ctx.Request().Context(), userID, sd.Playtime)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if exists {
		return ctx.String(http.StatusConflict, "duplicate save data")
	}
	// 4. 保存
	sd.UserId = userID
	if err := h.repo.InsertSave(ctx.Request().Context(), sd); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "success")
}

func (h *Handler) GetV2UsersUserIdData(ctx echo.Context, userId string) error {
	// 最新のセーブデータを取得
	sd, err := h.repo.GetLatestSave(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}
	// モデルそのまま返却
	return ctx.JSON(http.StatusOK, sd)
}

func (h *Handler) GetV2Statistics(ctx echo.Context) error {
	stats, err := h.repo.GetStatistics(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, stats)
}
