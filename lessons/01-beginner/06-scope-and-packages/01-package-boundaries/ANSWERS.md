# Answers

## Exercise 1
Call the exported `labelutil.Header` function directly. The uppercase `H` matters because exported names start with uppercase letters.

## Exercise 2
Create temporary variables with `strings.TrimSpace` and `strings.ToUpper`, then combine them with `fmt.Sprintf`.

## Exercise 3
Use a short-lived variable inside an `if` statement so the trimmed value stays local to the check.
