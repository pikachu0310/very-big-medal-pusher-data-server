package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func decodeUserIDParam(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", nil
	}
	if decoded, ok := decodeBase64ID(trimmed); ok {
		return decoded, nil
	}
	decoded, err := url.QueryUnescape(trimmed)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

func decodeBase64ID(val string) (string, bool) {
	encs := []*base64.Encoding{
		base64.RawURLEncoding,
		base64.URLEncoding,
		base64.RawStdEncoding,
		base64.StdEncoding,
	}
	normalizedVal := normalizeBase64(val)
	for _, enc := range encs {
		decoded, err := enc.DecodeString(val)
		if err != nil {
			continue
		}
		reencoded := enc.EncodeToString(decoded)
		if normalizeBase64(reencoded) == normalizedVal {
			return string(decoded), true
		}
	}
	return "", false
}

func normalizeBase64(val string) string {
	val = strings.ReplaceAll(val, "-", "+")
	val = strings.ReplaceAll(val, "_", "/")
	return strings.TrimRight(val, "=")
}

func buildSignedSaveData(model *models.SaveDataV2) (*models.SignedSaveData, error) {
	payload, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(payload)
	sig := signPayload(encoded, []byte(config.GetSecretKeyLoadV2()))

	return &models.SignedSaveData{
		Data: encoded,
		Sig:  sig,
	}, nil
}

func signPayload(message string, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
