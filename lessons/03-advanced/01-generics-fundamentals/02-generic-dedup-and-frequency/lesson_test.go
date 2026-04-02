package main

import (
	"reflect"
	"testing"
)

func TestUniquePreservesFirstSeenOrder(t *testing.T) {
	got := Unique([]string{"go", "go", "rust", "go", "zig", "rust"})
	want := []string{"go", "rust", "zig"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Unique() = %#v, want %#v", got, want)
	}
}

func TestCountBy(t *testing.T) {
	got := CountBy([]int{3, 3, 1, 2, 3, 2})
	want := map[int]int{1: 1, 2: 2, 3: 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CountBy() = %#v, want %#v", got, want)
	}
}

func TestMostFrequent(t *testing.T) {
	got, ok := MostFrequent([]string{"a", "b", "a", "c", "a", "b"})
	if !ok {
		t.Fatal("MostFrequent() ok = false, want true")
	}
	if got != "a" {
		t.Fatalf("MostFrequent() = %q, want %q", got, "a")
	}
}

func TestMostFrequentEmpty(t *testing.T) {
	_, ok := MostFrequent([]string{})
	if ok {
		t.Fatal("MostFrequent() ok = true, want false")
	}
}
