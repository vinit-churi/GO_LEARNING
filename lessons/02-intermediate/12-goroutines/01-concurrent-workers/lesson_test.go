package main

import (
	"reflect"
	"testing"
)

func TestDoubleAsync(t *testing.T) {
	got := DoubleAsync([]int{1, 2, 3})
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("DoubleAsync() = %#v, want %#v", got, want)
	}
}

func TestSum(t *testing.T) {
	if got := Sum([]int{2, 4, 6}); got != 12 {
		t.Fatalf("Sum() = %d", got)
	}
}

func TestDoubleAndSum(t *testing.T) {
	if got := DoubleAndSum([]int{1, 2, 3}); got != 12 {
		t.Fatalf("DoubleAndSum() = %d", got)
	}
}
