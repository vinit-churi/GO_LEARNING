package main

import (
	"strings"
	"testing"
)

// TestGetDayType tests the day of the week classifier
func TestGetDayType(t *testing.T) {
	tests := []struct {
		name     string
		day      string
		expected string
	}{
		// Weekdays
		{"Monday is weekday", "Monday", "Weekday"},
		{"Tuesday is weekday", "Tuesday", "Weekday"},
		{"Wednesday is weekday", "Wednesday", "Weekday"},
		{"Thursday is weekday", "Thursday", "Weekday"},
		{"Friday is weekday", "Friday", "Weekday"},
		
		// Weekends
		{"Saturday is weekend", "Saturday", "Weekend"},
		{"Sunday is weekend", "Sunday", "Weekend"},
		
		// Invalid
		{"Invalid day name", "Funday", "Invalid day"},
		{"Empty string", "", "Invalid day"},
		{"Lowercase monday", "monday", "Invalid day"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDayType(tt.day)
			if result != tt.expected {
				t.Errorf("GetDayType(%q) = %q, want %q", tt.day, result, tt.expected)
			}
		})
	}
}

// TestGetGradeMessage tests the grade evaluator
func TestGetGradeMessage(t *testing.T) {
	tests := []struct {
		name     string
		grade    string
		expected string
	}{
		{"Grade A", "A", "Excellent work!"},
		{"Grade B", "B", "Excellent work!"},
		{"Grade C", "C", "Good job!"},
		{"Grade D", "D", "You passed, but consider studying more."},
		{"Grade F", "F", "You need to retake this course."},
		{"Invalid grade X", "X", "Invalid grade"},
		{"Lowercase a", "a", "Invalid grade"},
		{"Empty string", "", "Invalid grade"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetGradeMessage(tt.grade)
			if result != tt.expected {
				t.Errorf("GetGradeMessage(%q) = %q, want %q", tt.grade, result, tt.expected)
			}
		})
	}
}

// TestClassifyNumber tests the number classifier
func TestClassifyNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected string
	}{
		// Negative
		{"Negative small", -1, "Negative"},
		{"Negative large", -100, "Negative"},
		{"Negative medium", -50, "Negative"},
		
		// Zero
		{"Zero", 0, "Zero"},
		
		// Small positive
		{"Small positive 1", 1, "Small positive"},
		{"Small positive 5", 5, "Small positive"},
		{"Small positive 10", 10, "Small positive"},
		
		// Medium positive
		{"Medium positive 11", 11, "Medium positive"},
		{"Medium positive 50", 50, "Medium positive"},
		{"Medium positive 100", 100, "Medium positive"},
		
		// Large positive
		{"Large positive 101", 101, "Large positive"},
		{"Large positive 1000", 1000, "Large positive"},
		{"Large positive 1000000", 1000000, "Large positive"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClassifyNumber(tt.number)
			if result != tt.expected {
				t.Errorf("ClassifyNumber(%d) = %q, want %q", tt.number, result, tt.expected)
			}
		})
	}
}

