//go:build solution

package main

import "fmt"

func totalMessages(days int) int {
	total := 0
	for day := 1; day <= days; day++ {
		total += day
	}

	return total
}

func buildCountdown(start int) []int {
	countdown := []int{}
	for value := start; value >= 1; value-- {
		countdown = append(countdown, value)
	}

	return countdown
}

func sumEvenNumbers(limit int) int {
	total := 0
	for value := 0; value <= limit; value++ {
		if value%2 == 0 {
			total += value
		}
	}

	return total
}

func main() {
	fmt.Println(totalMessages(4))
	fmt.Println(buildCountdown(5))
	fmt.Println(sumEvenNumbers(10))
}
