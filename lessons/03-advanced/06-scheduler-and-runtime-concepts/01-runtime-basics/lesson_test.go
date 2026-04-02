package main

import (
	"fmt"
	"testing"
)

func TestCPUCount(t *testing.T) {
	if got := CPUCount(); got <= 0 {
		t.Fatalf("CPUCount() = %d", got)
	}
}

func TestChooseWorkers(t *testing.T) {
	if got := ChooseWorkers(1); got != 1 {
		t.Fatalf("ChooseWorkers(1) = %d", got)
	}
	if got := ChooseWorkers(0); got != CPUCount() {
		t.Fatalf("ChooseWorkers(0) = %d", got)
	}
}

func TestParallelSummary(t *testing.T) {
	want := fmt.Sprintf("cpus=%d workers=%d", CPUCount(), ChooseWorkers(2))
	if got := ParallelSummary(2); got != want {
		t.Fatalf("ParallelSummary() = %q, want %q", got, want)
	}
}
