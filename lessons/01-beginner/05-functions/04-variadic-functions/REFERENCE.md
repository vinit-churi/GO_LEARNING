# Variadic Functions - Reference

## What Are Variadic Functions?

A **variadic function** is a function that can accept a variable number of arguments of the same type. In Go, you declare a variadic parameter by placing three dots (`...`) before the parameter's type.

```go
func myFunc(items ...string) {
    // items is a []string inside the function
}
```

## Syntax Patterns

### Basic Variadic Function

```go
func printAll(messages ...string) {
    for i, msg := range messages {
        fmt.Printf("%d: %s\n", i, msg)
    }
}

// Usage
printAll("Hello", "World")
printAll("Go", "is", "great")
printAll() // Valid - empty slice
```

### Mixing Regular and Variadic Parameters

The variadic parameter **must be the last parameter**:

```go
func greetAll(greeting string, names ...string) string {
    var result string
    for _, name := range names {
        result += greeting + " " + name + "! "
    }
    return result
}

// Usage
greetAll("Hello", "Alice", "Bob")        // "Hello Alice! Hello Bob! "
greetAll("Hi")                           // "" (empty names)
greetAll("Hey", "Charlie")               // "Hey Charlie! "
```

❌ **Invalid:**
```go
func invalid(items ...string, separator string) {} // Won't compile
```

### Passing Slices to Variadic Functions

Use the `...` operator to expand a slice into individual arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Direct arguments
result1 := sum(1, 2, 3, 4)    // 10

// From a slice
nums := []int{1, 2, 3, 4}
result2 := sum(nums...)       // 10
```

### Combining Approaches

```go
func combine(prefix string, words ...string) string {
    all := append([]string{prefix}, words...)
    return strings.Join(all, " ")
}

// Mix direct args and slice expansion
parts := []string{"is", "awesome"}
result := combine("Go", parts...)  // "Go is awesome"
```

### Variadic Functions with Slices

You can have a variadic parameter where each argument is itself a slice:

```go
func flatten(slices ...[]int) []int {
    var result []int
    for _, slice := range slices {
        result = append(result, slice...)
    }
    return result
}

// Usage
a := []int{1, 2, 3}
b := []int{4, 5}
c := []int{6}
all := flatten(a, b, c)  // []int{1, 2, 3, 4, 5, 6}
```

## How Variadic Parameters Work Internally

When you call a variadic function, Go creates a slice containing the variadic arguments:

```go
func display(items ...string) {
    // items is []string
    fmt.Printf("Type: %T\n", items)
    fmt.Printf("Length: %d\n", len(items))
}

display("a", "b", "c")
// Output:
// Type: []string
// Length: 3
```

**Important:** If no variadic arguments are passed, the slice is empty (`[]T{}`), not `nil`:

```go
func test(nums ...int) {
    fmt.Println(nums == nil)          // false
    fmt.Println(len(nums) == 0)       // true
}

test() // Prints: false, true
```

## Common Patterns

### 1. Aggregation Functions

```go
func min(numbers ...int) int {
    if len(numbers) == 0 {
        return 0 // or return an error
    }
    
    minimum := numbers[0]
    for _, num := range numbers[1:] {
        if num < minimum {
            minimum = num
        }
    }
    return minimum
}
```

### 2. String Formatting

```go
func format(template string, args ...interface{}) string {
    return fmt.Sprintf(template, args...)
}

// Usage
msg := format("User %s has %d points", "Alice", 100)
```

### 3. Logging and Debugging

```go
func log(level string, messages ...string) {
    timestamp := time.Now().Format("15:04:05")
    combined := strings.Join(messages, " ")
    fmt.Printf("[%s] %s: %s\n", timestamp, level, combined)
}

// Usage
log("INFO", "Server", "started", "successfully")
log("ERROR", "Connection failed")
```

### 4. Building Collections

```go
func makeSet(items ...string) map[string]bool {
    set := make(map[string]bool)
    for _, item := range items {
        set[item] = true
    }
    return set
}

// Usage
colors := makeSet("red", "blue", "red", "green")
// map[string]bool{"red": true, "blue": true, "green": true}
```

## Why Variadic Functions Exist

1. **Flexibility:** APIs can accept varying numbers of arguments without overloading
2. **Convenience:** Users don't need to manually create slices for common operations
3. **Readability:** `fmt.Println("a", "b", "c")` is clearer than `fmt.Println([]string{"a", "b", "c"})`
4. **Composability:** Easy to forward arguments to other variadic functions

## When To Use

✅ **Good use cases:**
- Functions naturally accepting variable inputs (sum, max, min)
- Logging and formatting functions
- Building collections or data structures
- Configuration functions with optional parameters

❌ **When NOT to use:**
- When you need type safety for each parameter
- When parameters have different meanings (use a struct instead)
- When you might want to pass `nil` (variadic parameters can't be nil)
- When performance is critical and you're creating many small slices

## When NOT To Use

### Use a Struct Instead

If your parameters have different meanings:

```go
// ❌ Not clear
func createUser(data ...string) User {
    // What's the order? What's required?
}

// ✅ Better
type UserData struct {
    Name     string
    Email    string
    Phone    string
}

func createUser(data UserData) User {
    // Clear and type-safe
}
```

### Use Options Pattern

For complex configuration:

```go
// ✅ Options pattern
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{port: 8080} // defaults
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
server := NewServer(WithPort(3000), WithTimeout(30))
```

## Common Mistakes

### 1. Variadic Parameter Not Last

```go
// ❌ Won't compile
func invalid(nums ...int, name string) {}

// ✅ Correct
func valid(name string, nums ...int) {}
```

### 2. Multiple Variadic Parameters

```go
// ❌ Won't compile
func invalid(names ...string, ages ...int) {}

// ✅ Use different approach
func valid(namesAndAges ...interface{}) {}
// or
func valid(data []NameAge) {} // struct slice
```

### 3. Forgetting the Spread Operator

```go
nums := []int{1, 2, 3}

// ❌ Won't compile (passing slice as single argument)
result := sum(nums)

// ✅ Correct (expanding slice)
result := sum(nums...)
```

### 4. Modifying Variadic Slice

The variadic parameter is a new slice, but modifying its elements affects the original if they're references:

```go
func modify(slices ...[]int) {
    for _, s := range slices {
        s[0] = 999 // Modifies original!
    }
}

original := []int{1, 2, 3}
modify(original)
fmt.Println(original) // [999, 2, 3]
```

### 5. Nil Check Confusion

```go
func process(items ...string) {
    // ❌ items will never be nil
    if items == nil {
        // This never executes
    }
    
    // ✅ Check for empty instead
    if len(items) == 0 {
        // This works
    }
}
```

## Performance Considerations

- Each variadic call allocates a new slice (small overhead)
- For hot paths, consider accepting a slice directly instead
- The spread operator (`...`) doesn't create a new slice, it just passes the existing one

```go
// If called millions of times:
func hotPath(nums ...int) {} // Allocates each time

// Consider:
func hotPath(nums []int) {} // No allocation
```

## Additional Resources

- Official Go Spec on Variadic Functions: https://go.dev/ref/spec#Passing_arguments_to_..._parameters
- Effective Go: https://go.dev/doc/effective_go#functions
- Go Blog - Arrays, slices (and variadic functions): https://go.dev/blog/slices
