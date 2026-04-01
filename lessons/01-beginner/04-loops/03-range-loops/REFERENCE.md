# Range Loops - Reference

## What is Range?

The `range` keyword is Go's idiomatic way to iterate over elements in data structures. It works with:
- Arrays and slices
- Strings
- Maps
- Channels (covered in advanced lessons)

Range automatically handles the iteration logic, making code cleaner and less error-prone than manual index management.

## Syntax Patterns

### Basic Range Over Slice

```go
numbers := []int{10, 20, 30, 40, 50}

// Get both index and value
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
// Output:
// Index: 0, Value: 10
// Index: 1, Value: 20
// ...
```

### Ignore Index with Underscore

```go
numbers := []int{10, 20, 30}

// Only need the value
for _, value := range numbers {
    fmt.Println(value)
}
```

The blank identifier `_` tells Go you're not using that variable. This is required because Go doesn't allow unused variables.

### Ignore Value (Index Only)

```go
numbers := []int{10, 20, 30}

// Only need the index
for index := range numbers {
    fmt.Printf("Position: %d\n", index)
}
```

Note: When you only need the index, you can simply omit the second variable entirely.

### Range Over String (Runes)

```go
text := "Go 🚀"

// Iterate over runes (Unicode code points)
for index, runeValue := range text {
    fmt.Printf("Index: %d, Rune: %c, Value: %d\n", index, runeValue, runeValue)
}
// Output:
// Index: 0, Rune: G, Value: 71
// Index: 1, Rune: o, Value: 111
// Index: 2, Rune:  , Value: 32
// Index: 3, Rune: 🚀, Value: 128640
```

**Important:** The index might not increase by 1 each iteration because multi-byte UTF-8 characters take multiple bytes. The index represents the byte offset, not the rune position.

### Range Over String (Bytes)

If you need to iterate over bytes instead of runes:

```go
text := "hello"

// Convert to byte slice first
for index, byteValue := range []byte(text) {
    fmt.Printf("Index: %d, Byte: %d\n", index, byteValue)
}
```

### Creating New Slices During Range

```go
numbers := []int{1, 2, 3, 4, 5}
var evens []int

for _, num := range numbers {
    if num%2 == 0 {
        evens = append(evens, num)
    }
}
// evens is now []int{2, 4}
```

## Why Range Exists

### Problems with Traditional Loops

```go
// Traditional loop - more error-prone
for i := 0; i < len(numbers); i++ {
    value := numbers[i]
    // use value
}
```

Issues:
- Easy to make off-by-one errors
- More verbose
- Manually managing index
- Risk of index out of bounds

### Benefits of Range

```go
// Range loop - cleaner and safer
for _, value := range numbers {
    // use value
}
```

Benefits:
- Cannot have index out of bounds
- Less code to write
- Clearer intent
- Go handles iteration logic

## Common Mistakes

### 1. Modifying Value Doesn't Change Original

```go
numbers := []int{1, 2, 3}

// This does NOT modify the slice!
for _, num := range numbers {
    num = num * 2  // Changes local copy only
}
// numbers is still []int{1, 2, 3}
```

**Fix:** Use the index to modify:

```go
for i := range numbers {
    numbers[i] = numbers[i] * 2
}
// numbers is now []int{2, 4, 6}
```

### 2. Reusing Loop Variables

```go
var pointers []*int

for _, num := range []int{1, 2, 3} {
    pointers = append(pointers, &num)  // BUG! All point to same variable
}
```

**Fix:** Create a new variable:

```go
for _, num := range []int{1, 2, 3} {
    num := num  // Create new variable
    pointers = append(pointers, &num)
}
```

### 3. Forgetting Strings Use Runes

```go
text := "café"
fmt.Println(len(text))  // prints 5 (bytes)

count := 0
for range text {
    count++
}
fmt.Println(count)  // prints 4 (runes)
```

### 4. Ignoring Return Values Incorrectly

```go
// Wrong - trying to ignore both values
for _ := range numbers {  // Syntax error
}

// Correct ways:
for range numbers {        // Ignore both (just iterate)
for i := range numbers {   // Use index only
for _, v := range numbers { // Use value only
for i, v := range numbers { // Use both
```

## When To Use Range

**✅ Use Range When:**
- Iterating over all elements in a collection
- You need both index and value
- You want cleaner, more readable code
- Working with strings as Unicode characters

**❌ Don't Use Range When:**
- You need to modify the loop index during iteration
- You need to iterate with a custom step (e.g., by 2s)
- You're iterating over a numeric range not backed by a collection
- You need to iterate backwards (use traditional for loop)

## Examples in Context

### Filter Slice

```go
func FilterEven(numbers []int) []int {
    var result []int
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}
```

### Find Maximum

```go
func FindMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}
```

### String Contains Character

```go
func ContainsChar(s string, target rune) bool {
    for _, char := range s {
        if char == target {
            return true
        }
    }
    return false
}
```

## Performance Considerations

- Range is as fast as traditional for loops (compiler optimizes them similarly)
- Range creates a copy of the value on each iteration (usually cheap for basic types)
- For large structs, consider ranging over pointers or indices

## Additional Resources

- Go Spec on Range: https://go.dev/ref/spec#For_statements
- Effective Go on For: https://go.dev/doc/effective_go#for
- Go Blog on Strings: https://go.dev/blog/strings
