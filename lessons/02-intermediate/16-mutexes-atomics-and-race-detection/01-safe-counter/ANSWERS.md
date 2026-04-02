# Safe Counter Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `Inc` locks around the mutation so concurrent increments do not race. See `SafeCounter.Inc` in `solution.go`.
- Exercise 2: `Value` also locks because reads must participate in synchronization too. See `SafeCounter.Value` in `solution.go`.
- Exercise 3: `IncrementAtomic` uses `atomic.AddInt64` for a single numeric counter. See `IncrementAtomic` in `solution.go`.
