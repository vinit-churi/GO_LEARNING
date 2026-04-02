package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetGetStatus(t *testing.T) {
	service := NewStatusService()
	service.SetStatus("billing", "ok")
	got, ok := service.GetStatus("billing")
	if !ok || got != "ok" {
		t.Fatalf("GetStatus() = %q, %v", got, ok)
	}
}

func TestGetStatusMissing(t *testing.T) {
	service := NewStatusService()
	if _, ok := service.GetStatus("missing"); ok {
		t.Fatal("GetStatus(missing) = true")
	}
}

func TestHandler(t *testing.T) {
	service := NewStatusService()
	service.SetStatus("billing", "ok")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/status?service=billing", nil)
	service.Handler()(recorder, request)
	if got := recorder.Body.String(); got != "{\"service\":\"billing\",\"status\":\"ok\"}\n" {
		t.Fatalf("Handler() body = %q", got)
	}
}
