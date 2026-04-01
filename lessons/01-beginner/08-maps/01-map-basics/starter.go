//go:build !solution
// +build !solution

package main

import "fmt"

func BuildScoreboard(names []string) map[string]int {
	scores := make(map[string]int)
	for _, name := range names {
		scores[name] = 0
	}
	return scores
}

func LookupScore(scores map[string]int, name string) (int, bool) {
	score, ok := scores[name]
	return score, ok
}

func CountWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	fmt.Println(BuildScoreboard([]string{"Ava", "Mia"}))
	fmt.Println(LookupScore(map[string]int{"Ava": 12}, "Ava"))
	fmt.Println(CountWords([]string{"go", "go", "test"}))
}
