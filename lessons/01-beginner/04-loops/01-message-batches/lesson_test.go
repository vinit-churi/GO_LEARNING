package main

import (
	"reflect"
	"testing"
)

func TestTotalMessages(t *testing.T) {
	testCases := []struct {
		days int
		want int
	}{
		{0, 0},
		{1, 1},
		{4, 10},
		{7, 28},
	}

	for _, tc := range testCases {
		if got := totalMessages(tc.days); got != tc.want {
			t.Fatalf("totalMessages(%d) = %d, want %d", tc.days, got, tc.want)
		}
	}
}

func TestBuildCountdown(t *testing.T) {
	testCases := []struct {
		start int
		want  []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{4, []int{4, 3, 2, 1}},
	}

	for _, tc := range testCases {
		if got := buildCountdown(tc.start); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("buildCountdown(%d) = %v, want %v", tc.start, got, tc.want)
		}
	}
}

func TestSumEvenNumbers(t *testing.T) {
	testCases := []struct {
		limit int
		want  int
	}{
		{0, 0},
		{1, 0},
		{6, 12},
		{10, 30},
	}

	for _, tc := range testCases {
		if got := sumEvenNumbers(tc.limit); got != tc.want {
			t.Fatalf("sumEvenNumbers(%d) = %d, want %d", tc.limit, got, tc.want)
		}
	}
}
