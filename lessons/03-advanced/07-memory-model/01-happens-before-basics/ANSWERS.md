# Happens-Before Basics Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `PublishOnce` shows a concrete safe-publication pattern instead of relying on ad hoc globals.
- `HandOff` uses a channel send and receive to establish visibility of the transferred value.
- `SnapshotLabel` keeps the one-time initialization behind a small wrapper.
