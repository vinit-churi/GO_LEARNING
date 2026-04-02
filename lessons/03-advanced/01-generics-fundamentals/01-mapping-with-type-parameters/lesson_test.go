package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapSlice(t *testing.T) {
	got := MapSlice([]int{1, 2, 3}, func(v int) string { return string(rune('a' + v - 1)) })
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("MapSlice() = %#v, want %#v", got, want)
	}
}

func TestFirstOrZero(t *testing.T) {
	if got := FirstOrZero([]int{9, 8, 7}); got != 9 {
		t.Fatalf("FirstOrZero(non-empty) = %d", got)
	}
	if got := FirstOrZero([]int{}); got != 0 {
		t.Fatalf("FirstOrZero(empty) = %d", got)
	}
}

func TestPairLabels(t *testing.T) {
	got := PairLabels([]int{2, 4}, func(v int) string { return fmt.Sprintf("v=%d", v) })
	want := []string{"v=2", "v=4"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("PairLabels() = %#v, want %#v", got, want)
	}
}
