package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWithRequestTimeoutHasDeadline(t *testing.T) {
	ctx, cancel := WithRequestTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	if _, ok := ctx.Deadline(); !ok {
		t.Fatal("expected derived context to have a deadline")
	}
}

func TestWaitForWorkCompletes(t *testing.T) {
	workDone := make(chan struct{})
	close(workDone)

	err := WaitForWork(context.Background(), workDone)
	if err != nil {
		t.Fatalf("WaitForWork() error = %v, want nil", err)
	}
}

func TestRunWithTimeoutCancelsSlowWork(t *testing.T) {
	workDone := make(chan struct{}) // never closed

	err := RunWithTimeout(context.Background(), 10*time.Millisecond, workDone)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("RunWithTimeout() error = %v, want %v", err, context.DeadlineExceeded)
	}
}
