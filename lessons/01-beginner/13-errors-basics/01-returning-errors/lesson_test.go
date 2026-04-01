package main

import "testing"

func TestParseQuantity(t *testing.T) {
	got, err := ParseQuantity(" 12 ")
	if err != nil || got != 12 {
		t.Fatalf("ParseQuantity returned (%d, %v)", got, err)
	}
}

func TestSafeDivide(t *testing.T) {
	got, err := SafeDivide(10, 2)
	if err != nil || got != 5 {
		t.Fatalf("SafeDivide returned (%d, %v)", got, err)
	}

	_, err = SafeDivide(10, 0)
	if err == nil {
		t.Fatal("SafeDivide should fail for zero divisor")
	}
}

func TestBuildOrderMessage(t *testing.T) {
	got, err := BuildOrderMessage(" Ava ", 3)
	if err != nil || got != "Ava ordered 3 items" {
		t.Fatalf("BuildOrderMessage returned (%q, %v)", got, err)
	}
}
