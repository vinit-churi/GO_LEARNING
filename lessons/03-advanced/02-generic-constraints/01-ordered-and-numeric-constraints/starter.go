//go:build !solution
// +build !solution

package main

import (
	"cmp"
	"fmt"
)

type Number interface {
	~int | ~int64 | ~float64
}

func SumNumbers[T Number](items []T) T {
	var total T
	for _, item := range items {
		total += item
	}
	return total
}

func MaxOrdered[T cmp.Ordered](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	max := items[0]
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}
	return max
}

func ClampOrdered[T cmp.Ordered](value, low, high T) T {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}

func main() {
	fmt.Println(SumNumbers([]int{1, 2, 3}), MaxOrdered([]string{"go", "zebra", "api"}))
}
