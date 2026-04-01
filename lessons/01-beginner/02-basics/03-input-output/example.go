//go:build ignore

// Example demonstrates the input/output exercises
package main

import (
	"fmt"
	"strings"

	io "github.com/vinit-churi/learning-go/lessons/01-beginner/02-basics/03-input-output"
)

func main() {
	fmt.Println("=== Exercise 1: Basic Print ===")
	io.BasicPrint()

	fmt.Println("\n=== Exercise 2: Formatted Output ===")
	io.FormattedOutput("Alice", 30, 1.65)
	io.FormattedOutput("Bob", 25, 1.80)

	fmt.Println("\n=== Exercise 3: Format Verbs Practice ===")
	fmt.Println("String example:")
	io.FormatVerbsPractice("hello")
	fmt.Println("\nInteger example:")
	io.FormatVerbsPractice(42)
	fmt.Println("\nFloat example:")
	io.FormatVerbsPractice(3.14)

	fmt.Println("\n=== Exercise 4: Reading User Input ===")
	// Simulating user input
	reader := strings.NewReader("Charlie 35")
	name, age, err := io.ReadUserInput(reader)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read from input - Name: %s, Age: %d\n", name, age)
	}
}
