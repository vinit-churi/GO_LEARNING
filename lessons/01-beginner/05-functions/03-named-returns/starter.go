//go:build starter
// +build starter

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Exercise 1: Basic Named Returns ===")
	length, width := rectangleDimensions(24, 20)
	fmt.Printf("Rectangle with area=24, perimeter=20: length=%d, width=%d\n", length, width)
	
	fmt.Println("\n=== Exercise 2: Named Returns with Error Handling ===")
	result, err := safeDivide(10, 2)
	fmt.Printf("10 / 2 = %.2f, error: %v\n", result, err)
	
	result, err = safeDivide(10, 0)
	fmt.Printf("10 / 0 = %.2f, error: %v\n", result, err)
	
	fmt.Println("\n=== Exercise 3: When to Use Named Returns ===")
	first, last, err := parseNameNamed("John Doe")
	fmt.Printf("parseNameNamed('John Doe'): first=%q, last=%q, error=%v\n", first, last, err)
	
	first, last, err = parseNameExplicit("Madonna")
	fmt.Printf("parseNameExplicit('Madonna'): first=%q, last=%q, error=%v\n", first, last, err)
	
	fmt.Println("\n=== Exercise 4: Complex Named Returns ===")
	sum, count, avg, err := calculateStats([]int{2, 4, 6, 8})
	fmt.Printf("Stats for [2,4,6,8]: sum=%d, count=%d, avg=%.2f, error=%v\n", sum, count, avg, err)
	
	sum, count, avg, err = calculateStats([]int{})
	fmt.Printf("Stats for []: sum=%d, count=%d, avg=%.2f, error=%v\n", sum, count, avg, err)
	
	fmt.Println("\n=== Exercise 5: Named Returns Best Practices ===")
	output, valid := validateAndFormat("  hello  ")
	fmt.Printf("validateAndFormat('  hello  '): output=%q, valid=%t\n", output, valid)
	
	output, valid = validateAndFormat("  hi  ")
	fmt.Printf("validateAndFormat('  hi  '): output=%q, valid=%t\n", output, valid)
}

// Exercise 1: Basic Named Returns
// TODO: Implement rectangleDimensions with named returns (length, width int)
// Calculate length and width of a rectangle given its area and perimeter
// Assume integer dimensions and length >= width
// Return 0, 0 if no valid solution exists
func rectangleDimensions(area, perimeter int) (length, width int) {
	// TODO: Implement this function
	// Hint: Try different values where length >= width
	// Check if length * width == area AND 2*(length+width) == perimeter
	// Use a loop or mathematical approach
	// Set length and width variables
	// Use naked return
	return
}

// Exercise 2: Named Returns with Error Handling
// TODO: Implement safeDivide with named returns (result float64, err error)
// Divide a by b, return error if b is zero
// Use naked returns (bare return statements)
func safeDivide(a, b float64) (result float64, err error) {
	// TODO: Check if b is zero
	// TODO: If zero, set err and return
	// TODO: Otherwise, calculate result and return
	return
}

// Exercise 3: When to Use Named Returns
// TODO: Implement parseNameNamed with named returns
// Split fullName on first space into firstName and lastName
// Return error if no space found (message: "invalid name format")
// Use named returns with naked return
func parseNameNamed(fullName string) (firstName, lastName string, err error) {
	// TODO: Find the first space
	// TODO: If no space, set err and return
	// TODO: Split on space, set firstName and lastName
	// TODO: Use naked return
	return
}

// TODO: Implement parseNameExplicit WITHOUT named returns
// Same logic as parseNameNamed but use explicit returns
// Compare the two implementations - which is clearer?
func parseNameExplicit(fullName string) (string, string, error) {
	// TODO: Same logic as above
	// TODO: Use explicit return statements
	// TODO: Return values at each return point
	return "", "", nil
}

// Exercise 4: Complex Named Returns
// TODO: Implement calculateStats with named returns
// Calculate sum, count, and average of a slice of integers
// Return error if slice is empty (message: "empty slice")
// Use at least one explicit return and one naked return
func calculateStats(numbers []int) (sum, count int, avg float64, err error) {
	// TODO: Check if slice is empty, use explicit return for error
	// TODO: Calculate sum using a loop
	// TODO: Set count from len(numbers)
	// TODO: Calculate average: float64(sum) / float64(count)
	// TODO: Use naked return for success case
	return
}

// Exercise 5: Named Returns Best Practices
// TODO: Implement validateAndFormat
// Use named returns for documentation: (output string, valid bool)
// But use EXPLICIT returns (not naked) for clarity
// This demonstrates best practice: named for documentation, explicit for clarity
func validateAndFormat(input string) (output string, valid bool) {
	// TODO: Trim whitespace from input using strings.TrimSpace
	// TODO: Check if trimmed length is at least 3
	// TODO: If too short, return "", false explicitly
	// TODO: If valid, return uppercase version and true explicitly
	// TODO: Use strings.ToUpper for conversion
	return
}
