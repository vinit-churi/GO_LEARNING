package main

import "testing"

func TestDeskHeader(t *testing.T) {
	if got := DeskHeader(" support "); got != "SUPPORT DESK" {
		t.Fatalf("DeskHeader returned %q", got)
	}
}

func TestBuildMemberCode(t *testing.T) {
	if got := BuildMemberCode("sales", " aria "); got != "SALES-ARIA" {
		t.Fatalf("BuildMemberCode returned %q", got)
	}
}

func TestHasShortName(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Mia", true},
		{" Liam ", true},
		{"Noah", true},
		{"Olivia", false},
	}

	for _, tt := range tests {
		if got := HasShortName(tt.name); got != tt.want {
			t.Fatalf("HasShortName(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}
