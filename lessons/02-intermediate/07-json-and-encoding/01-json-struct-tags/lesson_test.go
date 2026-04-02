package main

import "testing"

func TestEncodeEvent(t *testing.T) {
	got, err := EncodeEvent(Event{Service: "billing-api", Status: "ok", Count: 2})
	if err != nil {
		t.Fatalf("EncodeEvent() error = %v", err)
	}
	if got != `{"service":"billing-api","status":"ok","count":2}` {
		t.Fatalf("EncodeEvent() = %q", got)
	}
}

func TestDecodeEvent(t *testing.T) {
	got, err := DecodeEvent(`{"service":"billing-api","status":"ok","count":2}`)
	if err != nil {
		t.Fatalf("DecodeEvent() error = %v", err)
	}
	if got.Service != "billing-api" || got.Status != "ok" || got.Count != 2 {
		t.Fatalf("DecodeEvent() = %+v", got)
	}
}

func TestSummaryLabel(t *testing.T) {
	got := SummaryLabel(Event{Service: "billing-api", Status: "ok", Count: 2})
	if got != "billing-api:ok:2" {
		t.Fatalf("SummaryLabel() = %q", got)
	}
}
