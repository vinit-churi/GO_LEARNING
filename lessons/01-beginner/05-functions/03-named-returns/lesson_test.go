package main

import (
	"errors"
	"strings"
	"testing"
)

// Test Exercise 1: Basic Named Returns
func TestRectangleDimensions(t *testing.T) {
	tests := []struct {
		name      string
		area      int
		perimeter int
		wantLen   int
		wantWidth int
	}{
		{
			name:      "6x4 rectangle",
			area:      24,
			perimeter: 20,
			wantLen:   6,
			wantWidth: 4,
		},
		{
			name:      "4x3 rectangle",
			area:      12,
			perimeter: 14,
			wantLen:   4,
			wantWidth: 3,
		},
		{
			name:      "5x4 rectangle",
			area:      20,
			perimeter: 18,
			wantLen:   5,
			wantWidth: 4,
		},
		{
			name:      "no valid solution",
			area:      10,
			perimeter: 100,
			wantLen:   0,
			wantWidth: 0,
		},
		{
			name:      "square 5x5",
			area:      25,
			perimeter: 20,
			wantLen:   5,
			wantWidth: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length, width := rectangleDimensions(tt.area, tt.perimeter)
			if length != tt.wantLen || width != tt.wantWidth {
				t.Errorf("rectangleDimensions(%d, %d) = (%d, %d), want (%d, %d)",
					tt.area, tt.perimeter, length, width, tt.wantLen, tt.wantWidth)
			}
			
			// Verify the solution if non-zero
			if length != 0 && width != 0 {
				if length*width != tt.area {
					t.Errorf("length(%d) * width(%d) = %d, want area %d",
						length, width, length*width, tt.area)
				}
				if 2*(length+width) != tt.perimeter {
					t.Errorf("2*(length(%d) + width(%d)) = %d, want perimeter %d",
						length, width, 2*(length+width), tt.perimeter)
				}
				if length < width {
					t.Errorf("length(%d) should be >= width(%d)", length, width)
				}
			}
		})
	}
}

// Test Exercise 2: Named Returns with Error Handling
func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name       string
		a, b       float64
		wantResult float64
		wantErr    bool
	}{
		{
			name:       "normal division",
			a:          10,
			b:          2,
			wantResult: 5.0,
			wantErr:    false,
		},
		{
			name:       "division by zero",
			a:          10,
			b:          0,
			wantResult: 0.0,
			wantErr:    true,
		},
		{
			name:       "negative division",
			a:          -15,
			b:          3,
			wantResult: -5.0,
			wantErr:    false,
		},
		{
			name:       "decimal result",
			a:          7,
			b:          2,
			wantResult: 3.5,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := safeDivide(tt.a, tt.b)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("safeDivide(%v, %v) expected error, got nil", tt.a, tt.b)
				}
				if err != nil && !strings.Contains(err.Error(), "division by zero") {
					t.Errorf("error message should contain 'division by zero', got: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("safeDivide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.wantResult {
					t.Errorf("safeDivide(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.wantResult)
				}
			}
		})
	}
}

// Test Exercise 3: When to Use Named Returns
func TestParseNameNamed(t *testing.T) {
	tests := []struct {
		name          string
		fullName      string
		wantFirstName string
		wantLastName  string
		wantErr       bool
	}{
		{
			name:          "normal full name",
			fullName:      "John Doe",
			wantFirstName: "John",
			wantLastName:  "Doe",
			wantErr:       false,
		},
		{
			name:          "single name",
			fullName:      "Madonna",
			wantFirstName: "",
			wantLastName:  "",
			wantErr:       true,
		},
		{
			name:          "three part name",
			fullName:      "Ada Lovelace Byron",
			wantFirstName: "Ada",
			wantLastName:  "Lovelace Byron",
			wantErr:       false,
		},
		{
			name:          "empty string",
			fullName:      "",
			wantFirstName: "",
			wantLastName:  "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstName, lastName, err := parseNameNamed(tt.fullName)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseNameNamed(%q) expected error, got nil", tt.fullName)
				}
				if err != nil && !strings.Contains(err.Error(), "invalid name format") {
					t.Errorf("error message should contain 'invalid name format', got: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("parseNameNamed(%q) unexpected error: %v", tt.fullName, err)
				}
				if firstName != tt.wantFirstName || lastName != tt.wantLastName {
					t.Errorf("parseNameNamed(%q) = (%q, %q), want (%q, %q)",
						tt.fullName, firstName, lastName, tt.wantFirstName, tt.wantLastName)
				}
			}
		})
	}
}

