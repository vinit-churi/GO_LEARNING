package main

import (
	"testing"
	"unsafe"
)

func TestRecordSize(t *testing.T) {
	if got := RecordSize(); got != unsafe.Sizeof(Record{}) {
		t.Fatalf("RecordSize() = %d", got)
	}
}

func TestValueOffset(t *testing.T) {
	if got := ValueOffset(); got != unsafe.Offsetof(Record{}.Value) {
		t.Fatalf("ValueOffset() = %d", got)
	}
}

func TestHasPadding(t *testing.T) {
	want := unsafe.Offsetof(Record{}.Value) > unsafe.Sizeof(Record{}.Flag)
	if got := HasPadding(); got != want {
		t.Fatalf("HasPadding() = %v, want %v", got, want)
	}
}
