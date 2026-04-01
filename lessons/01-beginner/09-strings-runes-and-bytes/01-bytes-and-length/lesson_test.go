package main

import "testing"

func TestByteLength(t *testing.T) {
	if got := ByteLength("Go语言"); got != 8 {
		t.Fatalf("ByteLength returned %d", got)
	}
}

func TestFirstRune(t *testing.T) {
	if got := FirstRune("你好"); got != '你' {
		t.Fatalf("FirstRune returned %q", got)
	}
}

func TestReverseRunes(t *testing.T) {
	if got := ReverseRunes("Go语言"); got != "言语oG" {
		t.Fatalf("ReverseRunes returned %q", got)
	}
}
