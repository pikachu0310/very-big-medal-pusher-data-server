package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func (h *Handler) GetPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (h *Handler) GetData(ctx echo.Context, params models.GetDataParams) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}

func (h *Handler) GetUsersUserIdData(ctx echo.Context, userId string) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}

func (h *Handler) GetRankings(ctx echo.Context, params models.GetRankingsParams) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}

// GetTotalMedals は全ユーザーの最新 have_medal 合計を返すエンドポイント
func (h *Handler) GetTotalMedals(ctx echo.Context) error {
	// v1エンドポイントはもう使われなくなりました
	return ctx.JSON(http.StatusGone, map[string]string{
		"error": "This endpoint is deprecated and no longer available. Please use v4 endpoints instead.",
		"code":  "DEPRECATED_ENDPOINT",
	})
}

func verifySignature(data, sig string, secret []byte) bool {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(sig))
}
