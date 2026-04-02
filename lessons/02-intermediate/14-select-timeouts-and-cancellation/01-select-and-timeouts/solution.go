//go:build solution
// +build solution

package main

import (
	"fmt"
	"time"
)

func ReceiveWithTimeout(ch <-chan string, timeout time.Duration) string {
	select {
	case value := <-ch:
		return value
	case <-time.After(timeout):
		return "timeout"
	}
}

func ImmediateStatus(ch <-chan string) string {
	select {
	case value := <-ch:
		return value
	default:
		return "idle"
	}
}

func FirstResult(left, right <-chan string) string {
	select {
	case value := <-left:
		return value
	case value := <-right:
		return value
	}
}

func main() {
	ch := make(chan string, 1)
	ch <- "ready"
	fmt.Println(ReceiveWithTimeout(ch, 10*time.Millisecond))
}
