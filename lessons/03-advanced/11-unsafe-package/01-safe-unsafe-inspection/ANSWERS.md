# Safe Unsafe Inspection Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `HeaderSize`, `HeaderAlignment`, and `PayloadOffset` all report compiler-chosen layout facts.
- The lesson intentionally stops before pointer arithmetic so the risk stays low and the examples remain portable.
- Use these results as debugging aids, not as excuses to optimize blindly.
