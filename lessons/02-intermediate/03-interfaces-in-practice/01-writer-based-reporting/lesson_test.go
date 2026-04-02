package main

import (
	"bytes"
	"testing"
)

func TestWriteLine(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteLine(&buf, "ready"); err != nil {
		t.Fatalf("WriteLine() error = %v", err)
	}
	if got := buf.String(); got != "ready\n" {
		t.Fatalf("WriteLine() wrote %q", got)
	}
}

func TestEmitReport(t *testing.T) {
	var buf bytes.Buffer
	if err := EmitReport(&buf, "billing-api", 3); err != nil {
		t.Fatalf("EmitReport() error = %v", err)
	}
	if got := buf.String(); got != "service=billing-api count=3\n" {
		t.Fatalf("EmitReport() wrote %q", got)
	}
}

func TestBufferReport(t *testing.T) {
	if got := BufferReport("billing-api", 3); got != "service=billing-api count=3\n" {
		t.Fatalf("BufferReport() = %q", got)
	}
}
