package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestExercise1(t *testing.T) {
	output := captureOutput(exercise1)
	
	expected := "Hello, World!"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}

func TestExercise2(t *testing.T) {
	output := captureOutput(exercise2)
	
	// Should contain a greeting
	if !strings.Contains(output, "Hello") && !strings.Contains(output, "Hi") {
		t.Error("Expected output to contain a greeting")
	}
	
	// Should contain a welcome message
	if !strings.Contains(output, "Welcome") || !strings.Contains(output, "Go") {
		t.Error("Expected output to mention Go programming")
	}
	
	// Should have multiple lines
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) < 2 {
		t.Errorf("Expected at least 2 lines of output, got %d", len(lines))
	}
}

func TestExercise3(t *testing.T) {
	output := captureOutput(exercise3)
	
	lines := strings.Split(strings.TrimSpace(output), "\n")
	
	// Should have exactly 3 lines
	if len(lines) != 3 {
		t.Errorf("Expected exactly 3 lines of output, got %d", len(lines))
	}
	
	// Each line should have content
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			t.Errorf("Line %d should not be empty", i+1)
		}
	}
}

func TestExercise4(t *testing.T) {
	output := captureOutput(exercise4)
	
	lines := strings.Split(output, "\n")
	
	// ASCII art should have multiple lines
	if len(lines) < 3 {
		t.Errorf("Expected ASCII art with at least 3 lines, got %d", len(lines))
	}
	
	// Should contain some non-alphanumeric characters (typical of ASCII art)
	hasArt := false
	for _, line := range lines {
		if strings.ContainsAny(line, "|_/\\-=+*#@") {
			hasArt = true
			break
		}
	}
	
	if !hasArt {
		t.Error("Expected output to contain ASCII art characters")
	}
}

func TestMainDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()
	
	// Redirect stdout to prevent test output pollution
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()
	
	main()
}
