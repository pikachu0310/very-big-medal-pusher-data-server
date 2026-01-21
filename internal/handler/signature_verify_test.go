package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
)

func TestSignatureVerifyHMAC(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")

	secret := []byte("secret")
	data := "data=test"
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(data))
	expected := hex.EncodeToString(mac.Sum(nil))

	if !verifySignature(data, expected, secret) {
		t.Fatalf("expected valid signature to pass")
	}
	if verifySignature(data, "invalid", secret) {
		t.Fatalf("expected invalid signature to fail")
	}
}

func TestSignatureVerifyBypass(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "bypass-token")

	if !verifySignature("data=test", "bypass-token", []byte("secret")) {
		t.Fatalf("expected bypass token to allow signature verification")
	}
}

func TestSignatureVerifyBypassEmptyToken(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")

	if verifySignature("data=test", "bypass-token", []byte("secret")) {
		t.Fatalf("expected bypass to be disabled when token is empty")
	}
}

func TestUserSignatureVerifyHMAC(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")
	t.Setenv("LOAD", "load-secret")

	userID := "user-1"
	expected := signPayload(userID, []byte(config.GetSecretKeyLoadV2()))

	if !verifyUserSignature(userID, expected) {
		t.Fatalf("expected valid user signature to pass")
	}
	if verifyUserSignature(userID, "invalid") {
		t.Fatalf("expected invalid user signature to fail")
	}
}

func TestUserSignatureVerifyBypass(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "bypass-token")

	if !verifyUserSignature("user-1", "bypass-token") {
		t.Fatalf("expected bypass token to allow user signature verification")
	}
}

func TestUserSignatureVerifyBypassEmptyToken(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")

	if verifyUserSignature("user-1", "bypass-token") {
		t.Fatalf("expected bypass to be disabled when token is empty")
	}
}

func TestUserSignatureVerifyV4(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")
	t.Setenv("LOAD", "load-secret")

	decodedUserID := "user-1"
	rawUserID := "user-1%2Braw"
	expected := signPayload(decodedUserID, []byte(config.GetSecretKeyLoadV2()))

	if !verifyUserSignatureV4(rawUserID, decodedUserID, expected) {
		t.Fatalf("expected decoded user signature to pass")
	}
	if verifyUserSignatureV4(rawUserID, decodedUserID, "invalid") {
		t.Fatalf("expected invalid user signature to fail")
	}
}
