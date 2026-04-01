//go:build starter
// +build starter

package main

import "fmt"

func main() {
	// Test each exercise
	fmt.Println("Exercise 1: Countdown")
	fmt.Println(Countdown(5))
	fmt.Println()

	fmt.Println("Exercise 2: Sum Until Threshold")
	fmt.Println(SumUntilThreshold([]int{5, 10, 15, 20}, 25))
	fmt.Println()

	fmt.Println("Exercise 3: Read Until Sentinel")
	fmt.Println(ReadUntilSentinel([]string{"hello", "world", "STOP", "ignored"}, "STOP"))
	fmt.Println()

	fmt.Println("Exercise 4: Find First Divisor")
	fmt.Println(FindFirstDivisor(20, 3))
	fmt.Println(FindFirstDivisor(17, 2))
	fmt.Println()

	fmt.Println("Exercise 5: Collatz Sequence Length")
	fmt.Println(CollatzLength(10))
	fmt.Println(CollatzLength(27))
}

// Exercise 1: Countdown
// Count down from start to 1 using a while-style loop
// Returns a slice with all numbers in the countdown
func Countdown(start int) []int {
	// TODO: Implement countdown using for with only a condition
	// Hint: Use a slice to collect the numbers
	// Hint: Loop while n >= 1
	return nil
}

// Exercise 2: Sum Until Threshold
// Add numbers from slice until sum reaches or exceeds threshold
// Returns the sum when threshold is reached (may exceed it)
func SumUntilThreshold(numbers []int, threshold int) int {
	// TODO: Implement using while-style loop
	// Hint: Keep track of sum and current index
	// Hint: Stop when sum >= threshold OR when out of numbers
	return 0
}

// Exercise 3: Read Until Sentinel
// Read values from slice until sentinel is found
// Returns all values before the sentinel (sentinel is NOT included)
func ReadUntilSentinel(values []string, sentinel string) []string {
	// TODO: Implement using while-style loop
	// Hint: Check each value, stop when you find the sentinel
	// Hint: Don't include the sentinel in the result
	return nil
}

// Exercise 4: Find First Divisor
// Find the first number >= start that divides n evenly
// Returns the first divisor found
func FindFirstDivisor(n int, start int) int {
	// TODO: Implement using while-style loop
	// Hint: Use the modulo operator (%) to check for even division
	// Hint: Loop through candidates starting from 'start'
	// Hint: Return immediately when you find one
	return 0
}

// Exercise 5: Collatz Sequence Length
// Calculate the length of the Collatz sequence starting from n
// Rules: if even divide by 2, if odd multiply by 3 and add 1, stop at 1
// Returns the number of steps to reach 1
func CollatzLength(n int) int {
	// TODO: Implement using while-style loop
	// Hint: Loop while n != 1
	// Hint: Count each step
	// Hint: Use if-else to apply the rules
	return 0
}
