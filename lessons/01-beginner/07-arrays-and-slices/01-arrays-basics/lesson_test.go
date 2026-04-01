package main

import (
	"reflect"
	"testing"
)

func TestSumFirstThree(t *testing.T) {
	if got := SumFirstThree([3]int{2, 3, 4}); got != 9 {
		t.Fatalf("SumFirstThree returned %d", got)
	}
}

func TestAddTask(t *testing.T) {
	got := AddTask([]string{"plan", "code"}, "test")
	want := []string{"plan", "code", "test"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("AddTask returned %v", got)
	}
}

func TestWeekendWindow(t *testing.T) {
	got := WeekendWindow([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"})
	want := []string{"Sat", "Sun"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("WeekendWindow returned %v", got)
	}
}
