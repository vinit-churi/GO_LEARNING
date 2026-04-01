# Answers

## Exercise 1

Use `fmt.Sprintf("Hello, %s!", name)` so the punctuation stays consistent.

## Exercise 2

The function should format both the learner name and lesson count into one sentence. See `solution.go` built with the `solution` build tag.

## Exercise 3

This is a direct boolean comparison. The simplest answer is `completed >= 3`.

## Notes

- `starter.go` is the learner-facing file.
- `solution.go` contains a reference implementation behind the `solution` build tag so the default tests stay focused on the learner code.
