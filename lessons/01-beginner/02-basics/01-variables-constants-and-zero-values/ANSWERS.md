# Answers

## Exercise 1

Construct the struct with only the `Name` field set. The integer and boolean fields can rely on zero values.

## Exercise 2

Subtract `minutesSpent` from `maxDailyPracticeMinutes`, then clamp negative answers to `0`.

## Exercise 3

The session counts as started if either condition is true:

- `session.Minutes > 0`
- `session.Completed`

See `solution.go` for a direct reference implementation.
