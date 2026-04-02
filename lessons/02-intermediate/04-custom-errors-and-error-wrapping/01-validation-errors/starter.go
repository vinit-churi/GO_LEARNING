//go:build !solution
// +build !solution

package main

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Problem string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Problem)
}

func ValidateUsername(name string) error {
	if strings.TrimSpace(name) == "" {
		return ValidationError{Field: "username", Problem: "cannot be blank"}
	}
	return nil
}

func LoadUsername(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if err := ValidateUsername(trimmed); err != nil {
		return "", fmt.Errorf("load username: %w", err)
	}
	return trimmed, nil
}

func main() {
	name, err := LoadUsername(" Mina ")
	fmt.Println(name, errors.Unwrap(err))
}
