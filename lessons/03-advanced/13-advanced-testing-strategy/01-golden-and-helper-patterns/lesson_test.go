package main

import "testing"

func TestNormalizeReport(t *testing.T) {
	got := NormalizeReport(" name: alpha  \n count:   2 \n")
	want := "name: alpha\ncount: 2"
	if got != want {
		t.Fatalf("NormalizeReport() = %q, want %q", got, want)
	}
}

func TestBuildGolden(t *testing.T) {
	if got := BuildGolden("alpha", 2); got != "name: alpha\ncount: 2" {
		t.Fatalf("BuildGolden() = %q", got)
	}
}

func TestReportsMatch(t *testing.T) {
	left := " name: alpha\ncount: 2\n"
	right := "name:  alpha\n count: 2"
	if !ReportsMatch(left, right) {
		t.Fatal("ReportsMatch() = false")
	}
}
