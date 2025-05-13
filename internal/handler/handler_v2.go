package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
	"net/http"
	"strings"
)

func (h *Handler) GetV2Data(ctx echo.Context, params models.GetV2DataParams) error {
	// grab the raw query string, e.g. "data=%7B...%7D&user_id=pikachu0310&sig=..."
	rawQS := ctx.Request().URL.RawQuery

	// split off the signature portion
	parts := strings.SplitN(rawQS, "&sig=", 2)
	if len(parts) != 2 {
		return ctx.String(http.StatusBadRequest, "missing sig")
	}
	queryStr := parts[0] // exactly "data=<encodedJson>&user_id=<rawPlayerId>"
	sig := params.Sig    // echo already bound the decoded sig

	// verify
	if !verifySignature(queryStr, sig, generateUserSecret(params.UserId)) {
		return ctx.String(http.StatusUnauthorized, "invalid signature")
	}

	// 1. パラメータ文字列を復元 (data と user_id)
	userID := params.UserId

	// 3. 重複チェック (playtime)
	sd, err := domain.ParseSaveData(params.Data)
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

func (h *Handler) GetV2Statistics(ctx echo.Context) error {
	stats, err := h.repo.GetStatistics(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, stats)
}
