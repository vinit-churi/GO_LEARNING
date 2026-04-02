//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"strings"
)

func NormalizeReport(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i, line := range lines {
		lines[i] = strings.Join(strings.Fields(line), " ")
	}
	return strings.Join(lines, "\n")
}

func BuildGolden(name string, count int) string {
	return fmt.Sprintf("name: %s\ncount: %d", name, count)
}

func ReportsMatch(left, right string) bool {
	return NormalizeReport(left) == NormalizeReport(right)
}

func main() {}
