package main

import "testing"

func TestSummary(t *testing.T) {
	counter := Counter{Name: "Orders", Value: 4}
	if got := counter.Summary(); got != "Orders: 4" {
		t.Fatalf("Summary returned %q", got)
	}
}

func TestAdd(t *testing.T) {
	counter := Counter{Name: "Orders", Value: 4}
	counter.Add(3)
	if counter.Value != 7 {
		t.Fatalf("Add left value at %d", counter.Value)
	}
}

func TestReset(t *testing.T) {
	counter := Counter{Name: "Orders", Value: 4}
	counter.Reset()
	if counter.Value != 0 {
		t.Fatalf("Reset left value at %d", counter.Value)
	}
}
