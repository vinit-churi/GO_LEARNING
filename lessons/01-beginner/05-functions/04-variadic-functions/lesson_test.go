package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"single number", []int{5}, 5},
		{"multiple numbers", []int{1, 2, 3}, 6},
		{"with negative", []int{-5, 10, -3}, 2},
		{"all zeros", []int{0, 0, 0}, 0},
		{"empty", []int{}, 0},
		{"large numbers", []int{100, 200, 300}, 600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.numbers...)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d; want %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"single number", []int{42}, 42},
		{"multiple numbers", []int{1, 5, 3, 2}, 5},
		{"all negative", []int{-10, -5, -20}, -5},
		{"with zero", []int{-5, 0, 5}, 5},
		{"empty", []int{}, 0},
		{"duplicates", []int{3, 3, 3}, 3},
		{"ascending order", []int{1, 2, 3, 4, 5}, 5},
		{"descending order", []int{5, 4, 3, 2, 1}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.numbers...)
			if result != tt.expected {
				t.Errorf("Max(%v) = %d; want %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestJoinWithSeparator(t *testing.T) {
	tests := []struct {
		name      string
		separator string
		words     []string
		expected  string
	}{
		{"basic join", ", ", []string{"apple", "banana", "cherry"}, "apple, banana, cherry"},
		{"dash separator", " - ", []string{"Go", "is", "awesome"}, "Go - is - awesome"},
		{"single word", "|", []string{"one"}, "one"},
		{"empty words", ", ", []string{}, ""},
		{"two words", " ", []string{"Hello", "World"}, "Hello World"},
		{"pipe separator", "|", []string{"a", "b", "c"}, "a|b|c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JoinWithSeparator(tt.separator, tt.words...)
			if result != tt.expected {
				t.Errorf("JoinWithSeparator(%q, %v) = %q; want %q", 
					tt.separator, tt.words, result, tt.expected)
			}
		})
	}
}

func TestFilterEvens(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{"mixed numbers", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"all odd", []int{1, 3, 5, 7}, []int{}},
		{"all even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"empty", []int{}, []int{}},
		{"with zero", []int{0, 1, 2}, []int{0, 2}},
		{"negative evens", []int{-4, -3, -2, -1, 0}, []int{-4, -2, 0}},
		{"single even", []int{10}, []int{10}},
		{"single odd", []int{11}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterEvens(tt.numbers...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FilterEvens(%v) = %v; want %v", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestConcatenateAll(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			"multiple slices",
			[][]int{{1, 2}, {3, 4}, {5}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"single slice",
			[][]int{{10, 20}},
			[]int{10, 20},
		},
		{
			"empty input",
			[][]int{},
			[]int{},
		},
		{
			"slices with empty",
			[][]int{{1}, {}, {2, 3}},
			[]int{1, 2, 3},
		},
		{
			"all empty slices",
			[][]int{{}, {}, {}},
			[]int{},
		},
		{
			"different lengths",
			[][]int{{1}, {2, 3, 4}, {5, 6}},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConcatenateAll(tt.slices...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ConcatenateAll(%v) = %v; want %v", tt.slices, result, tt.expected)
			}
		})
	}
}

// Test that variadic functions work with spread operator
func TestSpreadOperator(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	
	t.Run("Sum with spread", func(t *testing.T) {
		result := Sum(nums...)
		expected := 15
		if result != expected {
			t.Errorf("Sum(nums...) = %d; want %d", result, expected)
		}
	})
	
	t.Run("Max with spread", func(t *testing.T) {
		result := Max(nums...)
		expected := 5
		if result != expected {
			t.Errorf("Max(nums...) = %d; want %d", result, expected)
		}
	})
	
	t.Run("FilterEvens with spread", func(t *testing.T) {
		result := FilterEvens(nums...)
		expected := []int{2, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FilterEvens(nums...) = %v; want %v", result, expected)
		}
	})
}

// Test edge cases
func TestEdgeCases(t *testing.T) {
	t.Run("Sum with single zero", func(t *testing.T) {
		if result := Sum(0); result != 0 {
			t.Errorf("Sum(0) = %d; want 0", result)
		}
	})
	
	t.Run("Max with all same values", func(t *testing.T) {
		if result := Max(7, 7, 7, 7); result != 7 {
			t.Errorf("Max(7,7,7,7) = %d; want 7", result)
		}
	})
	
	t.Run("JoinWithSeparator with empty strings", func(t *testing.T) {
		result := JoinWithSeparator(",", "", "a", "")
		expected := ",a,"
		if result != expected {
			t.Errorf("JoinWithSeparator with empty strings = %q; want %q", result, expected)
		}
	})
	
	t.Run("FilterEvens preserves order", func(t *testing.T) {
		result := FilterEvens(8, 2, 6, 4)
		expected := []int{8, 2, 6, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FilterEvens should preserve order: got %v; want %v", result, expected)
		}
	})
}

// Benchmark tests
func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(nums...)
	}
}

func BenchmarkMax(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(nums...)
	}
}

func BenchmarkFilterEvens(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterEvens(nums...)
	}
}

func BenchmarkConcatenateAll(b *testing.B) {
	slices := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatenateAll(slices...)
	}
}
