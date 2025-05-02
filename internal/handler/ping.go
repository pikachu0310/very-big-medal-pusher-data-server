package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
