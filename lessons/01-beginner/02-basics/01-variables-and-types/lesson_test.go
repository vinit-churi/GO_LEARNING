package variables

import (
	"testing"
)

func TestExercise1(t *testing.T) {
	name, age := Exercise1()

	// Check that name is not empty
	if name == "" {
		t.Error("Exercise1: name should not be empty")
	}

	// Check that age is positive
	if age <= 0 {
		t.Error("Exercise1: age should be a positive number")
	}

	// Check types (compilation already validates this, but we check values)
	if len(name) < 1 {
		t.Error("Exercise1: name should have at least one character")
	}

	t.Logf("Exercise1 passed: name=%q, age=%d", name, age)
}

func TestExercise2(t *testing.T) {
	temperature, isRaining, city := Exercise2()

	// Check that temperature is reasonable
	if temperature <= 0 || temperature > 120 {
		t.Errorf("Exercise2: temperature should be reasonable (0-120), got %f", temperature)
	}

	// Check that city is not empty
	if city == "" {
		t.Error("Exercise2: city should not be empty")
	}

	// Type checking is done at compile time, but we verify values make sense
	if temperature == 0.0 {
		t.Error("Exercise2: temperature should not be zero")
	}

	// isRaining can be true or false, both are valid
	t.Logf("Exercise2 passed: temperature=%f, isRaining=%t, city=%q", temperature, isRaining, city)
}

func TestExercise3(t *testing.T) {
	num, decimal, text, flag := Exercise3()

	// Test for zero values
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"int zero value", num, 0},
		{"float64 zero value", decimal, 0.0},
		{"string zero value", text, ""},
		{"bool zero value", flag, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("Exercise3 %s: got %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}

	t.Logf("Exercise3 passed: all zero values correct (int=%d, float64=%f, string=%q, bool=%t)",
		num, decimal, text, flag)
}

func TestExercise4(t *testing.T) {
	productName, price, inStock, quantity := Exercise4()

	// Define expected values
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"productName", productName, "Laptop"},
		{"price", price, 999.99},
		{"inStock", inStock, true},
		{"quantity", quantity, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("Exercise4 %s: got %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}

	t.Logf("Exercise4 passed: productName=%q, price=%f, inStock=%t, quantity=%d",
		productName, price, inStock, quantity)
}

// TestVariableDeclarationStyles tests different declaration styles to ensure learners understand them
func TestVariableDeclarationStyles(t *testing.T) {
	t.Run("var with explicit type", func(t *testing.T) {
		var x int = 42
		if x != 42 {
			t.Errorf("expected 42, got %d", x)
		}
	})

	t.Run("var with type inference", func(t *testing.T) {
		var x = 42
		if x != 42 {
			t.Errorf("expected 42, got %d", x)
		}
	})

	t.Run("short variable declaration", func(t *testing.T) {
		x := 42
		if x != 42 {
			t.Errorf("expected 42, got %d", x)
		}
	})

	t.Run("zero value declaration", func(t *testing.T) {
		var x int
		if x != 0 {
			t.Errorf("expected 0 (zero value), got %d", x)
		}
	})

	t.Run("multiple declaration", func(t *testing.T) {
		var (
			a = 1
			b = 2
			c = 3
		)
		if a != 1 || b != 2 || c != 3 {
			t.Errorf("expected 1, 2, 3, got %d, %d, %d", a, b, c)
		}
	})

	t.Run("parallel assignment", func(t *testing.T) {
		x, y, z := 1, 2, 3
		if x != 1 || y != 2 || z != 3 {
			t.Errorf("expected 1, 2, 3, got %d, %d, %d", x, y, z)
		}
	})
}

// TestTypeInference validates that Go correctly infers types
func TestTypeInference(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func() bool
	}{
		{
			name: "integer literal infers to int",
			testFunc: func() bool {
				x := 42
				var y int = x
				return y == 42
			},
		},
		{
			name: "float literal infers to float64",
			testFunc: func() bool {
				x := 3.14
				var y float64 = x
				return y == 3.14
			},
		},
		{
			name: "string literal infers to string",
			testFunc: func() bool {
				x := "hello"
				var y string = x
				return y == "hello"
			},
		},
		{
			name: "bool literal infers to bool",
			testFunc: func() bool {
				x := true
				var y bool = x
				return y == true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.testFunc() {
				t.Errorf("%s failed", tt.name)
			}
		})
	}
}

// BenchmarkVariableDeclaration benchmarks different declaration styles
// (They should all be equally fast - this is educational)
func BenchmarkVariableDeclaration(b *testing.B) {
	b.Run("var with explicit type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var x int = 42
			_ = x
		}
	})

	b.Run("short declaration", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := 42
			_ = x
		}
	})

	b.Run("var with inference", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var x = 42
			_ = x
		}
	})
}
