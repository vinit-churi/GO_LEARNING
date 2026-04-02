package main

import (
	"reflect"
	"testing"
)

func TestFeed(t *testing.T) {
	got := Collect(Feed([]int{1, 2, 3}))
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Feed() = %#v, want %#v", got, want)
	}
}

func TestSquareStream(t *testing.T) {
	got := Collect(SquareStream(Feed([]int{1, 2, 3})))
	want := []int{1, 4, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SquareStream() = %#v, want %#v", got, want)
	}
}

func TestCollect(t *testing.T) {
	got := Collect(Feed([]int{4, 5}))
	want := []int{4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Collect() = %#v, want %#v", got, want)
	}
}
