//go:build solution
// +build solution

package main

import (
	"fmt"
	"sync"
)

func PublishOnce(init func() string) string {
	var once sync.Once
	var value string
	once.Do(func() {
		value = init()
	})
	once.Do(func() {
		value = "should-not-run"
	})
	return value
}

func HandOff(value string) string {
	ch := make(chan string, 1)
	go func() {
		ch <- value
	}()
	return <-ch
}

func SnapshotLabel(init func() string) string {
	return "snapshot:" + PublishOnce(init)
}

func main() {
	fmt.Println(SnapshotLabel(func() string { return HandOff("ready") }))
}
