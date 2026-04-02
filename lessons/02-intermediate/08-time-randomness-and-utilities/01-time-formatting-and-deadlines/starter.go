//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"time"
)

func RFC3339Label(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func ExtendDeadline(start time.Time, minutes int) time.Time {
	return start.Add(time.Duration(minutes) * time.Minute)
}

func MinutesRemaining(now, deadline time.Time) int {
	return int(deadline.Sub(now) / time.Minute)
}

func main() {
	start := time.Date(2026, time.April, 2, 10, 0, 0, 0, time.UTC)
	fmt.Println(RFC3339Label(ExtendDeadline(start, 45)))
}
