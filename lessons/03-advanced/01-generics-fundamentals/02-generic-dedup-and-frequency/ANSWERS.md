# Generic Dedup and Frequency Answers

## Hints

- Track seen values in a map and append only the first time.
- Build `CountBy` first; `MostFrequent` can reuse it.
- For ties in `MostFrequent`, keep the first item that reached the max.

## Exercise Walkthroughs

- Exercise 1: `Unique` uses `seen := make(map[T]struct{})` and preserves input order by appending on first encounter.
- Exercise 2: `CountBy` increments a `map[T]int` in one pass.
- Exercise 3: `MostFrequent` scans items in input order and compares counts so ties remain deterministic.
