package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func (h *Handler) GetV2Data(ctx echo.Context, params models.GetV2DataParams) error {
	// v2エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}

func (h *Handler) GetV2UsersUserIdData(ctx echo.Context, userId string) error {
	// v2エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}

// GetV2Statistics returns combined rankings and total medals, with cache.
func (h *Handler) GetV2Statistics(ctx echo.Context) error {
	// v2エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}
