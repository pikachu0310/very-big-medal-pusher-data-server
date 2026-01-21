package handler

import (
	"crypto/hmac"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
)

func bypassSignature(sig string) bool {
	token := config.GetSignatureBypassToken()
	if token == "" || sig == "" {
		return false
	}
	return hmac.Equal([]byte(sig), []byte(token))
}
