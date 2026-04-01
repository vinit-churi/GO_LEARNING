//go:build !solution
// +build !solution

package main

import "fmt"

type Counter struct {
	Name  string
	Value int
}

func (c Counter) Summary() string {
	return fmt.Sprintf("%s: %d", c.Name, c.Value)
}

func (c *Counter) Add(amount int) {
	c.Value += amount
}

func (c *Counter) Reset() {
	c.Value = 0
}

func main() {
	counter := Counter{Name: "Orders", Value: 3}
	counter.Add(2)
	fmt.Println(counter.Summary())
	counter.Reset()
	fmt.Println(counter.Summary())
}
