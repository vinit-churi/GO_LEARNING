//go:build solution
// +build solution

package main

import (
	"fmt"
	"strings"
)

type Order struct {
	Customer string
	Cents    int
}

func ApplyDiscount(order *Order, percent int) {
	order.Cents -= order.Cents * percent / 100
}

func RenameCustomer(order *Order, name string) {
	order.Customer = strings.TrimSpace(name)
}

func SnapshotLabel(order Order) string {
	return fmt.Sprintf("%s:%d", order.Customer, order.Cents)
}

func main() {
	order := Order{Customer: "Mina", Cents: 2500}
	ApplyDiscount(&order, 20)
	RenameCustomer(&order, " Mina Patel ")
	fmt.Println(SnapshotLabel(order))
}
