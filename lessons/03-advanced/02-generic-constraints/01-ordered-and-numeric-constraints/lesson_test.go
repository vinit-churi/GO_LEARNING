package main

import "testing"

func TestSumNumbers(t *testing.T) {
	if got := SumNumbers([]int{1, 2, 3}); got != 6 {
		t.Fatalf("SumNumbers() = %d", got)
	}
}

func TestMaxOrdered(t *testing.T) {
	if got := MaxOrdered([]string{"go", "zebra", "api"}); got != "zebra" {
		t.Fatalf("MaxOrdered() = %q", got)
	}
}

func TestClampOrdered(t *testing.T) {
	if got := ClampOrdered(15, 0, 10); got != 10 {
		t.Fatalf("ClampOrdered(high) = %d", got)
	}
	if got := ClampOrdered(-2, 0, 10); got != 0 {
		t.Fatalf("ClampOrdered(low) = %d", got)
	}
}
