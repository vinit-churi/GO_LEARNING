# Input/Output in Go

**Difficulty:** Beginner  
**Prerequisites:** Variables and Types  
**Estimated Time:** 30 minutes

## Learning Objectives

By the end of this lesson, you will be able to:

- Use `fmt.Print` and `fmt.Println` to output text
- Format output using `fmt.Printf` with format verbs
- Understand common format verbs (`%s`, `%d`, `%f`, `%v`, `%T`)
- Read user input with `fmt.Scan`

## Concept Summary

Go's `fmt` package provides functions for formatted I/O operations. The most common functions are:

- **`fmt.Print`**: prints values without a newline
- **`fmt.Println`**: prints values with a newline at the end
- **`fmt.Printf`**: prints formatted output using format verbs
- **`fmt.Scan`**: reads input from standard input

Format verbs are special placeholders used with `fmt.Printf`:
- `%s` - string
- `%d` - decimal integer
- `%f` - floating-point number
- `%v` - default format for any value
- `%T` - type of the value

## Run Instructions

To run the starter code:
```bash
go run -tags=starter starter.go
```

To run the solution:
```bash
go run -tags=solution solution.go
```

## Test Instructions

Run all tests:
```bash
go test -v
```

Run tests with the starter code:
```bash
go test -v -tags=starter
```

Run tests with the solution code:
```bash
go test -v -tags=solution
```

## Exercises

### Exercise 1: Basic Printing

**Objective:** Practice using `fmt.Print` and `fmt.Println`

Implement the `BasicPrint` function that:
- Prints "Hello" using `fmt.Print`
- Prints a space
- Prints "World" using `fmt.Print`
- Prints "!" using `fmt.Println`

Expected output:
```
Hello World!
```

### Exercise 2: Formatted Output

**Objective:** Use `fmt.Printf` with format verbs

Implement the `FormattedOutput` function that takes:
- `name` (string)
- `age` (int)
- `height` (float64)

Print a formatted message: "Name: [name], Age: [age], Height: [height]m"

Example:
- Input: "Alice", 30, 1.65
- Output: "Name: Alice, Age: 30, Height: 1.65m"

### Exercise 3: Format Verbs Practice

**Objective:** Practice using different format verbs

Implement the `FormatVerbsPractice` function that takes a value of any type and:
- Prints the value using `%v` (default format)
- Prints the type using `%T`
- If the value is a string, also prints it using `%s` with quotes

Example for string "hello":
```
Value: hello
Type: string
String: "hello"
```

Example for integer 42:
```
Value: 42
Type: int
```

### Exercise 4: Reading User Input

**Objective:** Read input using `fmt.Scan`

Implement the `ReadUserInput` function that:
- Reads a name (string) and age (int) from the provided input
- Returns both values

Note: The function signature uses `io.Reader` to allow testing with different input sources.
