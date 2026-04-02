package main

import "testing"

func TestClassifyScore(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{name: "low", score: 10, want: "low"},
		{name: "medium", score: 60, want: "medium"},
		{name: "high", score: 95, want: "high"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClassifyScore(tt.score); got != tt.want {
				t.Fatalf("ClassifyScore(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestIsTerminal(t *testing.T) {
	tests := []struct {
		status string
		want   bool
	}{
		{status: "queued", want: false},
		{status: "done", want: true},
		{status: "failed", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			if got := IsTerminal(tt.status); got != tt.want {
				t.Fatalf("IsTerminal(%q) = %v, want %v", tt.status, got, tt.want)
			}
		})
	}
}

func TestBuildStatusLine(t *testing.T) {
	if got := BuildStatusLine("job-42", 91); got != "job-42:high" {
		t.Fatalf("BuildStatusLine() = %q", got)
	}
}
