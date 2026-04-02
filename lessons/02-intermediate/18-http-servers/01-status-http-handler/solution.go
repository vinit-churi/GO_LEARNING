//go:build solution
// +build solution

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIStatus struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func WriteJSON(w http.ResponseWriter, status APIStatus) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(status)
}

func StatusHandler(service string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, APIStatus{Service: service, Status: "ok"})
	}
}

func RouteLabel(r *http.Request) string {
	return fmt.Sprintf("%s %s", r.Method, r.URL.Path)
}

func main() {
	fmt.Println("use go test to exercise the HTTP handler")
}
