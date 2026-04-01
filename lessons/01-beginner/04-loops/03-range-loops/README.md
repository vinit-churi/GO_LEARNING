# Range Loops

**Difficulty:** Beginner  
**Prerequisites:** Basic for loops, slices, strings

## Learning Objectives

By the end of this lesson, you will:

- Understand how to use the `range` keyword to iterate over slices
- Learn to iterate over strings and work with runes
- Master using both index and value when ranging
- Know when and how to ignore index or value using underscore `_`
- Understand the difference between ranging over strings as bytes vs runes

## Concept Summary

The `range` keyword in Go provides a convenient way to iterate over elements in various data structures like slices, arrays, maps, and strings. Unlike traditional for loops that require manual index management, `range` handles iteration automatically.

Basic syntax:
```go
for index, value := range collection {
    // use index and value
}
```

The `range` expression returns two values on each iteration:
1. The index/key
2. The value at that position

You can ignore either value using the blank identifier `_`.

## How To Run

From this directory, run:

```bash
go run solution.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

To test with build tags:
```bash
# Test solution code (default)
go test -v

# Test starter code
go test -v -tags=starter
```

## Exercises

### Exercise 1: Sum of Slice Elements

Write a function `SumSlice` that takes a slice of integers and returns their sum using a range loop.

**Function Signature:**
```go
func SumSlice(numbers []int) int
```

**Example:**
```go
SumSlice([]int{1, 2, 3, 4, 5}) // returns 15
SumSlice([]int{10, 20, 30})    // returns 60
SumSlice([]int{})              // returns 0
```

### Exercise 2: Find Index of Element

Write a function `FindIndex` that takes a slice of strings and a target string, returning the index of the first occurrence (or -1 if not found). Use range to iterate.

**Function Signature:**
```go
func FindIndex(items []string, target string) int
```

**Example:**
```go
FindIndex([]string{"apple", "banana", "cherry"}, "banana")  // returns 1
FindIndex([]string{"apple", "banana", "cherry"}, "grape")   // returns -1
FindIndex([]string{}, "anything")                            // returns -1
```

### Exercise 3: Count Vowels in String

Write a function `CountVowels` that counts the number of vowels (a, e, i, o, u) in a string. Use range to iterate over the string as runes to properly handle Unicode characters.

**Function Signature:**
```go
func CountVowels(s string) int
```

**Example:**
```go
CountVowels("hello")           // returns 2 (e, o)
CountVowels("AEIOU")           // returns 5
CountVowels("xyz")             // returns 0
CountVowels("café")            // returns 1 (a only, é is not counted)
```

### Exercise 4: Double Each Element

Write a function `DoubleElements` that takes a slice of integers and returns a new slice with each element doubled. Use range with underscore to ignore the index.

**Function Signature:**
```go
func DoubleElements(numbers []int) []int
```

**Example:**
```go
DoubleElements([]int{1, 2, 3})    // returns []int{2, 4, 6}
DoubleElements([]int{5, 10, 15})  // returns []int{10, 20, 30}
DoubleElements([]int{})           // returns []int{}
```

### Exercise 5: Reverse String Using Runes

Write a function `ReverseString` that reverses a string by iterating over its runes. This ensures proper handling of multi-byte Unicode characters.

**Function Signature:**
```go
func ReverseString(s string) string
```

**Example:**
```go
ReverseString("hello")   // returns "olleh"
ReverseString("Go!")     // returns "!oG"
ReverseString("café")    // returns "éfac"
```

## Tips

- When you only need the value, use `for _, value := range slice` to ignore the index
- When you only need the index, use `for index := range slice` (omit the value entirely)
- Ranging over a string gives you runes (Unicode code points), not bytes
- An empty slice is safe to range over - the loop body simply won't execute
- The index and value are copies; modifying `value` won't change the original slice
