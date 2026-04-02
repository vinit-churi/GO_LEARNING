# Ordered and Numeric Constraints Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `SumNumbers` loops over the slice and accumulates into a zero-value total of the constrained type.
- `MaxOrdered` returns the zero value for empty input to keep the function total and simple for exercises.
- `ClampOrdered` uses direct comparisons because the ordered constraint guarantees them.
