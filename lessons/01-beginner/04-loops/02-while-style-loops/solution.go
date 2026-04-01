//go:build !starter
// +build !starter

package main

import "fmt"

func main() {
	// Test each exercise
	fmt.Println("Exercise 1: Countdown")
	fmt.Println(Countdown(5))
	fmt.Println(Countdown(3))
	fmt.Println()

	fmt.Println("Exercise 2: Sum Until Threshold")
	fmt.Println(SumUntilThreshold([]int{5, 10, 15, 20}, 25))
	fmt.Println(SumUntilThreshold([]int{1, 2, 3, 4}, 100))
	fmt.Println()

	fmt.Println("Exercise 3: Read Until Sentinel")
	fmt.Println(ReadUntilSentinel([]string{"hello", "world", "STOP", "ignored"}, "STOP"))
	fmt.Println(ReadUntilSentinel([]string{"a", "b", "c"}, "STOP"))
	fmt.Println()

	fmt.Println("Exercise 4: Find First Divisor")
	fmt.Println(FindFirstDivisor(20, 3))
	fmt.Println(FindFirstDivisor(17, 2))
	fmt.Println(FindFirstDivisor(100, 7))
	fmt.Println()

	fmt.Println("Exercise 5: Collatz Sequence Length")
	fmt.Println(CollatzLength(10))
	fmt.Println(CollatzLength(27))
	fmt.Println(CollatzLength(1))
}

// Exercise 1: Countdown
// Count down from start to 1 using a while-style loop
func Countdown(start int) []int {
	result := []int{}
	n := start
	
	for n >= 1 {
		result = append(result, n)
		n--
	}
	
	return result
}

// Exercise 2: Sum Until Threshold
// Add numbers from slice until sum reaches or exceeds threshold
func SumUntilThreshold(numbers []int, threshold int) int {
	sum := 0
	i := 0
	
	for sum < threshold && i < len(numbers) {
		sum += numbers[i]
		i++
	}
	
	return sum
}

// Exercise 3: Read Until Sentinel
// Read values from slice until sentinel is found
func ReadUntilSentinel(values []string, sentinel string) []string {
	result := []string{}
	i := 0
	
	for i < len(values) {
		if values[i] == sentinel {
			break
		}
		result = append(result, values[i])
		i++
	}
	
	return result
}

// Exercise 4: Find First Divisor
// Find the first number >= start that divides n evenly
func FindFirstDivisor(n int, start int) int {
	candidate := start
	
	for candidate <= n {
		if n%candidate == 0 {
			return candidate
		}
		candidate++
	}
	
	return n
}

// Exercise 5: Collatz Sequence Length
// Calculate the length of the Collatz sequence starting from n
func CollatzLength(n int) int {
	steps := 0
	
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	
	return steps
}
