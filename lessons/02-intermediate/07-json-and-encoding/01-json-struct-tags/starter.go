//go:build !solution
// +build !solution

package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Count   int    `json:"count,omitempty"`
}

func EncodeEvent(event Event) (string, error) {
	data, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecodeEvent(input string) (Event, error) {
	var event Event
	err := json.Unmarshal([]byte(input), &event)
	return event, err
}

func SummaryLabel(event Event) string {
	return fmt.Sprintf("%s:%s:%d", event.Service, event.Status, event.Count)
}

func main() {
	text, _ := EncodeEvent(Event{Service: "billing-api", Status: "ok", Count: 2})
	fmt.Println(text)
}
