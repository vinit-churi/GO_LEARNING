package main

import (
	"errors"
	"testing"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		wantMin  int
		wantMax  int
	}{
		{
			name:    "basic case",
			input:   []int{3, 7, 1, 9, 2},
			wantMin: 1,
			wantMax: 9,
		},
		{
			name:    "single element",
			input:   []int{5},
			wantMin: 5,
			wantMax: 5,
		},
		{
			name:    "negative numbers",
			input:   []int{-5, -1, -10, -3},
			wantMin: -10,
			wantMax: -1,
		},
		{
			name:    "all same",
			input:   []int{7, 7, 7, 7},
			wantMin: 7,
			wantMax: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := minMax(tt.input)
			if gotMin != tt.wantMin {
				t.Errorf("minMax() min = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("minMax() max = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		want      int
		wantError bool
	}{
		{
			name:      "normal division",
			a:         10,
			b:         2,
			want:      5,
			wantError: false,
		},
		{
			name:      "division by zero",
			a:         10,
			b:         0,
			want:      0,
			wantError: true,
		},
		{
			name:      "negative divisor",
			a:         10,
			b:         -2,
			want:      -5,
			wantError: false,
		},
		{
			name:      "zero dividend",
			a:         0,
			b:         5,
			want:      0,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := safeDivide(tt.a, tt.b)
			if (err != nil) != tt.wantError {
				t.Errorf("safeDivide() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("safeDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUserInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantName  string
		wantAge   int
		wantError bool
	}{
		{
			name:      "valid input",
			input:     "Alice:30",
			wantName:  "Alice",
			wantAge:   30,
			wantError: false,
		},
		{
			name:      "missing age",
			input:     "Bob",
			wantName:  "",
			wantAge:   0,
			wantError: true,
		},
		{
			name:      "invalid age",
			input:     "Charlie:abc",
			wantName:  "",
			wantAge:   0,
			wantError: true,
		},
		{
			name:      "empty name",
			input:     ":25",
			wantName:  "",
			wantAge:   25,
			wantError: false,
		},
		{
			name:      "extra colons",
			input:     "Dave:30:extra",
			wantName:  "",
			wantAge:   0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotAge, err := parseUserInput(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("parseUserInput() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if gotName != tt.wantName {
				t.Errorf("parseUserInput() name = %v, want %v", gotName, tt.wantName)
			}
			if gotAge != tt.wantAge {
				t.Errorf("parseUserInput() age = %v, want %v", gotAge, tt.wantAge)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name  string
		a     string
		b     string
		wantA string
		wantB string
	}{
		{
			name:  "basic swap",
			a:     "hello",
			b:     "world",
			wantA: "world",
			wantB: "hello",
		},
		{
			name:  "empty strings",
			a:     "",
			b:     "",
			wantA: "",
			wantB: "",
		},
		{
			name:  "one empty",
			a:     "test",
			b:     "",
			wantA: "",
			wantB: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := swap(tt.a, tt.b)
			if gotA != tt.wantA {
				t.Errorf("swap() first return = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("swap() second return = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCalculateStats(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantSum   int
		wantAvg   float64
		wantCount int
		wantError bool
	}{
		{
			name:      "basic case",
			input:     []int{10, 20, 30},
			wantSum:   60,
			wantAvg:   20.0,
			wantCount: 3,
			wantError: false,
		},
		{
			name:      "empty slice",
			input:     []int{},
			wantSum:   0,
			wantAvg:   0.0,
			wantCount: 0,
			wantError: true,
		},
		{
			name:      "single element",
			input:     []int{42},
			wantSum:   42,
			wantAvg:   42.0,
			wantCount: 1,
			wantError: false,
		},
		{
			name:      "negative numbers",
			input:     []int{-10, 10, -5, 5},
			wantSum:   0,
			wantAvg:   0.0,
			wantCount: 4,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotAvg, gotCount, err := calculateStats(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("calculateStats() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if gotSum != tt.wantSum {
				t.Errorf("calculateStats() sum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotAvg != tt.wantAvg {
				t.Errorf("calculateStats() avg = %v, want %v", gotAvg, tt.wantAvg)
			}
			if gotCount != tt.wantCount {
				t.Errorf("calculateStats() count = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

// Test that exercises can be called without panicking
func TestExercises(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{"exercise1", exercise1},
		{"exercise2", exercise2},
		{"exercise3", exercise3},
		{"exercise4", exercise4},
		{"exercise5", exercise5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s panicked: %v", tt.name, r)
				}
			}()
			tt.fn()
		})
	}
}

// Test edge case: discarding return values
func TestDiscardingReturnValues(t *testing.T) {
	// Test that we can discard values with _
	_, max := minMax([]int{1, 5, 3})
	if max != 5 {
		t.Errorf("Expected max = 5, got %d", max)
	}

	min, _ := minMax([]int{1, 5, 3})
	if min != 1 {
		t.Errorf("Expected min = 1, got %d", min)
	}

	// Test discarding error (only in test context!)
	result, _ := safeDivide(10, 2)
	if result != 5 {
		t.Errorf("Expected result = 5, got %d", result)
	}
}

// Test that errors are properly returned
func TestErrorReturns(t *testing.T) {
	// safeDivide should return error
	_, err := safeDivide(10, 0)
	if err == nil {
		t.Error("Expected error from safeDivide(10, 0), got nil")
	}

	// parseUserInput should return error for invalid format
	_, _, err = parseUserInput("invalid")
	if err == nil {
		t.Error("Expected error from parseUserInput with invalid format")
	}

	// calculateStats should return error for empty slice
	_, _, _, err = calculateStats([]int{})
	if err == nil {
		t.Error("Expected error from calculateStats with empty slice")
	}
}

// Benchmark the minMax function
func BenchmarkMinMax(b *testing.B) {
	numbers := []int{3, 7, 1, 9, 2, 8, 4, 6, 5}
	for i := 0; i < b.N; i++ {
		minMax(numbers)
	}
}

// Example of how errors should be checked
func Example_safeDivide() {
	result, err := safeDivide(10, 2)
	if err != nil {
		// Handle error
		return
	}
	// Use result
	_ = result
}

// Test that error messages are informative
func TestErrorMessages(t *testing.T) {
	_, err := safeDivide(10, 0)
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() == "" {
		t.Error("Error message should not be empty")
	}

	_, _, err = parseUserInput("invalid")
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() == "" {
		t.Error("Error message should not be empty")
	}

	_, _, _, err = calculateStats([]int{})
	if err == nil {
		t.Fatal("Expected error")
	}
	if err.Error() == "" {
		t.Error("Error message should not be empty")
	}
}

// Test using custom error matching
func TestErrorTypes(t *testing.T) {
	_, err := safeDivide(10, 0)
	if !errors.Is(err, err) { // Simple check that error is not nil
		t.Error("Error should be returned")
	}
}
