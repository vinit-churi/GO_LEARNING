//go:build solution
// +build solution

package main

import "fmt"

// Exercise 1: Simple Greeting Function
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// Exercise 2: Add Two Numbers
func Add(a, b int) int {
	return a + b
}

// Exercise 3: Print Information (No Return Value)
func PrintInfo(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

// Exercise 4: Calculate Rectangle Area Using Helper Functions
func Multiply(a, b int) int {
	return a * b
}

func CalculateRectangleArea(width, height int) int {
	return Multiply(width, height)
}

// Exercise 5: Temperature Converter
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
