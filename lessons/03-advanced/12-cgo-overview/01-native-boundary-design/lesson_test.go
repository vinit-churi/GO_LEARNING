package main

import "testing"

func TestNativeCall(t *testing.T) {
	if got := NativeCall(2, 3, FakeDriver{}); got != 5 {
		t.Fatalf("NativeCall() = %d", got)
	}
}

func TestSumPair(t *testing.T) {
	if got := SumPair(Pair{Left: 2, Right: 3}, FakeDriver{}); got != 5 {
		t.Fatalf("SumPair() = %d", got)
	}
}

func TestReportPair(t *testing.T) {
	if got := ReportPair(Pair{Left: 2, Right: 3}, FakeDriver{}); got != "sum=5" {
		t.Fatalf("ReportPair() = %q", got)
	}
}
