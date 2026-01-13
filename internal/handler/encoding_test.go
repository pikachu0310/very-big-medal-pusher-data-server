package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

func TestDecodeUserIDParam(t *testing.T) {
	encoded := base64.RawURLEncoding.EncodeToString([]byte("user-1"))
	decoded, err := decodeUserIDParam(encoded)
	if err != nil {
		t.Fatalf("decodeUserIDParam error: %v", err)
	}
	if decoded != "user-1" {
		t.Fatalf("decoded: got %q", decoded)
	}

	urlDecoded, err := decodeUserIDParam("user%201")
	if err != nil {
		t.Fatalf("decodeUserIDParam error: %v", err)
	}
	if urlDecoded != "user 1" {
		t.Fatalf("url decoded: got %q", urlDecoded)
	}

	if _, err := decodeUserIDParam("%ZZ"); err == nil {
		t.Fatalf("decodeUserIDParam expected error for invalid escape")
	}
}

func TestDecodeBase64ID(t *testing.T) {
	raw := base64.RawURLEncoding.EncodeToString([]byte("user-1"))
	decoded, ok := decodeBase64ID(raw)
	if !ok || decoded != "user-1" {
		t.Fatalf("decodeBase64ID raw failed: %v %q", ok, decoded)
	}

	std := base64.StdEncoding.EncodeToString([]byte("user-2"))
	decoded, ok = decodeBase64ID(std)
	if !ok || decoded != "user-2" {
		t.Fatalf("decodeBase64ID std failed: %v %q", ok, decoded)
	}
}

func TestNormalizeBase64(t *testing.T) {
	got := normalizeBase64("a-b_c==")
	if got != "a+b/c" {
		t.Fatalf("normalizeBase64 got %q", got)
	}
}

func TestBuildSignedSaveData(t *testing.T) {
	t.Setenv("LOAD", "load-secret")
	legacy := 1
	version := 4
	model := &models.SaveDataV2{
		Legacy:  &legacy,
		Version: &version,
	}

	signed, err := buildSignedSaveData(model)
	if err != nil {
		t.Fatalf("buildSignedSaveData error: %v", err)
	}

	expectedSig := signPayload(signed.Data, []byte("load-secret"))
	if signed.Sig != expectedSig {
		t.Fatalf("signature mismatch: got %q", signed.Sig)
	}

	decoded, err := base64.StdEncoding.DecodeString(signed.Data)
	if err != nil {
		t.Fatalf("decode data: %v", err)
	}

	var decodedModel models.SaveDataV2
	if err := json.Unmarshal(decoded, &decodedModel); err != nil {
		t.Fatalf("unmarshal decoded: %v", err)
	}
	if decodedModel.Version == nil || *decodedModel.Version != 4 {
		t.Fatalf("decoded version: got %#v", decodedModel.Version)
	}
}

func TestVerifySignature(t *testing.T) {
	secret := []byte("secret")
	message := "data=test"

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(message))
	expected := hex.EncodeToString(mac.Sum(nil))

	if !verifySignature(message, expected, secret) {
		t.Fatalf("verifySignature expected true")
	}
	if verifySignature(message, "bad", secret) {
		t.Fatalf("verifySignature expected false")
	}
}

func TestVerifyUserSignature(t *testing.T) {
	t.Setenv("LOAD", "load-secret")
	userID := "user-1"
	expected := signPayload(userID, []byte("load-secret"))
	if !verifyUserSignature(userID, expected) {
		t.Fatalf("verifyUserSignature expected true")
	}
	if verifyUserSignature(userID, "bad") {
		t.Fatalf("verifyUserSignature expected false")
	}
}

func TestVerifyUserSignatureV4(t *testing.T) {
	t.Setenv("LOAD", "load-secret")
	rawUserID := "user%201"
	decodedUserID := "user 1"
	expected := signPayload(decodedUserID, []byte("load-secret"))

	if !verifyUserSignatureV4(rawUserID, decodedUserID, expected) {
		t.Fatalf("verifyUserSignatureV4 expected true for decoded signature")
	}
	if verifyUserSignatureV4(rawUserID, decodedUserID, "bad") {
		t.Fatalf("verifyUserSignatureV4 expected false for invalid signature")
	}
}

func TestGenerateUserSecretV4(t *testing.T) {
	t.Setenv("SAVE", "save-secret")
	userID := "user-1"
	got := generateUserSecretV4(userID)

	mac := hmac.New(sha256.New, []byte("save-secret"))
	mac.Write([]byte(userID))
	expected := mac.Sum(nil)

	if !hmac.Equal(got, expected) {
		t.Fatalf("generateUserSecretV4 mismatch")
	}
}
