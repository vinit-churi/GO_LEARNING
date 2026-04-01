//go:build solution
// +build solution

package main

import "fmt"

func IsEven(n int) bool {
	return n%2 == 0
}

func GradeLabel(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "F"
	}
}

func Average(values []int) float64 {
	total := 0
	for _, value := range values {
		total += value
	}
	return float64(total) / float64(len(values))
}

func main() {
	fmt.Println(IsEven(4))
	fmt.Println(GradeLabel(86))
	fmt.Println(Average([]int{4, 6, 8}))
}
