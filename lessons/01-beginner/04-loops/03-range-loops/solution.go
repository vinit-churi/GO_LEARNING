//go:build !starter
// +build !starter

package main

import (
	"fmt"
	"strings"
)

// Exercise 1: Sum all numbers in a slice using range
func SumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Exercise 2: Find the index of target string in items slice
func FindIndex(items []string, target string) int {
	for index, item := range items {
		if item == target {
			return index
		}
	}
	return -1
}

// Exercise 3: Count vowels in a string
func CountVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	
	return count
}

// Exercise 4: Double each element in the slice
func DoubleElements(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	
	for _, num := range numbers {
		result = append(result, num*2)
	}
	
	return result
}

// Exercise 5: Reverse a string using runes
func ReverseString(s string) string {
	runes := []rune(s)
	
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	return string(runes)
}

func main() {
	fmt.Println("Exercise 1: Sum of Slice")
	fmt.Println(SumSlice([]int{1, 2, 3, 4, 5}))      // 15
	fmt.Println(SumSlice([]int{10, 20, 30}))         // 60
	fmt.Println(SumSlice([]int{}))                   // 0
	fmt.Println()

	fmt.Println("Exercise 2: Find Index")
	fmt.Println(FindIndex([]string{"apple", "banana", "cherry"}, "banana"))  // 1
	fmt.Println(FindIndex([]string{"apple", "banana", "cherry"}, "grape"))   // -1
	fmt.Println(FindIndex([]string{}, "anything"))                            // -1
	fmt.Println()

	fmt.Println("Exercise 3: Count Vowels")
	fmt.Println(CountVowels("hello"))      // 2
	fmt.Println(CountVowels("AEIOU"))      // 5
	fmt.Println(CountVowels("xyz"))        // 0
	fmt.Println(CountVowels("café"))       // 1
	fmt.Println()

	fmt.Println("Exercise 4: Double Elements")
	fmt.Println(DoubleElements([]int{1, 2, 3}))       // [2 4 6]
	fmt.Println(DoubleElements([]int{5, 10, 15}))     // [10 20 30]
	fmt.Println(DoubleElements([]int{}))              // []
	fmt.Println()

	fmt.Println("Exercise 5: Reverse String")
	fmt.Println(ReverseString("hello"))    // olleh
	fmt.Println(ReverseString("Go!"))      // !oG
	fmt.Println(ReverseString("café"))     // éfac
}
