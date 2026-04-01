//go:build starter
// +build starter

package main

import "fmt"

// Exercise 1: Sum all numbers in a slice using range
func SumSlice(numbers []int) int {
	// TODO: Use range to iterate over numbers and sum them
	// Hint: You can ignore the index with _
	return 0
}

// Exercise 2: Find the index of target string in items slice
func FindIndex(items []string, target string) int {
	// TODO: Use range to iterate and find the target
	// Hint: Return the index when you find a match
	// Return -1 if not found
	return -1
}

// Exercise 3: Count vowels in a string
func CountVowels(s string) int {
	// TODO: Use range to iterate over the string as runes
	// Hint: Create a helper function or check if rune is in "aeiouAEIOU"
	// Remember: ranging over a string gives you runes, not bytes
	return 0
}

// Exercise 4: Double each element in the slice
func DoubleElements(numbers []int) []int {
	// TODO: Create a new slice and append doubled values
	// Hint: Use _ to ignore the index since you only need the value
	return nil
}

// Exercise 5: Reverse a string using runes
func ReverseString(s string) string {
	// TODO: Convert string to runes, reverse them, convert back to string
	// Hint: Range over the string to get runes, prepend each to build reversed string
	// Or: collect runes in a slice, then reverse the slice
	return ""
}

func main() {
	fmt.Println("Exercise 1: Sum of Slice")
	fmt.Println(SumSlice([]int{1, 2, 3, 4, 5}))
	fmt.Println()

	fmt.Println("Exercise 2: Find Index")
	fmt.Println(FindIndex([]string{"apple", "banana", "cherry"}, "banana"))
	fmt.Println(FindIndex([]string{"apple", "banana", "cherry"}, "grape"))
	fmt.Println()

	fmt.Println("Exercise 3: Count Vowels")
	fmt.Println(CountVowels("hello"))
	fmt.Println(CountVowels("AEIOU"))
	fmt.Println()

	fmt.Println("Exercise 4: Double Elements")
	fmt.Println(DoubleElements([]int{1, 2, 3}))
	fmt.Println()

	fmt.Println("Exercise 5: Reverse String")
	fmt.Println(ReverseString("hello"))
	fmt.Println(ReverseString("café"))
}
