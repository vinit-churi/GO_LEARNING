//go:build solution
// +build solution

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsePair(input string) (string, int, error) {
	left, right, ok := strings.Cut(input, "=")
	if !ok {
		return "", 0, fmt.Errorf("missing separator")
	}
	name := strings.TrimSpace(left)
	count, err := strconv.Atoi(strings.TrimSpace(right))
	if err != nil {
		return "", 0, err
	}
	return name, count, nil
}

func MustParseCount(input string) int {
	_, count, err := ParsePair(input)
	if err != nil {
		return 0
	}
	return count
}

func FormatPair(name string, count int) string {
	return fmt.Sprintf("%s=%d", name, count)
}

func main() {
	fmt.Println(FormatPair("jobs", MustParseCount("jobs=3")))
}
