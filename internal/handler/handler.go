package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (h *Handler) GetData(ctx echo.Context, params models.GetDataParams) error {
	// ユーザーごとの秘密鍵を生成
	userSecret := generateUserSecret(params.UserId)

	// クエリパラメータを署名用に整列
	paramStr := createSortedParamString(params)

	// 署名の検証
	if !verifySignature(paramStr, params.Sig, userSecret) {
		return ctx.JSON(http.StatusBadRequest, "invalid signature")
	}

	// DBにデータを保存
	data := repository.GameData{
		ID:           uuid.New(),
		UserID:       params.UserId,
		Version:      params.V,
		InMedal:      params.InMedal,
		OutMedal:     params.OutMedal,
		SlotHit:      params.SlotHit,
		GetShirbe:    params.GetShirbe,
		StartSlot:    params.StartSlot,
		ShirbeBuy300: params.ShirbeBuy300,
		Medal1:       params.Medal1,
		Medal2:       params.Medal2,
		Medal3:       params.Medal3,
		Medal4:       params.Medal4,
		Medal5:       params.Medal5,
		RMedal:       params.RMedal,
		Second:       params.Second,
		Minute:       params.Minute,
		Hour:         params.Hour,
	}

	if err := h.repo.InsertGameData(ctx.Request().Context(), data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}

func generateUserSecret(userID string) []byte {
	h := hmac.New(sha256.New, []byte(GlobalSecret))
	h.Write([]byte(userID))
	return h.Sum(nil)
}

func createSortedParamString(params models.GetDataParams) string {
	paramMap := map[string]string{
		"v":             params.V,
		"user_id":       params.UserId,
		"in_medal":      fmt.Sprintf("%d", params.InMedal),
		"out_medal":     fmt.Sprintf("%d", params.OutMedal),
		"slot_hit":      fmt.Sprintf("%d", params.SlotHit),
		"get_shirbe":    fmt.Sprintf("%d", params.GetShirbe),
		"start_slot":    fmt.Sprintf("%d", params.StartSlot),
		"shirbe_buy300": fmt.Sprintf("%d", params.ShirbeBuy300),
		"medal_1":       fmt.Sprintf("%d", params.Medal1),
		"medal_2":       fmt.Sprintf("%d", params.Medal2),
		"medal_3":       fmt.Sprintf("%d", params.Medal3),
		"medal_4":       fmt.Sprintf("%d", params.Medal4),
		"medal_5":       fmt.Sprintf("%d", params.Medal5),
		"R_medal":       fmt.Sprintf("%d", params.RMedal),
		"second":        fmt.Sprintf("%.2f", params.Second),
		"minute":        fmt.Sprintf("%d", params.Minute),
		"hour":          fmt.Sprintf("%d", params.Hour),
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
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedMAC), []byte(sig))
}
