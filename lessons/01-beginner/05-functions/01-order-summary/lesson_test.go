package main

import "testing"

func TestSubtotal(t *testing.T) {
	testCases := []struct {
		prices []float64
		want   float64
	}{
		{[]float64{}, 0},
		{[]float64{10}, 10},
		{[]float64{10, 15.5, 4.5}, 30},
	}

	for _, tc := range testCases {
		if got := subtotal(tc.prices); got != tc.want {
			t.Fatalf("subtotal(%v) = %.2f, want %.2f", tc.prices, got, tc.want)
		}
	}
}

func TestApplyDiscount(t *testing.T) {
	testCases := []struct {
		amount   float64
		isMember bool
		want     float64
	}{
		{30, false, 30},
		{30, true, 27},
	}

	for _, tc := range testCases {
		if got := applyDiscount(tc.amount, tc.isMember); got != tc.want {
			t.Fatalf("applyDiscount(%.2f, %v) = %.2f, want %.2f", tc.amount, tc.isMember, got, tc.want)
		}
	}
}

func TestOrderSummary(t *testing.T) {
	got := orderSummary("Asha", []float64{10, 15.5, 4.5}, true)
	want := "Asha: subtotal=30.00, total=27.00"
	if got != want {
		t.Fatalf("orderSummary() = %q, want %q", got, want)
	}
}
