//go:build !starter
// +build !starter

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
func CheckPositive(n int) string {
	if n > 0 {
		return "positive"
	}
	return "not positive"
}

// Exercise 2: If-Else
func IsEven(n int) string {
	if n%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

// Exercise 3: If-Else If-Else Chain
func CompareNumbers(a, b int) string {
	if a == b {
		return "equal"
	} else if a > b {
		return "greater"
	} else {
		return "less"
	}
}

// Exercise 4: If With Short Statement
func DivideAndCheck(a, b int) string {
	if result := a / b; result > 10 {
		return "large"
	} else if result >= 1 {
		return "medium"
	} else {
		return "small"
	}
}

// Exercise 5: Complex Condition
func CheckRange(n int) string {
	if n >= 10 && n <= 20 {
		return "in range"
	}
	return "out of range"
}
