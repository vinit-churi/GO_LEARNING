# Context-Aware Runner Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `RunOnce` returns quickly on cancellation and otherwise forwards the work result.
- `LabelResult` keeps formatting separate from lifecycle concerns.
- `RunLabeled` composes the two helpers into a small project-like API.
