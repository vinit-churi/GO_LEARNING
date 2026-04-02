//go:build !solution
// +build !solution

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type APIStatus struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func FetchStatus(ctx context.Context, client *http.Client, url string) (APIStatus, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return APIStatus{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return APIStatus{}, err
	}
	defer resp.Body.Close()

	var status APIStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return APIStatus{}, err
	}
	return status, nil
}

func StatusLabel(status APIStatus) string {
	return fmt.Sprintf("%s:%s", status.Service, status.Status)
}

func FetchStatusLabel(ctx context.Context, client *http.Client, url string) (string, error) {
	status, err := FetchStatus(ctx, client, url)
	if err != nil {
		return "", err
	}
	return StatusLabel(status), nil
}

func main() {
	fmt.Println("use go test to exercise the HTTP client")
}
