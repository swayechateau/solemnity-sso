package crypt_test

import (
	"encoding/hex"
	"sso/pkg/crypt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	n := 10
	randomString, err := crypt.GenerateRandomString(n)
	if err != nil {
		t.Errorf("Error generating random string: %v", err)
	}

	decodedString, err := hex.DecodeString(randomString)
	if err != nil {
		t.Errorf("Error decoding random string: %v", err)
	}

	if len(decodedString) != n {
		t.Errorf("Generated random string length is incorrect. Expected: %d, Got: %d", n, len(decodedString))
	}
}
