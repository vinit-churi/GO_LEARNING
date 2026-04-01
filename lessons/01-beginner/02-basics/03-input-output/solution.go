//go:build !starter

package inputoutput

import (
	"fmt"
	"io"
)

// BasicPrint demonstrates the difference between fmt.Print and fmt.Println
// It should print "Hello World!" on a single line
func BasicPrint() {
	fmt.Print("Hello")
	fmt.Print(" ")
	fmt.Print("World")
	fmt.Println("!")
}

// FormattedOutput uses fmt.Printf to print formatted output
// It should print: "Name: [name], Age: [age], Height: [height]m"
func FormattedOutput(name string, age int, height float64) {
	fmt.Printf("Name: %s, Age: %d, Height: %.2fm\n", name, age, height)
}

// FormatVerbsPractice demonstrates different format verbs
// It prints the value using %v, the type using %T,
// and if it's a string, also prints it with quotes using %q
func FormatVerbsPractice(value interface{}) {
	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Type: %T\n", value)

	if str, ok := value.(string); ok {
		fmt.Printf("String: %q\n", str)
	}
}

// ReadUserInput reads a name (string) and age (int) from the provided reader
// It returns the name, age, and any error that occurred
func ReadUserInput(reader io.Reader) (string, int, error) {
	var name string
	var age int

	_, err := fmt.Fscan(reader, &name, &age)
	if err != nil {
		return "", 0, err
	}

	return name, age, nil
}
