# Concurrent Workers Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `DoubleAsync` launches one goroutine per value and stores into a pre-sized result slice by index. See `DoubleAsync` in `solution.go`.
- Exercise 2: `Sum` remains a simple sequential helper because not every operation needs concurrency. See `Sum` in `solution.go`.
- Exercise 3: `DoubleAndSum` composes both helpers to keep the concurrent part isolated. See `DoubleAndSum` in `solution.go`.
