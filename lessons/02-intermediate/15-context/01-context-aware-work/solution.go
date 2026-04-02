//go:build solution
// +build solution

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type traceKey struct{}

func TraceID(ctx context.Context) string {
	value, _ := ctx.Value(traceKey{}).(string)
	return value
}

func WorkOrCancel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Millisecond):
		return nil
	}
}

func LabelRequest(ctx context.Context, name string) string {
	return fmt.Sprintf("%s:%s", TraceID(ctx), name)
}

func main() {
	ctx := context.WithValue(context.Background(), traceKey{}, "trace-42")
	fmt.Println(LabelRequest(ctx, "billing"), errors.Is(WorkOrCancel(ctx), context.Canceled))
}
