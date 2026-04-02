package main

import (
	"sync"
	"testing"
)

func TestSafeCounterInc(t *testing.T) {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	if got := counter.Value(); got != 10 {
		t.Fatalf("Value() = %d", got)
	}
}

func TestSafeCounterValue(t *testing.T) {
	counter := SafeCounter{}
	counter.Inc()
	if got := counter.Value(); got != 1 {
		t.Fatalf("Value() = %d", got)
	}
}

func TestIncrementAtomic(t *testing.T) {
	var counter int64
	IncrementAtomic(&counter)
	IncrementAtomic(&counter)
	if counter != 2 {
		t.Fatalf("atomic counter = %d", counter)
	}
}
