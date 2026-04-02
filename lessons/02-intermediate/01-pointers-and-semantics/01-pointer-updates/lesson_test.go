package main

import "testing"

func TestApplyDiscount(t *testing.T) {
	order := Order{Customer: "Mina", Cents: 2500}
	ApplyDiscount(&order, 20)
	if order.Cents != 2000 {
		t.Fatalf("ApplyDiscount() left cents at %d", order.Cents)
	}
}

func TestRenameCustomer(t *testing.T) {
	order := Order{Customer: "Mina", Cents: 2500}
	RenameCustomer(&order, " Mina Patel ")
	if order.Customer != "Mina Patel" {
		t.Fatalf("RenameCustomer() left customer as %q", order.Customer)
	}
}

func TestSnapshotLabelUsesCopy(t *testing.T) {
	order := Order{Customer: "Mina", Cents: 2500}
	got := SnapshotLabel(order)
	if got != "Mina:2500" {
		t.Fatalf("SnapshotLabel() = %q", got)
	}
	if order.Customer != "Mina" || order.Cents != 2500 {
		t.Fatalf("SnapshotLabel() should not mutate the original order: %+v", order)
	}
}
