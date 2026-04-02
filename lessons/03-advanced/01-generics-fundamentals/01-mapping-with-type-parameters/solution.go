//go:build solution
// +build solution

package main

import "fmt"

func MapSlice[T, U any](items []T, fn func(T) U) []U {
	out := make([]U, len(items))
	for i, item := range items {
		out[i] = fn(item)
	}
	return out
}

func FirstOrZero[T any](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[0]
}

func PairLabels[T any](items []T, label func(T) string) []string {
	return MapSlice(items, label)
}

func main() {
	values := MapSlice([]int{1, 2, 3}, func(v int) string { return fmt.Sprintf("n=%d", v) })
	fmt.Println(values, FirstOrZero([]string{"go", "is", "fast"}))
}
