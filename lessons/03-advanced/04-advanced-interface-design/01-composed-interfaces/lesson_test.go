package main

import "testing"

func TestSaveReader(t *testing.T) {
	store := NewMemoryStore()
	if err := SaveReader(store, "alpha"); err != nil {
		t.Fatalf("SaveReader() error = %v", err)
	}
	if got, ok := store.Load("alpha"); !ok || got != "reader:alpha" {
		t.Fatalf("stored value = %q, %v", got, ok)
	}
}

func TestUpsertLabel(t *testing.T) {
	store := NewMemoryStore()
	if got := UpsertLabel(store, "alpha", "first"); got != "first" {
		t.Fatalf("UpsertLabel(initial) = %q", got)
	}
	if got := UpsertLabel(store, "alpha", "second"); got != "first" {
		t.Fatalf("UpsertLabel(existing) = %q", got)
	}
}

func TestCopyIfMissing(t *testing.T) {
	store := NewMemoryStore()
	store.Save("src", "payload")
	if err := CopyIfMissing(store, "src", "dst"); err != nil {
		t.Fatalf("CopyIfMissing() error = %v", err)
	}
	if got, ok := store.Load("dst"); !ok || got != "payload" {
		t.Fatalf("copied value = %q, %v", got, ok)
	}
}
