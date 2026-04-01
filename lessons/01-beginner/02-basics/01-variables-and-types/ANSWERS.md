# Answers and Explanations

## Exercise 1: Basic Variable Declaration

**Goal:** Learn to use the `var` keyword with explicit types.

### Hint
The pattern is: `var variableName type = value`

### Solution Explanation

```go
func Exercise1() (string, int) {
    var name string = "Alice"
    var age int = 30
    return name, age
}
```

**Key points:**
- We use `var` followed by the variable name
- The type comes after the variable name (this is different from C/Java)
- The `=` assigns the initial value
- We can return multiple values in Go using `(type1, type2)`

**Why this syntax?**
The `var` keyword with explicit types is useful when:
- You want to make the type clear to readers
- You're declaring at package level
- You're declaring a variable without initialization (relying on zero value)

**Alternative approaches:**
You could also write:
```go
var name string
name = "Alice"
var age int
age = 30
```

Or use type inference:
```go
var name = "Alice"
var age = 30
```

See `solution.go` for the reference implementation.

---

## Exercise 2: Short Variable Declaration

**Goal:** Master the `:=` syntax for concise variable declaration.

### Hint
The pattern is: `variableName := value` (type is inferred)

### Solution Explanation

```go
func Exercise2() (float64, bool, string) {
    temperature := 98.6
    isRaining := false
    city := "San Francisco"
    return temperature, isRaining, city
}
```

**Key points:**
- `:=` declares and initializes in one step
- Type is automatically inferred from the value
- `98.6` is inferred as `float64` (Go's default for decimal literals)
- `false` is inferred as `bool`
- `"San Francisco"` is inferred as `string`
- This syntax only works inside functions (not at package level)

**Why this syntax?**
The `:=` operator makes code more concise and is idiomatic Go. It's the most common way to declare variables inside functions.

**Common mistakes:**
- Using `:=` at package level (compile error)
- Using `:=` on an already declared variable in the same scope (error, unless at least one variable on the left is new)

**Type inference rules:**
- Integer literals → `int`
- Floating-point literals → `float64`
- String literals → `string`
- `true`/`false` → `bool`

See `solution.go` for the reference implementation.

---

## Exercise 3: Zero Values

**Goal:** Understand Go's zero value initialization.

### Hint
Declare variables with `var` but don't assign any value. Go will initialize them to their zero value.

### Solution Explanation

```go
func Exercise3() (int, float64, string, bool) {
    var num int
    var decimal float64
    var text string
    var flag bool
    return num, decimal, text, flag
}
```

**Key points:**
- `num` is initialized to `0` (zero value for `int`)
- `decimal` is initialized to `0.0` (zero value for `float64`)
- `text` is initialized to `""` (empty string, zero value for `string`)
- `flag` is initialized to `false` (zero value for `bool`)
- No explicit initialization is needed or provided

**Why zero values?**
Go's zero values eliminate undefined behavior. In languages like C, uninitialized variables contain garbage. In Go, every variable has a predictable, safe default value.

**Important insight:**
Zero values are often useful defaults:
- Counters start at 0
- Booleans default to false (safe default)
- Empty strings are valid strings
- This makes code more robust and predictable

**When to use:**
- When you need a variable but don't have its value yet
- When the zero value is the correct initial value
- When you want to emphasize that initialization happens later

See `solution.go` for the reference implementation.

---

## Exercise 4: Multiple Variable Declarations

**Goal:** Learn to declare multiple variables efficiently using a `var` block.

### Hint
Use the `var ()` block syntax to group related variable declarations together.

### Solution Explanation

```go
func Exercise4() (string, float64, bool, int) {
    var (
        productName string  = "Laptop"
        price       float64 = 999.99
        inStock     bool    = true
        quantity    int     = 5
    )
    return productName, price, inStock, quantity
}
```

**Key points:**
- The `var ()` block allows declaring multiple variables in one place
- Each variable can have a different type
- Variables are aligned for readability (optional but recommended)
- This syntax works at both package and function level
- All related variables (product information) are grouped together

**Why this syntax?**
Grouping related variables improves code organization and readability. It's especially useful for:
- Package-level constants and variables
- Configuration values
- Related state that belongs together

**Alternative approaches:**

Single line (less readable):
```go
var productName, price, inStock, quantity = "Laptop", 999.99, true, 5
```

Multiple `:=` statements (works, but less organized):
```go
productName := "Laptop"
price := 999.99
inStock := true
quantity := 5
```

**Best practice:**
Use `var ()` blocks when:
- Variables are related (like product attributes)
- You have multiple package-level variables
- You want to improve code organization

Use individual `:=` when:
- Variables are unrelated
- They're declared at different points in the function
- You want the most concise syntax

See `solution.go` for the reference implementation.

---

## General Tips

1. **Type inference is your friend:** Use `:=` inside functions for clean, concise code.

2. **Be explicit when needed:** Use `var` with explicit types when the type isn't obvious or when clarity matters.

3. **Zero values are safe:** Unlike C/C++, you can safely use variables declared with `var` without explicit initialization.

4. **Group related declarations:** Use `var ()` blocks to organize related variables, especially at package level.

5. **Experiment:** Try different declaration styles to see what compiles and what doesn't. Understanding the rules through experimentation is valuable.

## Common Pitfalls

### Shadowing Variables
```go
name := "Alice"
if true {
    name := "Bob"  // This creates a NEW variable in the inner scope
    fmt.Println(name)  // Prints "Bob"
}
fmt.Println(name)  // Prints "Alice" (original unchanged)
```

### Unused Variables
Go will not compile if you declare a variable but don't use it:
```go
func example() {
    name := "Alice"  // Compile error: name declared but not used
}
```

### Type Mismatch
You can't mix types without explicit conversion:
```go
var x int = 10
var y int32 = 20
// z := x + y  // Error
z := x + int(y)  // Correct
```

## Next Steps

Now that you understand variables and types:
1. Practice declaring variables in different ways
2. Experiment with type conversions
3. Move on to learning about constants and more complex types (arrays, slices, maps)
4. Read through the official Go Tour's section on [Basic Types](https://go.dev/tour/basics/11)
