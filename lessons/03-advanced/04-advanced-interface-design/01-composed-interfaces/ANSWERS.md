# Composed Interfaces Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `SaveReader` shows how a combined interface can still stay small and explicit.
- `UpsertLabel` reads first and falls back to a write path when the value is missing.
- `CopyIfMissing` demonstrates composition without leaking storage details into callers.
