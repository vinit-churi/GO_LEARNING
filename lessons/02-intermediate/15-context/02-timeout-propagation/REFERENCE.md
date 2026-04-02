# Timeout Propagation Reference

## Concept Explanations

`context.Context` carries cancellation and deadlines across API boundaries. Child contexts inherit cancellation from parents and can add tighter deadlines.

## Syntax Patterns

- `context.WithTimeout(parent, d)` creates a child context and cancel function.
- Use `select` with `<-ctx.Done()` in blocking loops.
- Always `defer cancel()` in the function that creates the child context.

## Common Mistakes

- Forgetting to call `cancel`, causing timers/resources to stay active.
- Ignoring `ctx.Done()` in worker functions.
- Returning generic errors instead of the context cancellation error.

## Why This Feature Exists

Context provides a uniform way to stop work when callers disconnect, deadlines pass, or higher-level workflows fail.

## When To Use It

Use context for request-scoped work, I/O, background tasks, and pipelines that need cancellation.

## When Not To Use It

Do not store optional business parameters in context. Keep it for cancellation, deadlines, and metadata meant for cross-cutting concerns.
