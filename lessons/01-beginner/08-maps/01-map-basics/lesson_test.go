package main

import (
	"reflect"
	"testing"
)

func TestBuildScoreboard(t *testing.T) {
	got := BuildScoreboard([]string{"Ava", "Mia"})
	want := map[string]int{"Ava": 0, "Mia": 0}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("BuildScoreboard returned %v", got)
	}
}

func TestLookupScore(t *testing.T) {
	score, ok := LookupScore(map[string]int{"Ava": 12}, "Ava")
	if !ok || score != 12 {
		t.Fatalf("LookupScore returned (%d, %v)", score, ok)
	}

	_, ok = LookupScore(map[string]int{"Ava": 12}, "Mia")
	if ok {
		t.Fatal("LookupScore should report missing key")
	}
}

func TestCountWords(t *testing.T) {
	got := CountWords([]string{"go", "test", "go", "map"})
	want := map[string]int{"go": 2, "test": 1, "map": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CountWords returned %v", got)
	}
}
