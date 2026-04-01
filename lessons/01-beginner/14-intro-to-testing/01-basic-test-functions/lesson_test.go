package main

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(4) || IsEven(5) {
		t.Fatal("IsEven returned the wrong result")
	}
}

func TestGradeLabel(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{82, "B"},
		{71, "C"},
		{40, "F"},
	}

	for _, tt := range tests {
		if got := GradeLabel(tt.score); got != tt.want {
			t.Fatalf("GradeLabel(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestAverage(t *testing.T) {
	if got := Average([]int{4, 6, 8}); got != 6 {
		t.Fatalf("Average returned %f", got)
	}
}
