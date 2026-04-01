# Answers

## Exercise 1

Return a map with one entry per comparison so the learner can inspect all results at once.

## Exercise 2

Use `age >= 0 && age <= 150` to keep the check inclusive on both sides.

## Exercise 3

For weekend detection, compare the day against both allowed values with `||`.

## Exercise 4

The empty check is `s == ""`, so the negated version is `!(s == "")`.

## Exercise 5

Voting requires both conditions, so use `&&`.

## Exercise 6

Bad weather only needs one condition to be true, so use `||`.

## Exercise 7

This function combines equality and range checks. Group the checks clearly even though precedence already works.

## Exercise 8

Premium users always get full access. Otherwise the learner must reason about membership plus age.

## Exercise 9

First define the in-range condition, then negate it with `!` for readability.
