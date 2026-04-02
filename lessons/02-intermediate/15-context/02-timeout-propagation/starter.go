//go:build !solution
// +build !solution

package main

import (
	"context"
	"fmt"
	"time"
)

func WithRequestTimeout(parent context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, d)
}

func WaitForWork(ctx context.Context, workDone <-chan struct{}) error {
	select {
	case <-workDone:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func RunWithTimeout(parent context.Context, d time.Duration, workDone <-chan struct{}) error {
	ctx, cancel := WithRequestTimeout(parent, d)
	defer cancel()

	return WaitForWork(ctx, workDone)
}

func main() {
	workDone := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Millisecond)
		close(workDone)
	}()

	err := RunWithTimeout(context.Background(), 50*time.Millisecond, workDone)
	fmt.Println(err)
}
