//go:build solution
// +build solution

package main

import "fmt"

func Feed(values []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			out <- value
		}
	}()
	return out
}

func SquareStream(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			out <- value * value
		}
	}()
	return out
}

func Collect(in <-chan int) []int {
	var values []int
	for value := range in {
		values = append(values, value)
	}
	return values
}

func main() {
	fmt.Println(Collect(SquareStream(Feed([]int{1, 2, 3}))))
}
