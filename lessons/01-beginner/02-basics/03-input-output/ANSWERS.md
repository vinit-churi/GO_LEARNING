# Input/Output - Answers

## Exercise 1: Basic Printing

**Solution:** See `BasicPrint` function in `solution.go`

**Explanation:**

The key difference between `fmt.Print` and `fmt.Println`:
- `fmt.Print` outputs text without adding a newline
- `fmt.Println` outputs text and adds a newline at the end

To print "Hello World!" on one line, we use `fmt.Print` for parts that should stay on the same line, and `fmt.Println` for the last part to add the newline:

```go
fmt.Print("Hello")
fmt.Print(" ")
fmt.Print("World")
fmt.Println("!")
```

**Alternative Approach:**

You could also accomplish this with a single `fmt.Println`:
```go
fmt.Println("Hello World!")
```

This is simpler but doesn't demonstrate the difference between `Print` and `Println`.

---

## Exercise 2: Formatted Output

**Solution:** See `FormattedOutput` function in `solution.go`

**Explanation:**

`fmt.Printf` allows us to create formatted strings using format verbs:
- `%s` for strings
- `%d` for decimal integers
- `%f` for floating-point numbers

```go
fmt.Printf("Name: %s, Age: %d, Height: %.2fm\n", name, age, height)
```

**Key Points:**

1. **Format Verbs Match Types:** Each `%` placeholder corresponds to an argument in order
2. **Precision Control:** `.2f` means 2 decimal places for the float
3. **Explicit Newline:** We add `\n` because `Printf` doesn't add one automatically

**Common Pitfall:**

Forgetting the `\n` at the end means the next output will continue on the same line.

---

## Exercise 3: Format Verbs Practice

**Solution:** See `FormatVerbsPractice` function in `solution.go`

**Explanation:**

This exercise demonstrates different format verbs:

```go
fmt.Printf("Value: %v\n", value)
fmt.Printf("Type: %T\n", value)
```

The `%v` verb provides the default representation of any value, while `%T` prints the type.

**Type Assertion for Strings:**

To print strings with quotes, we need to check if the value is actually a string:

```go
if str, ok := value.(string); ok {
    fmt.Printf("String: %q\n", str)
}
```

- `value.(string)` is a type assertion
- It returns two values: the string value and a boolean indicating success
- `%q` formats a string with quotes and proper escaping

**Why %v is Useful:**

The `%v` verb works with any type and uses sensible defaults. It's perfect for debugging because you don't need to know the exact type beforehand.

---

## Exercise 4: Reading User Input

**Solution:** See `ReadUserInput` function in `solution.go`

**Explanation:**

Reading input requires:
1. Variables to store the input (declared beforehand)
2. Pointers to those variables passed to `fmt.Fscan`
3. Error handling (good practice)

```go
var name string
var age int
_, err := fmt.Fscan(reader, &name, &age)
if err != nil {
    return "", 0, err
}
return name, age, nil
```

**Important Points:**

1. **Pointers Required:** `fmt.Fscan` needs pointers (`&name`, `&age`) to modify the variables
2. **`fmt.Fscan` vs `fmt.Scan`:** We use `fmt.Fscan` to read from any `io.Reader`, not just standard input. This makes the code testable.
3. **Error Handling:** Always check the error return value in production code
4. **Whitespace:** `Scan` automatically handles whitespace between values

**Real-World Usage:**

In a real program reading from standard input, you would use:
```go
fmt.Scan(&name, &age)
```

But for testable code, accepting an `io.Reader` parameter allows you to test with strings:
```go
reader := strings.NewReader("Alice 30")
name, age, _ := ReadUserInput(reader)
```

**Alternative: fmt.Scanln vs fmt.Scanf**

- `fmt.Scanln` stops at newline - useful for line-by-line input
- `fmt.Scanf` uses a format string for structured input:
  ```go
  fmt.Scanf("%s %d", &name, &age)
  ```

---

## General Tips

### Debugging Print Output

When debugging, `fmt.Printf("%#v\n", value)` is very helpful - it shows the Go syntax representation:
```go
fmt.Printf("%#v\n", "hello")  // Output: "hello"
fmt.Printf("%#v\n", 42)       // Output: 42
fmt.Printf("%#v\n", []int{1,2,3})  // Output: []int{1, 2, 3}
```

### Format Verb Cheat Sheet

For quick reference:
- `%v` - when you just want to see the value
- `%T` - when you want to know the type
- `%d` - integers
- `%f` - floats (use `%.2f` for 2 decimal places)
- `%s` - strings
- `%q` - quoted strings
- `%t` - booleans

### Reading Input in Real Programs

For robust input handling, consider:
- `bufio.Scanner` for line-by-line reading
- `bufio.Reader` for buffered reading
- Validation and error messages for invalid input
- Trimming whitespace: `strings.TrimSpace()`

Example:
```go
scanner := bufio.NewScanner(os.Stdin)
fmt.Print("Enter your name: ")
scanner.Scan()
name := scanner.Text()
```
