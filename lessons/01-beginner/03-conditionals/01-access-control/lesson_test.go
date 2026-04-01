package main

import "testing"

func TestAccessMessage(t *testing.T) {
	testCases := []struct {
		role   string
		active bool
		want   string
	}{
		{"admin", true, "admin access granted"},
		{"member", true, "member access granted"},
		{"guest", true, "access denied"},
		{"admin", false, "account inactive"},
	}

	for _, tc := range testCases {
		if got := accessMessage(tc.role, tc.active); got != tc.want {
			t.Fatalf("accessMessage(%q, %v) = %q, want %q", tc.role, tc.active, got, tc.want)
		}
	}
}

func TestShippingCost(t *testing.T) {
	testCases := []struct {
		subtotal float64
		premium  bool
		want     float64
	}{
		{10, false, 4.99},
		{50, false, 0},
		{125, false, 0},
		{10, true, 0},
	}

	for _, tc := range testCases {
		if got := shippingCost(tc.subtotal, tc.premium); got != tc.want {
			t.Fatalf("shippingCost(%.2f, %v) = %.2f, want %.2f", tc.subtotal, tc.premium, got, tc.want)
		}
	}
}

func TestPassOrRetry(t *testing.T) {
	testCases := []struct {
		score int
		want  string
	}{
		{95, "excellent"},
		{60, "pass"},
		{89, "pass"},
		{59, "retry"},
	}

	for _, tc := range testCases {
		if got := passOrRetry(tc.score); got != tc.want {
			t.Fatalf("passOrRetry(%d) = %q, want %q", tc.score, got, tc.want)
		}
	}
}
