package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
)

func verifySignature(data, sig string, secret []byte) bool {
	if bypassSignature(sig) {
		return true
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(sig))
}

// 署名検証：sig == HMAC-SHA256( key=<LoadSecret>, msg=userID )
func verifyUserSignature(userID, sig string) bool {
	if bypassSignature(sig) {
		return true
	}
	secret := []byte(config.GetSecretKeyLoadV2())

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(userID))
	expected := hex.EncodeToString(mac.Sum(nil))

	return strings.EqualFold(sig, expected)
}

// verifyUserSignatureV4 tries both decoded and raw user_id to keep compatibility with
// clients that sign either representation.
func verifyUserSignatureV4(rawUserID, decodedUserID, sig string) bool {
	if decodedUserID != "" && verifyUserSignature(decodedUserID, sig) {
		return true
	}
	if rawUserID != "" && rawUserID != decodedUserID {
		return verifyUserSignature(rawUserID, sig)
	}
	return false
}

func bypassSignature(sig string) bool {
	token := config.GetSignatureBypassToken()
	if token == "" || sig == "" {
		return false
	}
	return hmac.Equal([]byte(sig), []byte(token))
}
