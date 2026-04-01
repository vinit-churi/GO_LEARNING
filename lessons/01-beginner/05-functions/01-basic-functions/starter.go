//go:build !solution
// +build !solution

package main

import "fmt"

// Exercise 1: Simple Greeting Function
// Create a function that takes a name and returns a greeting message.
// TODO: Implement the Greet function
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// Exercise 2: Add Two Numbers
// Create a function that takes two integers and returns their sum.
// TODO: Implement the Add function
func Add(a, b int) int {
	return a + b
}

// Exercise 3: Print Information (No Return Value)
// Create a function that prints name and age information.
// TODO: Implement the PrintInfo function
func PrintInfo(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

// Exercise 4: Calculate Rectangle Area Using Helper Functions
// Create a Multiply function and use it in CalculateRectangleArea.
// TODO: Implement the Multiply function
func Multiply(a, b int) int {
	return a * b
}

// TODO: Implement the CalculateRectangleArea function
func CalculateRectangleArea(width, height int) int {
	return Multiply(width, height)
}

// Exercise 5: Temperature Converter
// Convert Celsius to Fahrenheit using the formula: F = C * 9/5 + 32
// TODO: Implement the CelsiusToFahrenheit function
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9.0/5.0 + 32.0
}

func main() {
	// Test your functions here
	fmt.Println("=== Exercise 1: Greet ===")
	fmt.Println(Greet("Alice"))

	fmt.Println("\n=== Exercise 2: Add ===")
	fmt.Println("5 + 3 =", Add(5, 3))

	fmt.Println("\n=== Exercise 3: PrintInfo ===")
	PrintInfo("Bob", 25)

	fmt.Println("\n=== Exercise 4: CalculateRectangleArea ===")
	fmt.Println("Area of 5x4 rectangle:", CalculateRectangleArea(5, 4))

	fmt.Println("\n=== Exercise 5: CelsiusToFahrenheit ===")
	fmt.Println("0°C =", CelsiusToFahrenheit(0.0), "°F")
	fmt.Println("100°C =", CelsiusToFahrenheit(100.0), "°F")
}
