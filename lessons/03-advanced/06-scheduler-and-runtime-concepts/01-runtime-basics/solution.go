//go:build solution
// +build solution

package main

import (
	"fmt"
	"runtime"
)

func CPUCount() int {
	return runtime.NumCPU()
}

func ChooseWorkers(limit int) int {
	cpus := CPUCount()
	if limit <= 0 {
		return cpus
	}
	if limit < cpus {
		return limit
	}
	return cpus
}

func ParallelSummary(limit int) string {
	return fmt.Sprintf("cpus=%d workers=%d", CPUCount(), ChooseWorkers(limit))
}

func main() {
	fmt.Println(ParallelSummary(8))
}
