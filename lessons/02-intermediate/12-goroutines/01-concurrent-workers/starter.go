//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"sync"
)

func DoubleAsync(values []int) []int {
	out := make([]int, len(values))
	var wg sync.WaitGroup

	for i, value := range values {
		wg.Add(1)
		go func(index, v int) {
			defer wg.Done()
			out[index] = v * 2
		}(i, value)
	}

	wg.Wait()
	return out
}

func Sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func DoubleAndSum(values []int) int {
	return Sum(DoubleAsync(values))
}

func main() {
	fmt.Println(DoubleAndSum([]int{1, 2, 3}))
}
