package handler

import "testing"

func TestVerifySignatureBypass(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "bypass-token")

	if !verifySignature("data=test", "bypass-token", []byte("secret")) {
		t.Fatalf("expected bypass token to allow signature verification")
	}
}

func TestVerifySignatureBypassEmptyToken(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")

	if verifySignature("data=test", "bypass-token", []byte("secret")) {
		t.Fatalf("expected bypass to be disabled when token is empty")
	}
}

func TestVerifyUserSignatureBypass(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "bypass-token")

	if !verifyUserSignature("user-1", "bypass-token") {
		t.Fatalf("expected bypass token to allow user signature verification")
	}
}

func TestVerifyUserSignatureBypassEmptyToken(t *testing.T) {
	t.Setenv("SIGNATURE_BYPASS_TOKEN", "")

	if verifyUserSignature("user-1", "bypass-token") {
		t.Fatalf("expected bypass to be disabled when token is empty")
	}
}
