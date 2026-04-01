//go:build solution
// +build solution

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseQuantity(input string) (int, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return 0, errors.New("quantity cannot be empty")
	}
	value, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, fmt.Errorf("invalid quantity %q", input)
	}
	return value, nil
}

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func BuildOrderMessage(name string, qty int) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name cannot be empty")
	}
	if qty <= 0 {
		return "", fmt.Errorf("quantity must be positive: %d", qty)
	}
	return fmt.Sprintf("%s ordered %d items", strings.TrimSpace(name), qty), nil
}

func main() {
	fmt.Println(ParseQuantity(" 12 "))
	fmt.Println(SafeDivide(10, 2))
	fmt.Println(BuildOrderMessage("Ava", 3))
}
