package main

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		expected []int
	}{
		{
			name:     "countdown from 5",
			start:    5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "countdown from 3",
			start:    3,
			expected: []int{3, 2, 1},
		},
		{
			name:     "countdown from 1",
			start:    1,
			expected: []int{1},
		},
		{
			name:     "countdown from 0",
			start:    0,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Countdown(tt.start)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Countdown(%d) = %v, want %v", tt.start, result, tt.expected)
			}
		})
	}
}

func TestSumUntilThreshold(t *testing.T) {
	tests := []struct {
		name      string
		numbers   []int
		threshold int
		expected  int
	}{
		{
			name:      "reaches threshold exactly",
			numbers:   []int{5, 10, 15, 20},
			threshold: 30,
			expected:  30,
		},
		{
			name:      "exceeds threshold",
			numbers:   []int{5, 10, 15, 20},
			threshold: 25,
			expected:  30,
		},
		{
			name:      "threshold higher than sum",
			numbers:   []int{1, 2, 3, 4},
			threshold: 100,
			expected:  10,
		},
		{
			name:      "empty slice",
			numbers:   []int{},
			threshold: 10,
			expected:  0,
		},
		{
			name:      "first element meets threshold",
			numbers:   []int{50, 20, 10},
			threshold: 25,
			expected:  50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumUntilThreshold(tt.numbers, tt.threshold)
			if result != tt.expected {
				t.Errorf("SumUntilThreshold(%v, %d) = %d, want %d", 
					tt.numbers, tt.threshold, result, tt.expected)
			}
		})
	}
}

func TestReadUntilSentinel(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		sentinel string
		expected []string
	}{
		{
			name:     "sentinel in middle",
			values:   []string{"hello", "world", "STOP", "ignored"},
			sentinel: "STOP",
			expected: []string{"hello", "world"},
		},
		{
			name:     "sentinel at start",
			values:   []string{"STOP", "hello", "world"},
			sentinel: "STOP",
			expected: []string{},
		},
		{
			name:     "sentinel at end",
			values:   []string{"hello", "world", "STOP"},
			sentinel: "STOP",
			expected: []string{"hello", "world"},
		},
		{
			name:     "no sentinel",
			values:   []string{"a", "b", "c"},
			sentinel: "STOP",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty slice",
			values:   []string{},
			sentinel: "STOP",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReadUntilSentinel(tt.values, tt.sentinel)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReadUntilSentinel(%v, %q) = %v, want %v",
					tt.values, tt.sentinel, result, tt.expected)
			}
		})
	}
}

func TestFindFirstDivisor(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		start    int
		expected int
	}{
		{
			name:     "find divisor of 20 starting at 3",
			n:        20,
			start:    3,
			expected: 4,
		},
		{
			name:     "prime number",
			n:        17,
			start:    2,
			expected: 17,
		},
		{
			name:     "divisor at start",
			n:        20,
			start:    5,
			expected: 5,
		},
		{
			name:     "start at 1",
			n:        42,
			start:    1,
			expected: 1,
		},
		{
			name:     "find divisor of 100",
			n:        100,
			start:    7,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindFirstDivisor(tt.n, tt.start)
			if result != tt.expected {
				t.Errorf("FindFirstDivisor(%d, %d) = %d, want %d",
					tt.n, tt.start, result, tt.expected)
			}
			// Verify it actually divides n
			if tt.n%result != 0 {
				t.Errorf("FindFirstDivisor(%d, %d) = %d, but %d does not divide %d",
					tt.n, tt.start, result, result, tt.n)
			}
		})
	}
}

func TestCollatzLength(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "collatz of 1",
			n:        1,
			expected: 0,
		},
		{
			name:     "collatz of 2",
			n:        2,
			expected: 1,
		},
		{
			name:     "collatz of 10",
			n:        10,
			expected: 6,
		},
		{
			name:     "collatz of 27",
			n:        27,
			expected: 111,
		},
		{
			name:     "collatz of 16",
			n:        16,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CollatzLength(tt.n)
			if result != tt.expected {
				t.Errorf("CollatzLength(%d) = %d, want %d",
					tt.n, result, tt.expected)
			}
		})
	}
}

func TestCountdownUsesWhileStyle(t *testing.T) {
	// This is a sanity check that the function works
	result := Countdown(5)
	if len(result) != 5 {
		t.Errorf("Countdown(5) should return 5 elements, got %d", len(result))
	}
	if result[0] != 5 || result[4] != 1 {
		t.Errorf("Countdown(5) should start at 5 and end at 1, got %v", result)
	}
}

func TestSumUntilThresholdStopsEarly(t *testing.T) {
	// Verify it doesn't process all elements
	numbers := []int{10, 20, 30, 40, 50}
	result := SumUntilThreshold(numbers, 35)
	
	// Should stop at 10+20+30=60, not process 40 and 50
	if result < 35 {
		t.Errorf("SumUntilThreshold should meet threshold, got %d", result)
	}
	if result > 100 {
		t.Errorf("SumUntilThreshold should stop early, got %d (processed all elements)", result)
	}
}

func TestReadUntilSentinelExcludesSentinel(t *testing.T) {
	values := []string{"a", "b", "STOP", "c"}
	result := ReadUntilSentinel(values, "STOP")
	
	for _, v := range result {
		if v == "STOP" {
			t.Error("ReadUntilSentinel should not include the sentinel value")
		}
	}
}
