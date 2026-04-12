//go:build starter

package inputoutput

import (
	"io"
)

// BasicPrint demonstrates the difference between fmt.Print and fmt.Println
// It should print "Hello World!" on a single line
func BasicPrint() {
	// TODO: Use fmt.Print to output "Hello"
	fmt.Print("Hello")
	// TODO: Use fmt.Print to output " "
	fmt.Print(" ")
	// TODO: Use fmt.Print to output "World"
	fmt.Print("World")
	// TODO: Use fmt.Println to output "!"
	fmt.Print("!")
}

// FormattedOutput uses fmt.Printf to print formatted output
// It should print: "Name: [name], Age: [age], Height: [height]m"
func FormattedOutput(name string, age int, height float64) {
	// TODO: Use fmt.Printf with format verbs to print the formatted message
	// Hint: Use %s for string, %d for int, %.2f for float with 2 decimal places
	// Don't forget to add \n at the end
	fmt.Printf("Name: %s, Age: %d, Height: %.2fm\n", name, age, height)
}

// FormatVerbsPractice demonstrates different format verbs
// It prints the value using %v, the type using %T,
// and if it's a string, also prints it with quotes using %q
func FormatVerbsPractice(value interface{}) {
	// TODO: Print "Value: " followed by the value using %v
	// TODO: Print "Type: " followed by the type using %T
	// TODO: If the value is a string, print "String: " followed by the quoted string using %q
	// Hint: Use type assertion to check if value is a string: str, ok := value.(string)
	fmt.Printf("Value: %v\nType: %T\n,String: %q", value, value, value)
}

// ReadUserInput reads a name (string) and age (int) from the provided reader
// It returns the name, age, and any error that occurred
func ReadUserInput(reader io.Reader) (string, int, error) {
	// TODO: Declare variables for name and age
	// TODO: Use fmt.Fscan to read name and age from the reader
	// Hint: Pass pointers to name and age using &
	// TODO: Return the name, age, and error
	_ = reader
	var Name string
	var Age int

	n, err := fmt.Scanf("%s %d", &Name, &Age)

	if err != nil {	
		return "", 0, nil // placeholder return
	}

	return Name, Age, nil
}
