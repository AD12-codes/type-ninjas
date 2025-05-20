package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGenerateHashPassword(t *testing.T) {
	password := "mock-password"

	hashed, err := GenerateHashedPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(hashed) == 0 {
		t.Fatal("expected non-empty hashed password")
	}

	err = bcrypt.CompareHashAndPassword(hashed, []byte(password))

	if err != nil {
		t.Fatalf("hashed password did not match original password: %v", err)
	}
}
