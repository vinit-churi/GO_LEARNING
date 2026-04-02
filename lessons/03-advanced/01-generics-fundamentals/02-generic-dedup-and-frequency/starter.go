//go:build !solution
// +build !solution

package main

import "fmt"

func Unique[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	out := make([]T, 0, len(items))

	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}

	return out
}

func CountBy[T comparable](items []T) map[T]int {
	counts := make(map[T]int, len(items))
	for _, item := range items {
		counts[item]++
	}
	return counts
}

func MostFrequent[T comparable](items []T) (T, bool) {
	var zero T
	if len(items) == 0 {
		return zero, false
	}

	counts := CountBy(items)
	best := items[0]
	bestCount := counts[best]

	for _, item := range items[1:] {
		if counts[item] > bestCount {
			best = item
			bestCount = counts[item]
		}
	}

	return best, true
}

func main() {
	fmt.Println(Unique([]string{"go", "go", "rust", "go", "zig"}))
}
