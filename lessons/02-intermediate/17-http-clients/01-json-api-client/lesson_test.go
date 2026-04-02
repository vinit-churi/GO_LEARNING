package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"service":"billing-api","status":"ok"}`)
	}))
	defer server.Close()

	got, err := FetchStatus(context.Background(), server.Client(), server.URL)
	if err != nil {
		t.Fatalf("FetchStatus() error = %v", err)
	}
	if got.Service != "billing-api" || got.Status != "ok" {
		t.Fatalf("FetchStatus() = %+v", got)
	}
}

func TestStatusLabel(t *testing.T) {
	if got := StatusLabel(APIStatus{Service: "billing-api", Status: "ok"}); got != "billing-api:ok" {
		t.Fatalf("StatusLabel() = %q", got)
	}
}

func TestFetchStatusLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"service":"billing-api","status":"ok"}`)
	}))
	defer server.Close()

	got, err := FetchStatusLabel(context.Background(), server.Client(), server.URL)
	if err != nil {
		t.Fatalf("FetchStatusLabel() error = %v", err)
	}
	if got != "billing-api:ok" {
		t.Fatalf("FetchStatusLabel() = %q", got)
	}
}
