package main

import "testing"

func TestNewBook(t *testing.T) {
	book := NewBook("Go Basics", 120)
	if book.Title != "Go Basics" || book.Pages != 120 {
		t.Fatalf("NewBook returned %+v", book)
	}
}

func TestTotalPages(t *testing.T) {
	books := []Book{{Title: "A", Pages: 10}, {Title: "B", Pages: 20}}
	if got := TotalPages(books); got != 30 {
		t.Fatalf("TotalPages returned %d", got)
	}
}

func TestPickupLabel(t *testing.T) {
	order := PickupOrder{
		Customer: Customer{Name: "Ava"},
		Branch:   Branch{City: "Pune"},
	}
	if got := PickupLabel(order); got != "Ava - Pune" {
		t.Fatalf("PickupLabel returned %q", got)
	}
}
