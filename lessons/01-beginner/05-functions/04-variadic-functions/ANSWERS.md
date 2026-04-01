# Variadic Functions - Answers

## Exercise 1: Sum Function

**Goal:** Create a function that sums any number of integers.

**Hints:**
- Initialize a variable to hold the running total
- Iterate through the variadic parameter (it's a slice!)
- Handle the case where no numbers are provided (should return 0)

**Solution Approach:**

The variadic parameter `numbers ...int` becomes a slice `[]int` inside the function. We iterate through it and accumulate the sum:

```go
func Sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

**Key Concepts:**
- The variadic parameter is accessed like a normal slice
- An empty variadic parameter results in an empty slice, so iterating returns 0 iterations
- No special handling needed for zero arguments—the loop simply doesn't execute

See `solution.go` for the complete implementation.

---

## Exercise 2: Max Function

**Goal:** Find the maximum value from any number of integers.

**Hints:**
- Need to handle the edge case: what if no arguments are provided?
- Start with the first element as the initial maximum
- Compare each subsequent element

**Solution Approach:**

First check if the slice is empty. Then use the first element as the initial max and compare against remaining elements:

```go
func Max(numbers ...int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    max := numbers[0]
    for _, num := range numbers[1:] {
        if num > max {
            max = num
        }
    }
    return max
}
```

**Key Concepts:**
- Always check `len()` when you need at least one element
- `numbers[1:]` skips the first element (we already used it as initial max)
- This works correctly with negative numbers

**Alternative Approach:**

You could also iterate from index 0 but use a separate flag or math.MinInt as the initial value. The above approach is more straightforward.

See `solution.go` for the complete implementation.

---

## Exercise 3: Join With Separator

**Goal:** Combine mixing a regular parameter with a variadic parameter.

**Hints:**
- The separator is a regular parameter (comes first)
- The words are variadic (comes last)
- Handle empty words gracefully
- Think about how to avoid extra separators at the start/end

**Solution Approach:**

Mix regular and variadic parameters. Build the string by iterating through words:

```go
func JoinWithSeparator(separator string, words ...string) string {
    if len(words) == 0 {
        return ""
    }
    
    result := words[0]
    for _, word := range words[1:] {
        result += separator + word
    }
    return result
}
```

**Key Concepts:**
- Variadic parameters **must** be last in the parameter list
- Start with the first word to avoid a leading separator
- Could also use `strings.Join()` but this demonstrates manual building

**Alternative Approach:**

```go
func JoinWithSeparator(separator string, words ...string) string {
    return strings.Join(words, separator)
}
```

The standard library `strings.Join` does exactly this! Always check if the standard library has what you need before reinventing the wheel. However, understanding how to build it yourself is valuable.

See `solution.go` for the complete implementation.

---

## Exercise 4: Filter Evens

**Goal:** Return a filtered slice from variadic input.

**Hints:**
- Create a new slice to hold results
- Use the modulo operator (`%`) to check if a number is even
- `append()` matching elements to the result slice
- Empty input should return an empty slice (not nil)

**Solution Approach:**

Create a result slice and append only even numbers:

```go
func FilterEvens(numbers ...int) []int {
    result := []int{}
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}
```

**Key Concepts:**
- Variadic functions can return slices
- A number is even if `num % 2 == 0`
- Starting with `[]int{}` ensures we return an empty slice, not `nil`
- The result maintains the original order

**Alternative Approach:**

You could pre-allocate capacity if you expect many evens:

```go
result := make([]int, 0, len(numbers)/2)
```

This is a minor optimization that reduces allocations during `append`.

See `solution.go` for the complete implementation.

---

## Exercise 5: Concatenate All (Bonus)

**Goal:** Work with variadic parameters where each argument is itself a slice.

**Hints:**
- The parameter is `...[]int`, meaning multiple slices
- Need to iterate through each slice
- Then iterate through each element in that slice
- Calculate total length first for efficient allocation (optional optimization)

**Solution Approach:**

Iterate through each slice, then each element:

```go
func ConcatenateAll(slices ...[]int) []int {
    // Calculate total length for efficient allocation
    totalLen := 0
    for _, slice := range slices {
        totalLen += len(slice)
    }
    
    // Pre-allocate result
    result := make([]int, 0, totalLen)
    
    // Append all elements
    for _, slice := range slices {
        result = append(result, slice...)
    }
    
    return result
}
```

**Key Concepts:**
- Variadic parameters can be any type, including slices
- Pre-calculating capacity avoids multiple allocations
- `append(result, slice...)` expands the slice elements
- This is essentially a "flatten" operation

**Alternative Approach (Simpler):**

```go
func ConcatenateAll(slices ...[]int) []int {
    result := []int{}
    for _, slice := range slices {
        result = append(result, slice...)
    }
    return result
}
```

This is simpler but may allocate more often. For small inputs, it's fine. For large inputs, pre-allocation helps.

See `solution.go` for the complete implementation.

---

## General Tips

1. **Think of variadic parameters as slices:** Inside the function, they behave exactly like slices
2. **Check for empty input:** If your logic requires at least one element, check `len()`
3. **Use the spread operator (`...`):** When passing a slice to a variadic function
4. **Standard library first:** Check if `strings`, `math`, or other packages have what you need
5. **Consider readability:** Sometimes a slice parameter is clearer than a variadic parameter

## Common Pitfalls

- **Forgetting the spread operator:** `sum(mySlice)` vs `sum(mySlice...)`
- **Putting variadic parameter in the middle:** It must be last
- **Checking for nil:** Variadic parameters are never nil, check `len()` instead
- **Not handling empty input:** Always consider what happens with zero arguments

## When You'd Use These In Real Code

- **Sum/Max/Min:** Utility functions for quick calculations
- **JoinWithSeparator:** Building strings, paths, SQL queries
- **FilterEvens:** Data processing pipelines, data cleaning
- **ConcatenateAll:** Combining results from multiple sources, batch operations

These patterns appear frequently in Go codebases, especially in data processing, CLI tools, and utility libraries.
