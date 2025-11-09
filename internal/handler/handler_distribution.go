package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCreditAllDistribution returns aggregated counts of credit_all ranges.
func (h *Handler) GetCreditAllDistribution(ctx echo.Context) error {
	resp, err := h.repo.GetCreditAllDistribution(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}
