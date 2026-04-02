# Time Formatting and Deadlines

## Difficulty
Intermediate

## Prerequisites
- Standard Library Essentials
- JSON and Encoding

## Learning Objectives
- Format timestamps with Go layouts
- Compute future deadlines with `time.Duration`
- Use UTC in deterministic tests

## Concept Summary

Time handling is easier when you normalize to explicit locations and use the standard `time` APIs. Durations are typed values, which makes arithmetic readable and safe.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: RFC3339 Label
Implement `RFC3339Label(t time.Time) string` using UTC.

### Exercise 2: Extend Deadline
Implement `ExtendDeadline(start time.Time, minutes int) time.Time`.

### Exercise 3: Minutes Remaining
Implement `MinutesRemaining(now, deadline time.Time) int` so it rounds down to whole minutes.
