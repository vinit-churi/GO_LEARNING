# JSON API Client Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `FetchStatus` constructs a request with context, issues it through the provided client, and decodes JSON into `APIStatus`. See `FetchStatus` in `solution.go`.
- Exercise 2: `StatusLabel` keeps display formatting separate from network code. See `StatusLabel` in `solution.go`.
- Exercise 3: `FetchStatusLabel` composes the fetch and formatting steps for a higher-level helper. See `FetchStatusLabel` in `solution.go`.
