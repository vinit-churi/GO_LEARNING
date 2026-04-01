//go:build starter
// +build starter

package main

import (
	"fmt"
)

func main() {
	// Exercise 1: Return two values
	fmt.Println("Exercise 1: Min and Max")
	// TODO: Call minMax with []int{3, 7, 1, 9, 2}
	// TODO: Print the min and max values
	
	fmt.Println()
	
	// Exercise 2: Return value and error
	fmt.Println("Exercise 2: Safe Division")
	// TODO: Call safeDivide(10, 2) and handle the result
	// TODO: Call safeDivide(10, 0) and handle the error
	
	fmt.Println()
	
	// Exercise 3: Parse user input
	fmt.Println("Exercise 3: Parse User Input")
	// TODO: Call parseUserInput("Alice:30") and print results
	// TODO: Call parseUserInput("Bob") and handle error
	
	fmt.Println()
	
	// Exercise 4: Swap values
	fmt.Println("Exercise 4: Swap Values")
	// TODO: Create two variables with values "hello" and "world"
	// TODO: Swap them using the swap function
	// TODO: Print the swapped values
	
	fmt.Println()
	
	// Exercise 5: Calculate statistics
	fmt.Println("Exercise 5: Calculate Statistics")
	// TODO: Call calculateStats with []int{10, 20, 30}
	// TODO: Print sum, average, and count
	// TODO: Call calculateStats with empty slice and handle error
}

// Exercise 1: Return minimum and maximum values from a slice
func minMax(numbers []int) (int, int) {
	// TODO: Initialize min and max to the first element
	// TODO: Loop through the slice and update min and max
	// TODO: Return min and max
	return 0, 0
}

// Exercise 2: Divide two integers safely, returning result and error
func safeDivide(a, b int) (int, error) {
	// TODO: Check if b is zero
	// TODO: If zero, return 0 and an error
	// TODO: Otherwise, return a/b and nil
	return 0, nil
}

// Exercise 3: Parse a "name:age" string into separate values
func parseUserInput(input string) (string, int, error) {
	// TODO: Split the input string by ":"
	// TODO: Validate that we have exactly 2 parts
	// TODO: Convert age string to integer
	// TODO: Return name, age, and error appropriately
	return "", 0, nil
}

// Exercise 4: Swap two string values
func swap(a, b string) (string, string) {
	// TODO: Return b and a (swapped)
	return "", ""
}

// Exercise 5: Calculate sum, average, and count from a slice
func calculateStats(numbers []int) (int, float64, int, error) {
	// TODO: Check if slice is empty and return error if so
	// TODO: Calculate sum by iterating through slice
	// TODO: Calculate average as float64
	// TODO: Return sum, average, count, and nil error
	return 0, 0.0, 0, nil
}
