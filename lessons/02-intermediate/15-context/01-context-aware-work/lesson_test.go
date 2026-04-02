package main

import (
	"context"
	"errors"
	"testing"
)

func TestTraceID(t *testing.T) {
	ctx := context.WithValue(context.Background(), traceKey{}, "trace-42")
	if got := TraceID(ctx); got != "trace-42" {
		t.Fatalf("TraceID() = %q", got)
	}
}

func TestWorkOrCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := WorkOrCancel(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("WorkOrCancel() error = %v", err)
	}
}

func TestLabelRequest(t *testing.T) {
	ctx := context.WithValue(context.Background(), traceKey{}, "trace-42")
	if got := LabelRequest(ctx, "billing"); got != "trace-42:billing" {
		t.Fatalf("LabelRequest() = %q", got)
	}
}
