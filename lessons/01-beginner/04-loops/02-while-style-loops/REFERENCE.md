# While-Style Loops - Reference

## The For-With-Condition Pattern

In Go, there is no `while` keyword. Instead, the `for` loop handles all iteration needs, including while-style loops:

```go
for condition {
    // executes while condition is true
}
```

This is equivalent to `while (condition)` in other languages.

### Comparison with Traditional For Loop

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style for loop
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

Both accomplish the same task, but the while-style is clearer when the loop logic doesn't fit the init-condition-post pattern.

## When To Use While-Style Loops

### Perfect Use Cases

1. **Unknown iteration count:** When you don't know beforehand how many times to loop
   ```go
   for sum < threshold {
       sum += nextValue()
   }
   ```

2. **Sentinel-based reading:** Processing data until a special marker
   ```go
   for line != "END" {
       process(line)
       line = readLine()
   }
   ```

3. **Condition-based termination:** When exit depends on complex logic
   ```go
   for !isComplete() {
       doWork()
   }
   ```

### When NOT To Use

1. **Known iteration count:** Use traditional for with counter
   ```go
   // Less clear
   i := 0
   for i < len(items) {
       process(items[i])
       i++
   }
   
   // Better
   for i := 0; i < len(items); i++ {
       process(items[i])
   }
   ```

2. **Iterating collections:** Use range instead
   ```go
   // Less idiomatic
   i := 0
   for i < len(items) {
       process(items[i])
       i++
   }
   
   // More idiomatic
   for i, item := range items {
       process(item)
   }
   ```

## Infinite Loops

An empty `for` creates an infinite loop:

```go
for {
    // runs forever unless broken
    if shouldStop() {
        break
    }
}
```

This is equivalent to `while (true)` and is commonly used in:
- Server main loops
- Event processing
- Workers waiting for tasks

## Loop Control Statements

### break

Exits the loop immediately:

```go
for sum < 100 {
    sum += getValue()
    if sum > threshold {
        break  // exit loop early
    }
}
```

### continue

Skips to the next iteration:

```go
for num < 100 {
    num++
    if num%2 == 0 {
        continue  // skip even numbers
    }
    process(num)
}
```

## Sentinel Values

A sentinel is a special value that signals "end of data":

```go
func ReadUntilEmpty(scanner *bufio.Scanner) []string {
    var lines []string
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {  // empty line is sentinel
            break
        }
        lines = append(lines, line)
    }
    return lines
}
```

Common sentinels:
- Empty string `""`
- Special keywords (`"END"`, `"QUIT"`, `"STOP"`)
- Negative numbers when only positive are valid
- `nil` for pointers
- `0` for counts

## Common Patterns

### Accumulator Pattern

```go
sum := 0
for value < limit {
    sum += value
    value = getNext()
}
```

### Search Pattern

```go
found := false
for !found && hasMore() {
    item := getNext()
    if matches(item) {
        found = true
    }
}
```

### Reader Pattern

```go
for scanner.Scan() {
    line := scanner.Text()
    if line == sentinel {
        break
    }
    process(line)
}
```

## Common Mistakes

1. **Forgetting to update condition:** Leads to infinite loops
   ```go
   // WRONG - infinite loop!
   x := 0
   for x < 10 {
       fmt.Println(x)
       // forgot to increment x
   }
   ```

2. **Off-by-one errors:** Especially with > vs >=
   ```go
   // Might miss the target value
   for value < target {  // should be <= ?
       value++
   }
   ```

3. **Complex conditions:** Hard to reason about when they'll be false
   ```go
   // Hard to verify this terminates
   for a < b && c > d && !flag {
       // complex logic
   }
   ```

4. **Side effects in conditions:** Can be confusing
   ```go
   // Calling function in condition
   for readNext() != nil {  // side effect happens on each check
       process()
   }
   ```

## Performance Considerations

- While-style loops have no inherent performance difference from traditional for loops
- The condition is checked before each iteration, so complex condition logic can add overhead
- Keep condition checks simple and fast
- Move invariant computations outside the loop

## Why Go Uses For For Everything

Go's designers chose to have one loop construct (`for`) that can handle all looping needs:

1. **Simplicity:** One concept to learn instead of three (for, while, do-while)
2. **Consistency:** All loops use the same keyword
3. **Flexibility:** Different forms of `for` cover all use cases
4. **No redundancy:** Why have multiple keywords that do similar things?

This is part of Go's philosophy of minimalism and orthogonality.

## Additional Resources

- Go Language Specification - For statements: https://go.dev/ref/spec#For_statements
- Effective Go - Control structures: https://go.dev/doc/effective_go#control-structures
- Go by Example - For loop: https://gobyexample.com/for
