# Mapping With Type Parameters Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `MapSlice` allocates a destination slice once and applies the mapper by index.
- `FirstOrZero` relies on the zero value of `T`, which keeps the API small and idiomatic.
- `PairLabels` composes the earlier generic helper instead of duplicating the loop.
