package main

import (
	"testing"
	"time"
)

func TestRFC3339Label(t *testing.T) {
	ts := time.Date(2026, time.April, 2, 10, 0, 0, 0, time.FixedZone("IST", 5*60*60+30*60))
	if got := RFC3339Label(ts); got != "2026-04-02T04:30:00Z" {
		t.Fatalf("RFC3339Label() = %q", got)
	}
}

func TestExtendDeadline(t *testing.T) {
	start := time.Date(2026, time.April, 2, 10, 0, 0, 0, time.UTC)
	got := ExtendDeadline(start, 45)
	want := time.Date(2026, time.April, 2, 10, 45, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("ExtendDeadline() = %v, want %v", got, want)
	}
}

func TestMinutesRemaining(t *testing.T) {
	now := time.Date(2026, time.April, 2, 10, 15, 0, 0, time.UTC)
	deadline := time.Date(2026, time.April, 2, 11, 0, 0, 0, time.UTC)
	if got := MinutesRemaining(now, deadline); got != 45 {
		t.Fatalf("MinutesRemaining() = %d", got)
	}
}
