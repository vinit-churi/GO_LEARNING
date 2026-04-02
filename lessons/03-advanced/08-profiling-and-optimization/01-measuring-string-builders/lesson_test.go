package main

import "testing"

func TestJoinWithBuilder(t *testing.T) {
	if got := JoinWithBuilder([]string{"api", "v1", "status"}); got != "api/v1/status" {
		t.Fatalf("JoinWithBuilder() = %q", got)
	}
}

func TestJoinWithPlus(t *testing.T) {
	if got := JoinWithPlus([]string{"api", "v1", "status"}); got != "api/v1/status" {
		t.Fatalf("JoinWithPlus() = %q", got)
	}
}

func TestReportLine(t *testing.T) {
	if got := ReportLine([]string{"api", "v1"}); got != "path=api/v1" {
		t.Fatalf("ReportLine() = %q", got)
	}
}

func BenchmarkJoinWithBuilder(b *testing.B) {
	parts := []string{"api", "v1", "status", "health", "ready"}
	for b.Loop() {
		_ = JoinWithBuilder(parts)
	}
}

func BenchmarkJoinWithPlus(b *testing.B) {
	parts := []string{"api", "v1", "status", "health", "ready"}
	for b.Loop() {
		_ = JoinWithPlus(parts)
	}
}
