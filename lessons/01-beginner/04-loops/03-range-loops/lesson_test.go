package main

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "sum of positive numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "sum with larger numbers",
			numbers:  []int{10, 20, 30},
			expected: 60,
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "single element",
			numbers:  []int{42},
			expected: 42,
		},
		{
			name:     "negative numbers",
			numbers:  []int{-5, -10, -15},
			expected: -30,
		},
		{
			name:     "mixed positive and negative",
			numbers:  []int{10, -5, 20, -10},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSlice(tt.numbers)
			if result != tt.expected {
				t.Errorf("SumSlice(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name     string
		items    []string
		target   string
		expected int
	}{
		{
			name:     "find element in middle",
			items:    []string{"apple", "banana", "cherry"},
			target:   "banana",
			expected: 1,
		},
		{
			name:     "find first element",
			items:    []string{"apple", "banana", "cherry"},
			target:   "apple",
			expected: 0,
		},
		{
			name:     "find last element",
			items:    []string{"apple", "banana", "cherry"},
			target:   "cherry",
			expected: 2,
		},
		{
			name:     "element not found",
			items:    []string{"apple", "banana", "cherry"},
			target:   "grape",
			expected: -1,
		},
		{
			name:     "empty slice",
			items:    []string{},
			target:   "anything",
			expected: -1,
		},
		{
			name:     "duplicate elements - returns first",
			items:    []string{"apple", "banana", "apple"},
			target:   "apple",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindIndex(tt.items, tt.target)
			if result != tt.expected {
				t.Errorf("FindIndex(%v, %q) = %d; expected %d", tt.items, tt.target, result, tt.expected)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "simple word with vowels",
			input:    "hello",
			expected: 2,
		},
		{
			name:     "all vowels uppercase",
			input:    "AEIOU",
			expected: 5,
		},
		{
			name:     "all vowels lowercase",
			input:    "aeiou",
			expected: 5,
		},
		{
			name:     "no vowels",
			input:    "xyz",
			expected: 0,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "mixed case",
			input:    "Hello World",
			expected: 3,
		},
		{
			name:     "unicode with accented vowel",
			input:    "café",
			expected: 1,
		},
		{
			name:     "sentence with spaces",
			input:    "Go is awesome",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountVowels(tt.input)
			if result != tt.expected {
				t.Errorf("CountVowels(%q) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDoubleElements(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{
			name:     "simple numbers",
			numbers:  []int{1, 2, 3},
			expected: []int{2, 4, 6},
		},
		{
			name:     "larger numbers",
			numbers:  []int{5, 10, 15},
			expected: []int{10, 20, 30},
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			numbers:  []int{7},
			expected: []int{14},
		},
		{
			name:     "negative numbers",
			numbers:  []int{-3, -7, -10},
			expected: []int{-6, -14, -20},
		},
		{
			name:     "zeros",
			numbers:  []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DoubleElements(tt.numbers)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DoubleElements(%v) = %v; expected %v", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "with punctuation",
			input:    "Go!",
			expected: "!oG",
		},
		{
			name:     "unicode characters",
			input:    "café",
			expected: "éfac",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "emoji",
			input:    "Go🚀",
			expected: "🚀oG",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Test that main doesn't panic
func TestMainDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()

	main()
}