func TestParseNameExplicit(t *testing.T) {
	tests := []struct {
		name          string
		fullName      string
		wantFirstName string
		wantLastName  string
		wantErr       bool
	}{
		{
			name:          "normal full name",
			fullName:      "John Doe",
			wantFirstName: "John",
			wantLastName:  "Doe",
			wantErr:       false,
		},
		{
			name:          "single name",
			fullName:      "Madonna",
			wantFirstName: "",
			wantLastName:  "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstName, lastName, err := parseNameExplicit(tt.fullName)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseNameExplicit(%q) expected error, got nil", tt.fullName)
				}
			} else {
				if err != nil {
					t.Errorf("parseNameExplicit(%q) unexpected error: %v", tt.fullName, err)
				}
				if firstName != tt.wantFirstName || lastName != tt.wantLastName {
					t.Errorf("parseNameExplicit(%q) = (%q, %q), want (%q, %q)",
						tt.fullName, firstName, lastName, tt.wantFirstName, tt.wantLastName)
				}
			}
		})
	}
}

// Test Exercise 4: Complex Named Returns
func TestCalculateStats(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		wantSum int
		wantCnt int
		wantAvg float64
		wantErr bool
	}{
		{
			name:    "even numbers",
			numbers: []int{2, 4, 6, 8},
			wantSum: 20,
			wantCnt: 4,
			wantAvg: 5.0,
			wantErr: false,
		},
		{
			name:    "empty slice",
			numbers: []int{},
			wantSum: 0,
			wantCnt: 0,
			wantAvg: 0.0,
			wantErr: true,
		},
		{
			name:    "single element",
			numbers: []int{10},
			wantSum: 10,
			wantCnt: 1,
			wantAvg: 10.0,
			wantErr: false,
		},
		{
			name:    "negative and positive",
			numbers: []int{-5, 0, 5},
			wantSum: 0,
			wantCnt: 3,
			wantAvg: 0.0,
			wantErr: false,
		},
		{
			name:    "sequential numbers",
			numbers: []int{1, 2, 3, 4, 5},
			wantSum: 15,
			wantCnt: 5,
			wantAvg: 3.0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, count, avg, err := calculateStats(tt.numbers)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("calculateStats(%v) expected error, got nil", tt.numbers)
				}
				if err != nil && !strings.Contains(err.Error(), "empty slice") {
					t.Errorf("error message should contain 'empty slice', got: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("calculateStats(%v) unexpected error: %v", tt.numbers, err)
				}
				if sum != tt.wantSum {
					t.Errorf("sum = %d, want %d", sum, tt.wantSum)
				}
				if count != tt.wantCnt {
					t.Errorf("count = %d, want %d", count, tt.wantCnt)
				}
				if avg != tt.wantAvg {
					t.Errorf("avg = %f, want %f", avg, tt.wantAvg)
				}
			}
		})
	}
}

// Test Exercise 5: Named Returns Best Practices
func TestValidateAndFormat(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOutput string
		wantValid  bool
	}{
		{
			name:       "valid input with spaces",
			input:      "  hello  ",
			wantOutput: "HELLO",
			wantValid:  true,
		},
		{
			name:       "too short",
			input:      "  hi  ",
			wantOutput: "",
			wantValid:  false,
		},
		{
			name:       "exactly 3 chars",
			input:      "  abc  ",
			wantOutput: "ABC",
			wantValid:  true,
		},
		{
			name:       "long valid input",
			input:      "   programming   ",
			wantOutput: "PROGRAMMING",
			wantValid:  true,
		},
		{
			name:       "two chars after trim",
			input:      "  go  ",
			wantOutput: "",
			wantValid:  false,
		},
		{
			name:       "no spaces needed",
			input:      "okay",
			wantOutput: "OKAY",
			wantValid:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, valid := validateAndFormat(tt.input)
			
			if output != tt.wantOutput {
				t.Errorf("output = %q, want %q", output, tt.wantOutput)
			}
			if valid != tt.wantValid {
				t.Errorf("valid = %t, want %t", valid, tt.wantValid)
			}
		})
	}
}

// Test that named returns are actually used in the function signatures
func TestNamedReturnSignatures(t *testing.T) {
	// This is a compile-time check more than runtime
	// If these compile, the signatures are correct
	
	var _ func(int, int) (int, int) = rectangleDimensions
	var _ func(float64, float64) (float64, error) = safeDivide
	var _ func(string) (string, string, error) = parseNameNamed
	var _ func(string) (string, string, error) = parseNameExplicit
	var _ func([]int) (int, int, float64, error) = calculateStats
	var _ func(string) (string, bool) = validateAndFormat
}

// Test that safeDivide returns proper error type
func TestSafeDivideErrorType(t *testing.T) {
	_, err := safeDivide(10, 0)
	if err == nil {
		t.Fatal("expected error for division by zero")
	}
	
	// Verify it's a proper error
	var _ error = err
	
	// Verify we can create our own error for comparison
	expectedErr := errors.New("division by zero")
	if err.Error() != expectedErr.Error() {
		t.Errorf("error message = %q, want %q", err.Error(), expectedErr.Error())
	}
}
