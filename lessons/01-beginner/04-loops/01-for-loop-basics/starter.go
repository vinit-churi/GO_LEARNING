//go:build starter
// +build starter

package main

import "fmt"

// Exercise 1: CountToN returns a slice containing numbers from 1 to n (inclusive).
// Use a classic for loop with init, condition, and post components.
func CountToN(n int) []int {
	// TODO: Implement this function
	// Hint: Create a slice, use a for loop from 1 to n, append each number
	return nil
}

// Exercise 2: MultiplicationTable generates a multiplication table of the given size.
// Use nested for loops to fill a 2D slice.
func MultiplicationTable(size int) [][]int {
	// TODO: Implement this function
	// Hint: Create a 2D slice, use nested loops
	// Element at [i][j] should be (i+1) * (j+1)
	return nil
}

// Exercise 3: FindFirstDivisible finds the first number divisible by the divisor.
// Use a for loop with break to exit early when found.
func FindFirstDivisible(numbers []int, divisor int) int {
	// TODO: Implement this function
	// Hint: Loop through numbers, check if num % divisor == 0
	// Return the number immediately when found, or -1 if not found
	return -1
}

// Exercise 4: SumOddNumbers sums only the odd numbers in a slice.
// Use a for loop with continue to skip even numbers.
func SumOddNumbers(numbers []int) int {
	// TODO: Implement this function
	// Hint: Loop through numbers, use continue to skip even numbers (num % 2 == 0)
	// Add odd numbers to sum
	return 0
}

// Exercise 5: ReadUntilNegative reads numbers from a slice until a negative number is found.
// Use an infinite loop pattern with a break condition.
func ReadUntilNegative(numbers []int) []int {
	// TODO: Implement this function
	// Hint: Use for loop with break conditions:
	// - Break if you reach the end of the slice
	// - Break if you find a negative number
	return nil
}

func main() {
	fmt.Println("Exercise 1: CountToN")
	fmt.Println(CountToN(5))
	fmt.Println(CountToN(0))
	fmt.Println()

	fmt.Println("Exercise 2: MultiplicationTable")
	table := MultiplicationTable(3)
	for _, row := range table {
		fmt.Println(row)
	}
	fmt.Println()

	fmt.Println("Exercise 3: FindFirstDivisible")
	fmt.Println(FindFirstDivisible([]int{3, 5, 10, 15}, 5))
	fmt.Println(FindFirstDivisible([]int{1, 2, 3}, 7))
	fmt.Println()

	fmt.Println("Exercise 4: SumOddNumbers")
	fmt.Println(SumOddNumbers([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumOddNumbers([]int{2, 4, 6}))
	fmt.Println()

	fmt.Println("Exercise 5: ReadUntilNegative")
	fmt.Println(ReadUntilNegative([]int{1, 2, -1, 3}))
	fmt.Println(ReadUntilNegative([]int{5, 10, 15}))
}
