package main

import "testing"

func TestCheckPositive(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"positive number", 5, "positive"},
		{"negative number", -3, "not positive"},
		{"zero", 0, "not positive"},
		{"large positive", 1000, "positive"},
		{"large negative", -1000, "not positive"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckPositive(tt.input)
			if result != tt.expected {
				t.Errorf("CheckPositive(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"even number", 4, "even"},
		{"odd number", 7, "odd"},
		{"zero is even", 0, "even"},
		{"negative even", -4, "even"},
		{"negative odd", -7, "odd"},
		{"two is even", 2, "even"},
		{"one is odd", 1, "odd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCompareNumbers(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected string
	}{
		{"a greater than b", 5, 3, "greater"},
		{"a less than b", 2, 7, "less"},
		{"a equal to b", 4, 4, "equal"},
		{"both zero", 0, 0, "equal"},
		{"a negative, b positive", -5, 3, "less"},
		{"a positive, b negative", 5, -3, "greater"},
		{"both negative, a greater", -2, -5, "greater"},
		{"both negative, a less", -5, -2, "less"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareNumbers(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("CompareNumbers(%d, %d) = %q, want %q", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivideAndCheck(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected string
	}{
		{"large result", 100, 5, "large"},
		{"medium result", 20, 4, "medium"},
		{"small result", 5, 10, "small"},
		{"exactly 11 is large", 110, 10, "large"},
		{"exactly 10 is medium", 100, 10, "medium"},
		{"exactly 1 is medium", 10, 10, "medium"},
		{"zero is small", 0, 10, "small"},
		{"negative result is small", -50, 10, "small"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DivideAndCheck(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("DivideAndCheck(%d, %d) = %q, want %q (division result: %d)", 
					tt.a, tt.b, result, tt.expected, tt.a/tt.b)
			}
		})
	}
}

func TestCheckRange(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"in middle of range", 15, "in range"},
		{"below range", 5, "out of range"},
		{"at lower bound", 10, "in range"},
		{"at upper bound", 20, "in range"},
		{"above range", 25, "out of range"},
		{"just below lower bound", 9, "out of range"},
		{"just above upper bound", 21, "out of range"},
		{"far below range", -100, "out of range"},
		{"far above range", 100, "out of range"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckRange(tt.input)
			if result != tt.expected {
				t.Errorf("CheckRange(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests to show performance (optional but educational)
func BenchmarkCheckPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckPositive(42)
	}
}

func BenchmarkIsEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEven(42)
	}
}

func BenchmarkCompareNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareNumbers(42, 21)
	}
}

func BenchmarkDivideAndCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideAndCheck(100, 5)
	}
}

func BenchmarkCheckRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckRange(15)
	}
}
