package main

import (
	"testing"
	"unsafe"
)

func TestHeaderSize(t *testing.T) {
	if got := HeaderSize(); got != unsafe.Sizeof(Header{}) {
		t.Fatalf("HeaderSize() = %d", got)
	}
}

func TestHeaderAlignment(t *testing.T) {
	if got := HeaderAlignment(); got != unsafe.Alignof(Header{}) {
		t.Fatalf("HeaderAlignment() = %d", got)
	}
}

func TestPayloadOffset(t *testing.T) {
	if got := PayloadOffset(); got != unsafe.Offsetof(Header{}.Payload) {
		t.Fatalf("PayloadOffset() = %d", got)
	}
}
