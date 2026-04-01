//go:build starter
// +build starter

package main

import "fmt"

// Exercise 1: Sum Function
// Implement a variadic function that calculates the sum of any number of integers.
// If no arguments are provided, return 0.
func Sum(numbers ...int) int {
	// TODO: Implement sum logic
	return 0
}

// Exercise 2: Max Function
// Implement a variadic function that returns the maximum value from any number of integers.
// If no arguments are provided, return 0.
func Max(numbers ...int) int {
	// TODO: Handle empty input
	// TODO: Initialize max with first element
	// TODO: Compare remaining elements
	return 0
}

// Exercise 3: Join With Separator
// Implement a variadic function that takes a separator string as the first parameter
// and any number of strings to join.
func JoinWithSeparator(separator string, words ...string) string {
	// TODO: Handle empty words
	// TODO: Build result string with separator between words
	return ""
}

// Exercise 4: Filter Evens
// Implement a variadic function that accepts any number of integers
// and returns a slice containing only the even numbers.
func FilterEvens(numbers ...int) []int {
	// TODO: Create result slice
	// TODO: Check each number if it's even (hint: use % 2)
	// TODO: Append even numbers to result
	return []int{}
}

// Exercise 5: Concatenate All (Bonus)
// Implement a variadic function that accepts multiple slices of integers
// and returns a single slice containing all elements.
func ConcatenateAll(slices ...[]int) []int {
	// TODO: Calculate total length (optional optimization)
	// TODO: Create result slice
	// TODO: Iterate through each slice and append its elements
	return []int{}
}

func main() {
	fmt.Println("=== Exercise 1: Sum ===")
	fmt.Println("Sum(1, 2, 3):", Sum(1, 2, 3))
	fmt.Println("Sum(10, 20):", Sum(10, 20))
	fmt.Println("Sum():", Sum())
	
	fmt.Println("\n=== Exercise 2: Max ===")
	fmt.Println("Max(1, 5, 3, 2):", Max(1, 5, 3, 2))
	fmt.Println("Max(-10, -5, -20):", Max(-10, -5, -20))
	fmt.Println("Max():", Max())
	
	fmt.Println("\n=== Exercise 3: Join With Separator ===")
	fmt.Println("Join with ', ':", JoinWithSeparator(", ", "apple", "banana", "cherry"))
	fmt.Println("Join with ' - ':", JoinWithSeparator(" - ", "Go", "is", "awesome"))
	fmt.Println("Join with empty words:", JoinWithSeparator(", "))
	
	fmt.Println("\n=== Exercise 4: Filter Evens ===")
	fmt.Println("Filter [1,2,3,4,5,6]:", FilterEvens(1, 2, 3, 4, 5, 6))
	fmt.Println("Filter [1,3,5,7]:", FilterEvens(1, 3, 5, 7))
	fmt.Println("Filter []:", FilterEvens())
	
	fmt.Println("\n=== Exercise 5: Concatenate All ===")
	result := ConcatenateAll([]int{1, 2}, []int{3, 4}, []int{5})
	fmt.Println("Concatenate [[1,2], [3,4], [5]]:", result)
	
	// Bonus: Using spread operator with a slice
	nums := []int{5, 10, 15, 20}
	fmt.Println("\n=== Bonus: Spread Operator ===")
	fmt.Println("Sum(nums...):", Sum(nums...))
	fmt.Println("Max(nums...):", Max(nums...))
}
