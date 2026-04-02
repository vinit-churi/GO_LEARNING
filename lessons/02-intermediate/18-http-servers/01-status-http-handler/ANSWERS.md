# Status HTTP Handler Answers

## Hints

- Start with the smallest helper before composing the later exercises.
- Keep outputs stable so tests read cleanly.
- Reuse the earlier functions instead of duplicating logic.

## Exercise Walkthroughs

- Exercise 1: `WriteJSON` sets the JSON content type and encodes the provided status value. See `WriteJSON` in `solution.go`.
- Exercise 2: `StatusHandler` closes over the service name and returns a handler function. See `StatusHandler` in `solution.go`.
- Exercise 3: `RouteLabel` keeps request labeling independent from the actual response logic. See `RouteLabel` in `solution.go`.
