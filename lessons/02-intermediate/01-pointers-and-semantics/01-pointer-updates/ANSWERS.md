# Pointer Updates Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `ApplyDiscount` computes the reduced price and writes it back through the pointer. See `ApplyDiscount` in `solution.go`.
- Exercise 2: `RenameCustomer` trims the input before assigning it so tests do not depend on incidental whitespace. See `RenameCustomer` in `solution.go`.
- Exercise 3: `SnapshotLabel` accepts a value copy on purpose, formats it, and leaves the caller-owned order unchanged. See `SnapshotLabel` in `solution.go`.
