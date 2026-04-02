# In-Memory Status Service Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `SetStatus` and `GetStatus` keep map access behind the service boundary.
- `Handler` reads the query parameter, looks up the value, and writes a compact JSON response.
- Using `httptest` ensures the project is exercised through the same API a caller would use.
