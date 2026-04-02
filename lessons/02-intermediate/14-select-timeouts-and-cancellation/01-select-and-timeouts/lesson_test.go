package main

import (
	"testing"
	"time"
)

func TestReceiveWithTimeout(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "ready"
	if got := ReceiveWithTimeout(ch, 10*time.Millisecond); got != "ready" {
		t.Fatalf("ReceiveWithTimeout() = %q", got)
	}
}

func TestImmediateStatus(t *testing.T) {
	ch := make(chan string)
	if got := ImmediateStatus(ch); got != "idle" {
		t.Fatalf("ImmediateStatus() = %q", got)
	}
}

func TestFirstResult(t *testing.T) {
	left := make(chan string)
	right := make(chan string, 1)
	right <- "fast"
	if got := FirstResult(left, right); got != "fast" {
		t.Fatalf("FirstResult() = %q", got)
	}
}
