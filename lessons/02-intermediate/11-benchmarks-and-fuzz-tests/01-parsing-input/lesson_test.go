package main

import "testing"

func TestParsePair(t *testing.T) {
	name, count, err := ParsePair("jobs=3")
	if err != nil {
		t.Fatalf("ParsePair() error = %v", err)
	}
	if name != "jobs" || count != 3 {
		t.Fatalf("ParsePair() = %q, %d", name, count)
	}
}

func TestMustParseCount(t *testing.T) {
	if got := MustParseCount("bad input"); got != 0 {
		t.Fatalf("MustParseCount() = %d", got)
	}
}

func TestFormatPair(t *testing.T) {
	if got := FormatPair("jobs", 3); got != "jobs=3" {
		t.Fatalf("FormatPair() = %q", got)
	}
}

func BenchmarkParsePair(b *testing.B) {
	for b.Loop() {
		_, _, err := ParsePair("jobs=3")
		if err != nil {
			b.Fatalf("ParsePair() error = %v", err)
		}
	}
}

func FuzzParsePair(f *testing.F) {
	f.Add("jobs=3")
	f.Add("workers=12")
	f.Add("broken")

	f.Fuzz(func(t *testing.T, input string) {
		_, _, _ = ParsePair(input)
	})
}
