//go:build !solution
// +build !solution

package main

import (
	"context"
	"fmt"
)

func RunOnce(ctx context.Context, work func() string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return work(), nil
	}
}

func LabelResult(value string) string {
	return "result:" + value
}

func RunLabeled(ctx context.Context, work func() string) (string, error) {
	value, err := RunOnce(ctx, work)
	if err != nil {
		return "", err
	}
	return LabelResult(value), nil
}

func main() {
	value, _ := RunLabeled(context.Background(), func() string { return "ready" })
	fmt.Println(value)
}
