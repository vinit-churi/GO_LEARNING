//go:build solution

package main

import "fmt"

func subtotal(prices []float64) float64 {
	total := 0.0
	for _, price := range prices {
		total += price
	}

	return total
}

func applyDiscount(amount float64, isMember bool) float64 {
	if !isMember {
		return amount
	}

	return amount * 0.90
}

func orderSummary(name string, prices []float64, isMember bool) string {
	rawSubtotal := subtotal(prices)
	total := applyDiscount(rawSubtotal, isMember)

	return fmt.Sprintf("%s: subtotal=%.2f, total=%.2f", name, rawSubtotal, total)
}

func main() {
	fmt.Println(orderSummary("Asha", []float64{10, 15.5, 4.5}, true))
}
