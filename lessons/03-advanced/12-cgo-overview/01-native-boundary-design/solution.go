//go:build solution
// +build solution

package main

import "fmt"

type Pair struct {
	Left  int
	Right int
}

type NativeAdder interface {
	Add(a, b int) int
}

type FakeDriver struct{}

func (FakeDriver) Add(a, b int) int {
	return a + b
}

func NativeCall(a, b int, driver NativeAdder) int {
	return driver.Add(a, b)
}

func SumPair(pair Pair, driver NativeAdder) int {
	return NativeCall(pair.Left, pair.Right, driver)
}

func ReportPair(pair Pair, driver NativeAdder) string {
	return fmt.Sprintf("sum=%d", SumPair(pair, driver))
}

func main() {
	fmt.Println(ReportPair(Pair{Left: 2, Right: 3}, FakeDriver{}))
}
