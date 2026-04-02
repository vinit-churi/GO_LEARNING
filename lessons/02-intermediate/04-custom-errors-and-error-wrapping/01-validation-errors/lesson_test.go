package main

import (
	"errors"
	"testing"
)

func TestValidationErrorMessage(t *testing.T) {
	err := ValidationError{Field: "username", Problem: "cannot be blank"}
	if got := err.Error(); got != "username: cannot be blank" {
		t.Fatalf("Error() = %q", got)
	}
}

func TestValidateUsername(t *testing.T) {
	err := ValidateUsername("   ")
	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("ValidateUsername() error should unwrap to ValidationError: %v", err)
	}
}

func TestLoadUsernameWrapsValidationError(t *testing.T) {
	_, err := LoadUsername("   ")
	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("LoadUsername() error should unwrap to ValidationError: %v", err)
	}
	if validationErr.Field != "username" {
		t.Fatalf("wrapped ValidationError.Field = %q", validationErr.Field)
	}
}
