package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// Test Exercise 1: Greet
func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Simple name", "Alice", "Hello, Alice!"},
		{"Another name", "Bob", "Hello, Bob!"},
		{"Name with space", "John Doe", "Hello, John Doe!"},
		{"Empty string", "", "Hello, !"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Test Exercise 2: Add
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive numbers", 5, 3, 8},
		{"With zero", 10, 0, 10},
		{"Negative numbers", -5, -3, -8},
		{"Mixed signs", -5, 10, 5},
		{"Large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Test Exercise 3: PrintInfo
func TestPrintInfo(t *testing.T) {
	tests := []struct {
		name     string
		nameArg  string
		age      int
		expected string
	}{
		{"Standard input", "Bob", 25, "Name: Bob, Age: 25\n"},
		{"Different person", "Alice", 30, "Name: Alice, Age: 30\n"},
		{"Young age", "Charlie", 5, "Name: Charlie, Age: 5\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintInfo(tt.nameArg, tt.age)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			if output != tt.expected {
				t.Errorf("PrintInfo(%q, %d) printed %q; want %q", tt.nameArg, tt.age, output, tt.expected)
			}
		})
	}
}

// Test Exercise 4: Multiply
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Simple multiplication", 3, 4, 12},
		{"With one", 7, 1, 7},
		{"With zero", 5, 0, 0},
		{"Negative numbers", -3, 4, -12},
		{"Both negative", -3, -4, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Test Exercise 4: CalculateRectangleArea
func TestCalculateRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		expected int
	}{
		{"Standard rectangle", 5, 4, 20},
		{"Square", 6, 6, 36},
		{"Wide rectangle", 10, 2, 20},
		{"Tall rectangle", 3, 8, 24},
		{"Unit square", 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateRectangleArea(tt.width, tt.height)
			if result != tt.expected {
				t.Errorf("CalculateRectangleArea(%d, %d) = %d; want %d", tt.width, tt.height, result, tt.expected)
			}
		})
	}
}

// Test Exercise 5: CelsiusToFahrenheit
func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
		epsilon  float64
	}{
		{"Freezing point", 0.0, 32.0, 0.01},
		{"Boiling point", 100.0, 212.0, 0.01},
		{"Body temperature", 37.0, 98.6, 0.01},
		{"Negative temperature", -40.0, -40.0, 0.01},
		{"Room temperature", 25.0, 77.0, 0.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CelsiusToFahrenheit(tt.celsius)
			if abs(result-tt.expected) > tt.epsilon {
				t.Errorf("CelsiusToFahrenheit(%f) = %f; want %f", tt.celsius, result, tt.expected)
			}
		})
	}
}

// Helper function to calculate absolute value of float64
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Benchmark Exercise 2: Add
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(5, 3)
	}
}

// Benchmark Exercise 4: CalculateRectangleArea
func BenchmarkCalculateRectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateRectangleArea(5, 4)
	}
}

// Example for Exercise 1
func ExampleGreet() {
	fmt.Println(Greet("World"))
	// Output: Hello, World!
}

// Example for Exercise 2
func ExampleAdd() {
	fmt.Println(Add(10, 20))
	// Output: 30
}

// Example for Exercise 5
func ExampleCelsiusToFahrenheit() {
	fmt.Printf("%.1f\n", CelsiusToFahrenheit(0))
	// Output: 32.0
}
