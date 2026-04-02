//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func SaveNote(path, text string) error {
	return os.WriteFile(path, []byte(text), 0o644)
}

func LoadNote(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func DuplicateNote(src, dst string) error {
	text, err := LoadNote(src)
	if err != nil {
		return err
	}
	return SaveNote(dst, text)
}

func main() {
	fmt.Println("use go test to exercise file operations")
}
