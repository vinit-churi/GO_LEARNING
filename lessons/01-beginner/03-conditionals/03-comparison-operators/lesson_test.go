package main

import (
	"reflect"
	"testing"
)

func TestCompareIntegers(t *testing.T) {
	got := CompareIntegers(5, 3)
	want := map[string]bool{
		"equal":          false,
		"notEqual":       true,
		"lessThan":       false,
		"greaterThan":    true,
		"lessOrEqual":    false,
		"greaterOrEqual": true,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CompareIntegers(5, 3) = %v, want %v", got, want)
	}
}

func TestIsValidAge(t *testing.T) {
	testCases := []struct {
		age  int
		want bool
	}{
		{-1, false},
		{0, true},
		{25, true},
		{150, true},
		{151, false},
	}

	for _, tc := range testCases {
		if got := IsValidAge(tc.age); got != tc.want {
			t.Fatalf("IsValidAge(%d) = %v, want %v", tc.age, got, tc.want)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	if !IsWeekend("Saturday") || !IsWeekend("Sunday") || IsWeekend("Monday") {
		t.Fatalf("IsWeekend logic returned unexpected results")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty("") {
		t.Fatalf("IsNotEmpty(\"\") = true, want false")
	}
	if !IsNotEmpty("go") {
		t.Fatalf("IsNotEmpty(\"go\") = false, want true")
	}
}

func TestCanVote(t *testing.T) {
	testCases := []struct {
		age       int
		isCitizen bool
		want      bool
	}{
		{25, true, true},
		{17, true, false},
		{25, false, false},
	}

	for _, tc := range testCases {
		if got := CanVote(tc.age, tc.isCitizen); got != tc.want {
			t.Fatalf("CanVote(%d, %v) = %v, want %v", tc.age, tc.isCitizen, got, tc.want)
		}
	}
}

func TestNeedsUmbrella(t *testing.T) {
	testCases := []struct {
		rain bool
		snow bool
		want bool
	}{
		{true, false, true},
		{false, true, true},
		{false, false, false},
	}

	for _, tc := range testCases {
		if got := NeedsUmbrella(tc.rain, tc.snow); got != tc.want {
			t.Fatalf("NeedsUmbrella(%v, %v) = %v, want %v", tc.rain, tc.snow, got, tc.want)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	testCases := []struct {
		password        string
		confirmPassword string
		length          int
		want            bool
	}{
		{"secret123", "secret123", 9, true},
		{"secret123", "secret999", 9, false},
		{"secret", "secret", 6, false},
		{"secret123", "secret123", 129, false},
	}

	for _, tc := range testCases {
		if got := IsValidPassword(tc.password, tc.confirmPassword, tc.length); got != tc.want {
			t.Fatalf("IsValidPassword(%q, %q, %d) = %v, want %v", tc.password, tc.confirmPassword, tc.length, got, tc.want)
		}
	}
}

func TestCheckAccessLevel(t *testing.T) {
	testCases := []struct {
		age       int
		isMember  bool
		isPremium bool
		want      string
	}{
		{25, true, true, "full access"},
		{25, true, false, "full access"},
		{16, true, false, "limited access"},
		{25, false, false, "no access"},
	}

	for _, tc := range testCases {
		if got := CheckAccessLevel(tc.age, tc.isMember, tc.isPremium); got != tc.want {
			t.Fatalf("CheckAccessLevel(%d, %v, %v) = %q, want %q", tc.age, tc.isMember, tc.isPremium, got, tc.want)
		}
	}
}

func TestIsOutsideRange(t *testing.T) {
	testCases := []struct {
		value int
		min   int
		max   int
		want  bool
	}{
		{5, 1, 10, false},
		{15, 1, 10, true},
		{1, 1, 10, false},
		{0, 1, 10, true},
	}

	for _, tc := range testCases {
		if got := IsOutsideRange(tc.value, tc.min, tc.max); got != tc.want {
			t.Fatalf("IsOutsideRange(%d, %d, %d) = %v, want %v", tc.value, tc.min, tc.max, got, tc.want)
		}
	}
}
