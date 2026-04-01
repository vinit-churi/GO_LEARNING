//go:build !starter
// +build !starter

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Exercise 1: Basic Named Returns ===")
	exercise1()
	
	fmt.Println("\n=== Exercise 2: Named Returns with Error Handling ===")
	exercise2()
	
	fmt.Println("\n=== Exercise 3: When to Use Named Returns ===")
	exercise3()
	
	fmt.Println("\n=== Exercise 4: Complex Named Returns ===")
	exercise4()
	
	fmt.Println("\n=== Exercise 5: Named Returns Best Practices ===")
	exercise5()
}

// Exercise 1 Demo
func exercise1() {
	tests := []struct {
		area, perimeter int
	}{
		{24, 20},  // 6 x 4
		{12, 14},  // 4 x 3
		{20, 18},  // 5 x 4
		{10, 100}, // no solution
	}
	
	for _, test := range tests {
		length, width := rectangleDimensions(test.area, test.perimeter)
		fmt.Printf("Area=%d, Perimeter=%d → length=%d, width=%d\n", 
			test.area, test.perimeter, length, width)
	}
}

// Exercise 1: Basic Named Returns
func rectangleDimensions(area, perimeter int) (length, width int) {
	// Try all possible widths from 1 to area
	for w := 1; w*w <= area; w++ {
		if area%w == 0 { // w divides area evenly
			l := area / w
			// Check if this gives us the right perimeter
			if 2*(l+w) == perimeter {
				length = l
				width = w
				return // naked return with found values
			}
		}
	}
	// No valid solution found, return zero values (already initialized)
	return
}

// Exercise 2 Demo
func exercise2() {
	tests := []struct {
		a, b float64
	}{
		{10, 2},
		{10, 0},
		{15, 3},
		{7, 0},
	}
	
	for _, test := range tests {
		result, err := safeDivide(test.a, test.b)
		if err != nil {
			fmt.Printf("%.1f / %.1f = error: %v\n", test.a, test.b, err)
		} else {
			fmt.Printf("%.1f / %.1f = %.2f\n", test.a, test.b, result)
		}
	}
}

// Exercise 2: Named Returns with Error Handling
func safeDivide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return // naked return: result=0.0 (zero value), err=error
	}
	result = a / b
	return // naked return: result=calculated, err=nil (zero value)
}

// Exercise 3 Demo
func exercise3() {
	names := []string{"John Doe", "Madonna", "Ada Lovelace", ""}
	
	fmt.Println("Using parseNameNamed (named returns + naked return):")
	for _, name := range names {
		first, last, err := parseNameNamed(name)
		if err != nil {
			fmt.Printf("  %q → error: %v\n", name, err)
		} else {
			fmt.Printf("  %q → first=%q, last=%q\n", name, first, last)
		}
	}
	
	fmt.Println("\nUsing parseNameExplicit (explicit returns):")
	for _, name := range names {
		first, last, err := parseNameExplicit(name)
		if err != nil {
			fmt.Printf("  %q → error: %v\n", name, err)
		} else {
			fmt.Printf("  %q → first=%q, last=%q\n", name, first, last)
		}
	}
}

// Exercise 3A: Named Returns Implementation
func parseNameNamed(fullName string) (firstName, lastName string, err error) {
	index := strings.Index(fullName, " ")
	if index == -1 {
		err = errors.New("invalid name format")
		return // naked return: firstName="", lastName="", err=error
	}
	firstName = fullName[:index]
	lastName = fullName[index+1:]
	return // naked return: firstName=set, lastName=set, err=nil
}

// Exercise 3B: Explicit Returns Implementation
func parseNameExplicit(fullName string) (string, string, error) {
	index := strings.Index(fullName, " ")
	if index == -1 {
		return "", "", errors.New("invalid name format")
	}
	firstName := fullName[:index]
	lastName := fullName[index+1:]
	return firstName, lastName, nil
}

// Exercise 4 Demo
func exercise4() {
	tests := [][]int{
		{2, 4, 6, 8},
		{1, 2, 3, 4, 5},
		{10},
		{},
		{-5, 0, 5},
	}
	
	for _, test := range tests {
		sum, count, avg, err := calculateStats(test)
		if err != nil {
			fmt.Printf("%v → error: %v\n", test, err)
		} else {
			fmt.Printf("%v → sum=%d, count=%d, avg=%.2f\n", test, sum, count, avg)
		}
	}
}

// Exercise 4: Complex Named Returns
func calculateStats(numbers []int) (sum, count int, avg float64, err error) {
	if len(numbers) == 0 {
		return 0, 0, 0.0, errors.New("empty slice") // explicit return for error case
	}
	
	// Calculate sum
	for _, num := range numbers {
		sum += num
	}
	
	// Set count
	count = len(numbers)
	
	// Calculate average
	avg = float64(sum) / float64(count)
	
	return // naked return: all values are set, no error
}

// Exercise 5 Demo
func exercise5() {
	tests := []string{
		"  hello  ",
		"  hi  ",
		"go",
		"   programming   ",
		"  a  ",
		"okay",
	}
	
	for _, test := range tests {
		output, valid := validateAndFormat(test)
		fmt.Printf("%q → output=%q, valid=%t\n", test, output, valid)
	}
}

// Exercise 5: Named Returns Best Practices
// Named returns for documentation, explicit returns for clarity
func validateAndFormat(input string) (output string, valid bool) {
	trimmed := strings.TrimSpace(input)
	
	if len(trimmed) < 3 {
		return "", false // explicit return: clear what we're returning
	}
	
	return strings.ToUpper(trimmed), true // explicit return: clear what we're returning
}
