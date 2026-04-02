package main

import (
	"context"
	"errors"
	"testing"
)

func TestRunOnce(t *testing.T) {
	got, err := RunOnce(context.Background(), func() string { return "ready" })
	if err != nil || got != "ready" {
		t.Fatalf("RunOnce() = %q, %v", got, err)
	}
}

func TestLabelResult(t *testing.T) {
	if got := LabelResult("ready"); got != "result:ready" {
		t.Fatalf("LabelResult() = %q", got)
	}
}

func TestRunLabeledCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := RunLabeled(ctx, func() string { return "ready" })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("RunLabeled() error = %v", err)
	}
}
