package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	recorder := httptest.NewRecorder()
	WriteJSON(recorder, APIStatus{Service: "billing-api", Status: "ok"})
	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q", got)
	}
	if got := recorder.Body.String(); got != "{\"service\":\"billing-api\",\"status\":\"ok\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestStatusHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/status", nil)
	StatusHandler("billing-api")(recorder, request)
	if got := recorder.Body.String(); got != "{\"service\":\"billing-api\",\"status\":\"ok\"}\n" {
		t.Fatalf("body = %q", got)
	}
}

func TestRouteLabel(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/status", nil)
	if got := RouteLabel(request); got != "POST /status" {
		t.Fatalf("RouteLabel() = %q", got)
	}
}
