package main

import "testing"

func TestGreeting(t *testing.T) {
	if got := greeting("Riya"); got != "Hello, Riya!" {
		t.Fatalf("greeting() = %q, want %q", got, "Hello, Riya!")
	}
}

func TestLearnerProfile(t *testing.T) {
	if got := learnerProfile("Sam", 4); got != "Sam has completed 4 lessons." {
		t.Fatalf("learnerProfile() = %q, want %q", got, "Sam has completed 4 lessons.")
	}
}

func TestIsReadyForNextLesson(t *testing.T) {
	testCases := []struct {
		completed int
		want      bool
	}{
		{0, false},
		{2, false},
		{3, true},
		{10, true},
	}

	for _, tc := range testCases {
		if got := isReadyForNextLesson(tc.completed); got != tc.want {
			t.Fatalf("isReadyForNextLesson(%d) = %v, want %v", tc.completed, got, tc.want)
		}
	}
}
