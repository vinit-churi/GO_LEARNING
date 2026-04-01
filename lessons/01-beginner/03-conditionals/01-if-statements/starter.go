//go:build starter
// +build starter

package main

import "fmt"

func main() {
	// Test Exercise 1
	fmt.Println("Exercise 1: Basic If Statement")
	fmt.Println(CheckPositive(5))
	fmt.Println(CheckPositive(-3))
	fmt.Println(CheckPositive(0))
	fmt.Println()

	// Test Exercise 2
	fmt.Println("Exercise 2: If-Else")
	fmt.Println(IsEven(4))
	fmt.Println(IsEven(7))
	fmt.Println(IsEven(0))
	fmt.Println()

	// Test Exercise 3
	fmt.Println("Exercise 3: If-Else If-Else Chain")
	fmt.Println(CompareNumbers(5, 3))
	fmt.Println(CompareNumbers(2, 7))
	fmt.Println(CompareNumbers(4, 4))
	fmt.Println()

	// Test Exercise 4
	fmt.Println("Exercise 4: If With Short Statement")
	fmt.Println(DivideAndCheck(100, 5))
	fmt.Println(DivideAndCheck(20, 4))
	fmt.Println(DivideAndCheck(5, 10))
	fmt.Println()

	// Test Exercise 5
	fmt.Println("Exercise 5: Complex Condition")
	fmt.Println(CheckRange(15))
	fmt.Println(CheckRange(5))
	fmt.Println(CheckRange(10))
	fmt.Println(CheckRange(20))
	fmt.Println(CheckRange(25))
}

// Exercise 1: Basic If Statement
// Write a function that returns "positive" if n > 0, otherwise "not positive"
func CheckPositive(n int) string {
	// TODO: Implement this function
	return ""
}

// Exercise 2: If-Else
// Write a function that returns "even" if n is even, otherwise "odd"
// Hint: Use the modulo operator %
func IsEven(n int) string {
	// TODO: Implement this function
	return ""
}

// Exercise 3: If-Else If-Else Chain
// Write a function that compares two numbers
// Returns "equal" if a == b, "greater" if a > b, "less" if a < b
func CompareNumbers(a, b int) string {
	// TODO: Implement this function
	return ""
}

// Exercise 4: If With Short Statement
// Write a function that divides a by b and checks the result
// Use an if statement with short variable declaration
// Returns "large" if result > 10, "medium" if 1 <= result <= 10, "small" if result < 1
func DivideAndCheck(a, b int) string {
	// TODO: Implement this function using if with short statement
	// Hint: if result := a / b; result > 10 { ... }
	return ""
}

// Exercise 5: Complex Condition
// Write a function that returns "in range" if n is between 10 and 20 (inclusive)
// Otherwise returns "out of range"
// Hint: Use the && operator to combine conditions
func CheckRange(n int) string {
	// TODO: Implement this function
	return ""
}
