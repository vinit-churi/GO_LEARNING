//go:build !solution
// +build !solution

package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

type Customer struct {
	Name string
}

type Branch struct {
	City string
}

type PickupOrder struct {
	Customer Customer
	Branch   Branch
}

func NewBook(title string, pages int) Book {
	return Book{Title: title, Pages: pages}
}

func TotalPages(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.Pages
	}
	return total
}

func PickupLabel(order PickupOrder) string {
	return fmt.Sprintf("%s - %s", order.Customer.Name, order.Branch.City)
}

func main() {
	books := []Book{NewBook("Go Basics", 120), NewBook("Maps", 80)}
	fmt.Println(TotalPages(books))
	fmt.Println(PickupLabel(PickupOrder{
		Customer: Customer{Name: "Ava"},
		Branch:   Branch{City: "Pune"},
	}))
}
