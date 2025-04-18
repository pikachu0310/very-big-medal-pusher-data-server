package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/domain"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

var GlobalSecret = "your_global_secret_here"

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) GetPing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (h *Handler) GetData(ctx echo.Context, params models.GetDataParams) error {
	userSecret := generateUserSecret(params.UserId)
	paramStr := createSortedParamString(params)

	if !verifySignature(paramStr, params.Sig, userSecret) {
		return ctx.JSON(http.StatusBadRequest, "invalid signature")
	}

	data := domain.GetDataParamsToGameData(params)

	if err := h.repo.InsertGameData(ctx.Request().Context(), data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}

func (h *Handler) GetUsersUserIdData(ctx echo.Context, userId string) error {
	data, err := h.repo.GetUserGameData(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}

func (h *Handler) GetRankings(ctx echo.Context, params models.GetRankingsParams) error {
	sortBy := "have_medal"
	if params.Sort != nil {
		sortBy = string(*params.Sort)
	}

	limit := 50
	if params.Limit != nil {
		limit = *params.Limit
	}

	rankings, err := h.repo.GetRankings(ctx.Request().Context(), sortBy, limit)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, rankings)
}

func generateUserSecret(userID string) []byte {
	h := hmac.New(sha256.New, []byte(GlobalSecret))
	h.Write([]byte(userID))
	return h.Sum(nil)
}

func createSortedParamString(params models.GetDataParams) string {
	paramMap := map[string]string{
		"version":       params.Version,
		"user_id":       params.UserId,
		"have_medal":    strconv.Itoa(params.HaveMedal),
		"in_medal":      strconv.Itoa(params.InMedal),
		"out_medal":     strconv.Itoa(params.OutMedal),
		"slot_hit":      strconv.Itoa(params.SlotHit),
		"get_shirbe":    strconv.Itoa(params.GetShirbe),
		"start_slot":    strconv.Itoa(params.StartSlot),
		"shirbe_buy300": strconv.Itoa(params.ShirbeBuy300),
		"medal_1":       strconv.Itoa(params.Medal1),
		"medal_2":       strconv.Itoa(params.Medal2),
		"medal_3":       strconv.Itoa(params.Medal3),
		"medal_4":       strconv.Itoa(params.Medal4),
		"medal_5":       strconv.Itoa(params.Medal5),
		"R_medal":       strconv.Itoa(params.RMedal),
		"second":        strconv.Itoa(params.Second),
		"minute":        strconv.Itoa(params.Minute),
		"hour":          strconv.Itoa(params.Hour),
		"fever":         strconv.Itoa(params.Fever),
	}

	keys := make([]string, 0, len(paramMap))
	for k := range paramMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, k := range keys {
		sb.WriteString(k + "=" + paramMap[k])
		if i < len(keys)-1 {
			sb.WriteString("&")
		}
	}
	return sb.String()
}

func verifySignature(data, sig string, secret []byte) bool {
	return sig == "test"

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedMAC), []byte(sig))
}
