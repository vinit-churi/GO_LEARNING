package main

import (
	"reflect"
	"testing"
)

func TestCountToN(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "count to 5",
			n:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "count to 1",
			n:        1,
			expected: []int{1},
		},
		{
			name:     "zero returns empty",
			n:        0,
			expected: []int{},
		},
		{
			name:     "negative returns empty",
			n:        -5,
			expected: []int{},
		},
		{
			name:     "count to 10",
			n:        10,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountToN(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountToN(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestMultiplicationTable(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected [][]int
	}{
		{
			name: "size 3",
			size: 3,
			expected: [][]int{
				{1, 2, 3},
				{2, 4, 6},
				{3, 6, 9},
			},
		},
		{
			name: "size 1",
			size: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "size 2",
			size: 2,
			expected: [][]int{
				{1, 2},
				{2, 4},
			},
		},
		{
			name: "size 4",
			size: 4,
			expected: [][]int{
				{1, 2, 3, 4},
				{2, 4, 6, 8},
				{3, 6, 9, 12},
				{4, 8, 12, 16},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MultiplicationTable(tt.size)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MultiplicationTable(%d) = %v, want %v", tt.size, result, tt.expected)
			}
		})
	}
}

func TestFindFirstDivisible(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		divisor  int
		expected int
	}{
		{
			name:     "find first divisible by 5",
			numbers:  []int{3, 5, 10, 15},
			divisor:  5,
			expected: 5,
		},
		{
			name:     "no number divisible",
			numbers:  []int{1, 2, 3},
			divisor:  7,
			expected: -1,
		},
		{
			name:     "first number is divisible",
			numbers:  []int{10, 15, 20},
			divisor:  5,
			expected: 10,
		},
		{
			name:     "last number is divisible",
			numbers:  []int{1, 2, 3, 4, 5},
			divisor:  5,
			expected: 5,
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			divisor:  5,
			expected: -1,
		},
		{
			name:     "divisor is zero",
			numbers:  []int{1, 2, 3},
			divisor:  0,
			expected: -1,
		},
		{
			name:     "multiple divisible numbers",
			numbers:  []int{6, 12, 18, 24},
			divisor:  3,
			expected: 6,
		},
		{
			name:     "divisible by 1",
			numbers:  []int{5, 7, 9},
			divisor:  1,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindFirstDivisible(tt.numbers, tt.divisor)
			if result != tt.expected {
				t.Errorf("FindFirstDivisible(%v, %d) = %d, want %d", tt.numbers, tt.divisor, result, tt.expected)
			}
		})
	}
}

func TestSumOddNumbers(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "mixed odd and even",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 9, // 1 + 3 + 5
		},
		{
			name:     "all even",
			numbers:  []int{2, 4, 6},
			expected: 0,
		},
		{
			name:     "all odd",
			numbers:  []int{1, 3, 5, 7},
			expected: 16, // 1 + 3 + 5 + 7
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "single odd number",
			numbers:  []int{7},
			expected: 7,
		},
		{
			name:     "single even number",
			numbers:  []int{8},
			expected: 0,
		},
		{
			name:     "negative odd numbers",
			numbers:  []int{-1, -3, -5},
			expected: -9, // -1 + -3 + -5
		},
		{
			name:     "mixed positive and negative",
			numbers:  []int{-1, 2, 3, -4, 5},
			expected: 7, // -1 + 3 + 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumOddNumbers(tt.numbers)
			if result != tt.expected {
				t.Errorf("SumOddNumbers(%v) = %d, want %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestReadUntilNegative(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{
			name:     "negative in middle",
			numbers:  []int{1, 2, -1, 3},
			expected: []int{1, 2},
		},
		{
			name:     "no negative",
			numbers:  []int{5, 10, 15},
			expected: []int{5, 10, 15},
		},
		{
			name:     "negative at start",
			numbers:  []int{-1, 2, 3},
			expected: []int{},
		},
		{
			name:     "negative at end",
			numbers:  []int{1, 2, 3, -1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: []int{},
		},
		{
			name:     "single positive",
			numbers:  []int{5},
			expected: []int{5},
		},
		{
			name:     "single negative",
			numbers:  []int{-5},
			expected: []int{},
		},
		{
			name:     "zero is not negative",
			numbers:  []int{0, 1, 2},
			expected: []int{0, 1, 2},
		},
		{
			name:     "multiple negatives",
			numbers:  []int{1, 2, -1, -2, -3},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReadUntilNegative(tt.numbers)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReadUntilNegative(%v) = %v, want %v", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestMainDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()
	
	main()
}
