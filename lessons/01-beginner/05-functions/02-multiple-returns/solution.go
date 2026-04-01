//go:build !starter
// +build !starter

package main

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
)

func main() {
	exercise1()
	fmt.Println()
	
	exercise2()
	fmt.Println()
	
	exercise3()
	fmt.Println()
	
	exercise4()
	fmt.Println()
	
	exercise5()
}

// Exercise 1: Return two values
func exercise1() {
	fmt.Println("Exercise 1: Min and Max")
	min, max := minMax([]int{3, 7, 1, 9, 2})
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}

// Exercise 2: Return value and error
func exercise2() {
	fmt.Println("Exercise 2: Safe Division")
	
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}
	
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %d\n", result)
	}
}

// Exercise 3: Parse user input
func exercise3() {
	fmt.Println("Exercise 3: Parse User Input")
	
	name, age, err := parseUserInput("Alice:30")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
	
	name, age, err = parseUserInput("Bob")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}

// Exercise 4: Swap values
func exercise4() {
	fmt.Println("Exercise 4: Swap Values")
	a, b := "hello", "world"
	fmt.Printf("Before swap: a=%s, b=%s\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After swap: a=%s, b=%s\n", a, b)
}

// Exercise 5: Calculate statistics
func exercise5() {
	fmt.Println("Exercise 5: Calculate Statistics")
	
	sum, avg, count, err := calculateStats([]int{10, 20, 30})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Sum: %d, Average: %.2f, Count: %d\n", sum, avg, count)
	}
	
	sum, avg, count, err = calculateStats([]int{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Sum: %d, Average: %.2f, Count: %d\n", sum, avg, count)
	}
}

// Exercise 1: Return minimum and maximum values from a slice
func minMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	
	min := numbers[0]
	max := numbers[0]
	
	for _, num := range numbers[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	
	return min, max
}

// Exercise 2: Divide two integers safely, returning result and error
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Exercise 3: Parse a "name:age" string into separate values
func parseUserInput(input string) (string, int, error) {
	parts := strings.Split(input, ":")
	
	if len(parts) != 2 {
		return "", 0, errors.New("invalid format: expected 'name:age'")
	}
	
	name := parts[0]
	age, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid age: %v", err)
	}
	
	return name, age, nil
}

// Exercise 4: Swap two string values
func swap(a, b string) (string, string) {
	return b, a
}

// Exercise 5: Calculate sum, average, and count from a slice
func calculateStats(numbers []int) (int, float64, int, error) {
	if len(numbers) == 0 {
		return 0, 0.0, 0, errors.New("cannot calculate stats for empty slice")
	}
	
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	
	count := len(numbers)
	avg := float64(sum) / float64(count)
	
	return sum, avg, count, nil
}
