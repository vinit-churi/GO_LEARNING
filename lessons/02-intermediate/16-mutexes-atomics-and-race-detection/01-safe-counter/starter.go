//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func IncrementAtomic(counter *int64) {
	atomic.AddInt64(counter, 1)
}

func main() {
	var atomicCount int64
	counter := SafeCounter{}
	counter.Inc()
	IncrementAtomic(&atomicCount)
	fmt.Println(counter.Value(), atomicCount)
}
