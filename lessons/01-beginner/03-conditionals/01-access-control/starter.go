package main

import "fmt"

func accessMessage(role string, active bool) string {
	if !active {
		return "account inactive"
	}
	if role == "admin" {
		return "admin access granted"
	}
	if role == "member" {
		return "member access granted"
	}

	return "access denied"
}

func shippingCost(subtotal float64, premium bool) float64 {
	if premium || subtotal >= 50 {
		return 0
	}

	return 4.99
}

func passOrRetry(score int) string {
	if score >= 90 {
		return "excellent"
	}
	if score >= 60 {
		return "pass"
	}

	return "retry"
}

func main() {
	fmt.Println(accessMessage("admin", true))
	fmt.Println(shippingCost(42, false))
	fmt.Println(passOrRetry(88))
}
