//go:build solution
// +build solution

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func CleanTags(tags []string) []string {
	cleaned := make([]string, 0, len(tags))
	for _, tag := range tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		if tag != "" {
			cleaned = append(cleaned, tag)
		}
	}
	return cleaned
}

func SortedScores(raw []string) ([]int, error) {
	scores := make([]int, 0, len(raw))
	for _, item := range raw {
		value, err := strconv.Atoi(strings.TrimSpace(item))
		if err != nil {
			return nil, err
		}
		scores = append(scores, value)
	}
	slices.Sort(scores)
	return scores, nil
}

func JoinTags(tags []string) string {
	return strings.Join(CleanTags(tags), ",")
}

func main() {
	scores, _ := SortedScores([]string{"9", " 2", "5"})
	fmt.Println(JoinTags([]string{" Go ", "testing", ""}))
	fmt.Println(scores)
}
