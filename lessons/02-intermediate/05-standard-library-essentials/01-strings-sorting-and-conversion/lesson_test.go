package main

import (
	"reflect"
	"testing"
)

func TestCleanTags(t *testing.T) {
	got := CleanTags([]string{" Go ", "testing", "", " API "})
	want := []string{"go", "testing", "api"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CleanTags() = %#v, want %#v", got, want)
	}
}

func TestSortedScores(t *testing.T) {
	got, err := SortedScores([]string{"9", " 2", "5"})
	if err != nil {
		t.Fatalf("SortedScores() error = %v", err)
	}
	want := []int{2, 5, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SortedScores() = %#v, want %#v", got, want)
	}
}

func TestJoinTags(t *testing.T) {
	if got := JoinTags([]string{" Go ", "testing", " API "}); got != "go,testing,api" {
		t.Fatalf("JoinTags() = %q", got)
	}
}
