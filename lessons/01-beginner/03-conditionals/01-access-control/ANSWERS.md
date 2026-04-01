# Answers

## Exercise 1

Check inactivity first because it overrides all role-based access messages.

## Exercise 2

This is a good case for early returns:

- premium -> `0`
- subtotal >= 50 -> `0`
- otherwise -> `4.99`

## Exercise 3

Use descending checks: first `>= 90`, then `>= 60`, then the fallback.
