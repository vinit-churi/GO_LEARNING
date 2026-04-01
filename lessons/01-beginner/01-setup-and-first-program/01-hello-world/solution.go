//go:build !starter
// +build !starter

package main

import "fmt"

func main() {
	exercise1()
	fmt.Println()
	
	exercise2()
	fmt.Println()
	
	exercise3()
	fmt.Println()
	
	exercise4()
}

// Exercise 1: Hello, World!
func exercise1() {
	fmt.Println("Hello, World!")
}

// Exercise 2: Greet Yourself
func exercise2() {
	fmt.Println("Hello, Alice!")
	fmt.Println("Welcome to Go programming.")
}

// Exercise 3: Multiple Lines
func exercise3() {
	fmt.Println("Hello, Go learner!")
	fmt.Println("You are learning the Go programming language.")
	fmt.Println("Keep practicing and you'll master it!")
}

// Exercise 4: ASCII Art
func exercise4() {
	fmt.Println(" ____  ____  ")
	fmt.Println("|  __||  _ \\ ")
	fmt.Println("| |_  | |_) |")
	fmt.Println("|  _| |  _ < ")
	fmt.Println("|_|   |_| \\_\\")
	fmt.Println()
	fmt.Println("Welcome to Go!")
}
