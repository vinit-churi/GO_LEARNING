package main

import "testing"

func TestNewPracticeSession(t *testing.T) {
	session := newPracticeSession("maps")

	if session.Name != "maps" {
		t.Fatalf("Name = %q, want %q", session.Name, "maps")
	}
	if session.Minutes != 0 {
		t.Fatalf("Minutes = %d, want 0", session.Minutes)
	}
	if session.Completed {
		t.Fatalf("Completed = true, want false")
	}
}

func TestRemainingPracticeMinutes(t *testing.T) {
	testCases := []struct {
		spent int
		want  int
	}{
		{0, 90},
		{35, 55},
		{90, 0},
		{120, 0},
	}

	for _, tc := range testCases {
		if got := remainingPracticeMinutes(tc.spent); got != tc.want {
			t.Fatalf("remainingPracticeMinutes(%d) = %d, want %d", tc.spent, got, tc.want)
		}
	}
}

func TestHasStarted(t *testing.T) {
	testCases := []struct {
		session PracticeSession
		want    bool
	}{
		{PracticeSession{Name: "intro"}, false},
		{PracticeSession{Name: "loops", Minutes: 10}, true},
		{PracticeSession{Name: "functions", Completed: true}, true},
	}

	for _, tc := range testCases {
		if got := hasStarted(tc.session); got != tc.want {
			t.Fatalf("hasStarted(%+v) = %v, want %v", tc.session, got, tc.want)
		}
	}
}
