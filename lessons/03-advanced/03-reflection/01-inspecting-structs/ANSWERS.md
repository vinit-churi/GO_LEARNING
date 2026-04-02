# Inspecting Structs Answers

## Hints

- Start from the smallest helper and compose outward.
- Keep the implementation explicit enough that a reader can trace it quickly.
- Reuse previous functions when the later exercise naturally builds on them.

## Exercise Walkthroughs

- `DescribeStruct` unwraps a pointer, checks that the final value is a struct, and formats the result.
- `FieldNames` iterates over `NumField` and reads each field name from the reflected type.
- `IsZeroValue` defers the comparison rules to `reflect.Value.IsZero`.
