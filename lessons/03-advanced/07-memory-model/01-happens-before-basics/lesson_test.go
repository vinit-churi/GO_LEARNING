package main

import "testing"

func TestPublishOnce(t *testing.T) {
	calls := 0
	got := PublishOnce(func() string {
		calls++
		return "ready"
	})
	if got != "ready" {
		t.Fatalf("PublishOnce() = %q", got)
	}
	if calls != 1 {
		t.Fatalf("init calls = %d", calls)
	}
}

func TestHandOff(t *testing.T) {
	if got := HandOff("value"); got != "value" {
		t.Fatalf("HandOff() = %q", got)
	}
}

func TestSnapshotLabel(t *testing.T) {
	if got := SnapshotLabel(func() string { return "ready" }); got != "snapshot:ready" {
		t.Fatalf("SnapshotLabel() = %q", got)
	}
}
