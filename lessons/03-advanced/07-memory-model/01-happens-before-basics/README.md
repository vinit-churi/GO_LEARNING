# Happens-Before Basics

## Difficulty
Advanced

## Prerequisites
- Channels
- Mutexes and Atomics

## Learning Objectives
- Use channels to hand off data safely between goroutines
- Use `sync.Once` for one-time initialization
- Recognize that unsynchronized access is not made safe by timing assumptions

## Concept Summary

The Go memory model explains which synchronization operations make writes visible across goroutines. The practical goal is not memorizing theory, but knowing which tools establish safe ordering.

## Run Instructions

```bash
go run starter.go
```

## Test Instructions

```bash
go test -v
```

## Exercises

### Exercise 1: Publish Once
Implement `PublishOnce(init func() string) string` using `sync.Once`.

### Exercise 2: Hand Off
Implement `HandOff(value string) string` using a channel to transfer the value between goroutines.

### Exercise 3: Snapshot Label
Implement `SnapshotLabel(init func() string) string` using the one-time publisher.