// TestDescribeSeason tests the season describer with fallthrough
func TestDescribeSeason(t *testing.T) {
	tests := []struct {
		name            string
		season          string
		expectedContains []string // What strings should be in the result
		shouldNotContain []string // What strings should NOT be in the result
	}{
		{
			name:            "Spring",
			season:          "Spring",
			expectedContains: []string{"Flowers bloom"},
			shouldNotContain: []string{"Hot and sunny", "Leaves change", "Cold and snowy"},
		},
		{
			name:            "Summer with fallthrough",
			season:          "Summer",
			expectedContains: []string{"Hot and sunny", "Leaves change color"},
			shouldNotContain: []string{"Flowers bloom", "Cold and snowy"},
		},
		{
			name:            "Fall",
			season:          "Fall",
			expectedContains: []string{"Leaves change color"},
			shouldNotContain: []string{"Hot and sunny", "Flowers bloom", "Cold and snowy"},
		},
		{
			name:            "Autumn (synonym for Fall)",
			season:          "Autumn",
			expectedContains: []string{"Leaves change color"},
			shouldNotContain: []string{"Hot and sunny", "Flowers bloom", "Cold and snowy"},
		},
		{
			name:            "Winter",
			season:          "Winter",
			expectedContains: []string{"Cold and snowy"},
			shouldNotContain: []string{"Hot and sunny", "Leaves change", "Flowers bloom"},
		},
		{
			name:            "Invalid season",
			season:          "Rainy",
			expectedContains: []string{"Unknown season"},
			shouldNotContain: []string{"Hot and sunny", "Leaves change", "Flowers bloom", "Cold and snowy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DescribeSeason(tt.season)
			
			// Check for expected strings
			for _, expected := range tt.expectedContains {
				if !strings.Contains(result, expected) {
					t.Errorf("DescribeSeason(%q) = %q, should contain %q", tt.season, result, expected)
				}
			}
			
			// Check for strings that should NOT be present
			for _, notExpected := range tt.shouldNotContain {
				if strings.Contains(result, notExpected) {
					t.Errorf("DescribeSeason(%q) = %q, should NOT contain %q", tt.season, result, notExpected)
				}
			}
		})
	}
	
	// Special test for Summer fallthrough behavior
	t.Run("Summer fallthrough creates multiline result", func(t *testing.T) {
		result := DescribeSeason("Summer")
		lines := strings.Split(result, "\n")
		if len(lines) < 2 {
			t.Errorf("DescribeSeason(\"Summer\") should produce multiple lines due to fallthrough, got %q", result)
		}
	})
}

// TestClassifyHTTPStatus tests HTTP status code classification
func TestClassifyHTTPStatus(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected string
	}{
		// Success (2xx)
		{"200 OK", 200, "Success"},
		{"201 Created", 201, "Success"},
		{"204 No Content", 204, "Success"},
		{"299 edge case", 299, "Success"},
		
		// Redirection (3xx)
		{"300 Multiple Choices", 300, "Redirection"},
		{"301 Moved Permanently", 301, "Redirection"},
		{"302 Found", 302, "Redirection"},
		{"304 Not Modified", 304, "Redirection"},
		{"399 edge case", 399, "Redirection"},
		
		// Client Error (4xx)
		{"400 Bad Request", 400, "Client Error"},
		{"401 Unauthorized", 401, "Client Error"},
		{"403 Forbidden", 403, "Client Error"},
		{"404 Not Found", 404, "Client Error"},
		{"418 I'm a teapot", 418, "Client Error"},
		{"499 edge case", 499, "Client Error"},
		
		// Server Error (5xx)
		{"500 Internal Server Error", 500, "Server Error"},
		{"501 Not Implemented", 501, "Server Error"},
		{"502 Bad Gateway", 502, "Server Error"},
		{"503 Service Unavailable", 503, "Server Error"},
		{"599 edge case", 599, "Server Error"},
		
		// Unknown
		{"100 Continue", 100, "Unknown Status"},
		{"199 edge case", 199, "Unknown Status"},
		{"600 Invalid", 600, "Unknown Status"},
		{"999 Invalid", 999, "Unknown Status"},
		{"0 Invalid", 0, "Unknown Status"},
		{"-1 Invalid", -1, "Unknown Status"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClassifyHTTPStatus(tt.code)
			if result != tt.expected {
				t.Errorf("ClassifyHTTPStatus(%d) = %q, want %q", tt.code, result, tt.expected)
			}
		})
	}
}

// TestMainDoesNotPanic ensures main runs without crashing
func TestMainDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()
	
	main()
}

// Benchmark tests (optional, but good practice)
func BenchmarkGetDayType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDayType("Monday")
	}
}

func BenchmarkClassifyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassifyNumber(50)
	}
}

func BenchmarkClassifyHTTPStatus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassifyHTTPStatus(404)
	}
}
