# Embedded Services

## Difficulty
Intermediate

## Prerequisites
- Structs
- Methods
- Pointers and Semantics

## Learning Objectives
- Compose behavior from smaller structs
- Use embedding to promote fields and methods
- Decide when explicit fields are clearer than embedding

## Concept Summary

Go favors composition over inheritance. Embedding lets one struct reuse fields and methods from another while keeping the relationship explicit in the type definition.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Full Name
Implement `FullName()` on `AuditInfo` using the promoted fields.

### Exercise 2: Service Summary
Implement `Summary()` on `DeployService` so it uses the embedded `AuditInfo`.

### Exercise 3: Touch Audit Data
Implement `Touch(name string)` so it updates the embedded audit owner.
