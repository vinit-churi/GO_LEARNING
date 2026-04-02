# Channel Pipelines Reference

## Concept Explanations

Channels let goroutines communicate through explicit message passing. Small pipeline stages can be easier to reason about than shared mutable state.

## Syntax Patterns

- Keep helpers small and name them for the caller.
- Prefer standard library packages when they already model the behavior you need.
- Use tests to lock down edge cases and make the intended contract explicit.

## Common Mistakes

- Skipping error handling when the lesson API returns an `error`.
- Hiding side effects that should be visible in the function signature.
- Repeating logic instead of composing the helper functions from the lesson.

## Why This Feature Exists

It gives Go programs a small, explicit tool for this part of application design without needing framework magic.

## When To Use It

Use it when the code is clearer, safer, or more testable with this feature.

## When Not To Use It

Do not force the pattern when a simpler sequential or concrete approach would be easier to understand.

## Notes

- Use directional channel types like `<-chan T` to communicate intent clearly.
- Producers should close their output channels when no more values will be sent.
- Consumers range over a channel until it is closed.
- Pipelines work well when each stage does one job and owns its output channel.
- Common mistake: closing a channel from the receiver side or closing it more than once.
- Official docs: Go blog on pipelines and cancellation.
