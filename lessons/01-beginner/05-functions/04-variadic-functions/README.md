# Variadic Functions

**Difficulty:** Beginner  
**Prerequisites:** 
- Basic Functions (05-functions/01-basic-functions)
- Slices (03-data-structures/02-slices)

## Learning Objectives

By the end of this lesson, you will:

- Understand variadic parameters and the `...` syntax
- Write functions that accept variable numbers of arguments
- Pass slices to variadic functions using the spread operator
- Combine regular and variadic parameters correctly
- Implement common variadic patterns (sum, max, min, etc.)

## Concept Summary

A **variadic function** accepts zero or more arguments of a specified type. In Go, you declare a variadic parameter using three dots (`...`) before the type. Inside the function, the variadic parameter is treated as a slice.

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Can be called with any number of arguments
sum(1, 2)           // 3
sum(1, 2, 3, 4, 5)  // 15
sum()               // 0
```

You can also pass a slice to a variadic function using `...`:

```go
nums := []int{1, 2, 3, 4}
result := sum(nums...)  // Expands slice into individual arguments
```

## How To Run

From this directory, run:

```bash
# Run the solution
go run solution.go

# Or with starter code
go run -tags starter starter.go
```

## How To Test

From this directory, run:

```bash
go test -v
```

## Exercises

### Exercise 1: Sum Function

Implement a variadic function `Sum` that calculates the sum of any number of integers.

**Function Signature:**
```go
func Sum(numbers ...int) int
```

**Examples:**
```go
Sum(1, 2, 3)        // returns 6
Sum(10, 20)         // returns 30
Sum()               // returns 0
Sum(100)            // returns 100
```

### Exercise 2: Max Function

Implement a variadic function `Max` that returns the maximum value from any number of integers. If no arguments are provided, return 0.

**Function Signature:**
```go
func Max(numbers ...int) int
```

**Examples:**
```go
Max(1, 5, 3, 2)     // returns 5
Max(-10, -5, -20)   // returns -5
Max(42)             // returns 42
Max()               // returns 0
```

### Exercise 3: Join With Separator

Implement a variadic function `JoinWithSeparator` that takes a separator string as the first parameter and any number of strings to join.

**Function Signature:**
```go
func JoinWithSeparator(separator string, words ...string) string
```

**Examples:**
```go
JoinWithSeparator(", ", "apple", "banana", "cherry")  // returns "apple, banana, cherry"
JoinWithSeparator(" - ", "Go", "is", "awesome")       // returns "Go - is - awesome"
JoinWithSeparator("|", "one")                         // returns "one"
JoinWithSeparator(", ")                               // returns ""
```

### Exercise 4: Filter Evens

Implement a variadic function `FilterEvens` that accepts any number of integers and returns a slice containing only the even numbers, in the order they appeared.

**Function Signature:**
```go
func FilterEvens(numbers ...int) []int
```

**Examples:**
```go
FilterEvens(1, 2, 3, 4, 5, 6)     // returns []int{2, 4, 6}
FilterEvens(1, 3, 5, 7)            // returns []int{}
FilterEvens(2, 4, 6)               // returns []int{2, 4, 6}
FilterEvens()                      // returns []int{}
```

### Exercise 5: Concatenate All (Bonus)

Implement a variadic function `ConcatenateAll` that accepts slices of integers (using variadic slices) and returns a single slice containing all elements.

**Function Signature:**
```go
func ConcatenateAll(slices ...[]int) []int
```

**Examples:**
```go
ConcatenateAll([]int{1, 2}, []int{3, 4}, []int{5})   // returns []int{1, 2, 3, 4, 5}
ConcatenateAll([]int{10, 20})                         // returns []int{10, 20}
ConcatenateAll()                                      // returns []int{}
```

## Tips

- A variadic parameter must be the last parameter in a function signature
- Inside the function, treat the variadic parameter as a slice
- Use `slice...` to expand a slice when calling a variadic function
- You can only have one variadic parameter per function
- An empty variadic parameter results in an empty slice, not `nil`
