# Generic Safe Cache Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `Set` and `Get` both take the lock because map access is not safe without synchronization.
- `Size` keeps read access consistent with writes by using the same mutex.
- The generic key type is constrained to `comparable` because Go maps require it.
