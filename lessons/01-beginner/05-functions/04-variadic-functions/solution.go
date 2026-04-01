//go:build !starter
// +build !starter

package main

import "fmt"

// Exercise 1: Sum Function
// Calculates the sum of any number of integers.
func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Exercise 2: Max Function
// Returns the maximum value from any number of integers.
// Returns 0 if no arguments are provided.
func Max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	
	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// Exercise 3: Join With Separator
// Joins multiple strings with a separator between them.
func JoinWithSeparator(separator string, words ...string) string {
	if len(words) == 0 {
		return ""
	}
	
	result := words[0]
	for _, word := range words[1:] {
		result += separator + word
	}
	return result
}

// Exercise 4: Filter Evens
// Returns a slice containing only the even numbers from the input.
func FilterEvens(numbers ...int) []int {
	result := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

// Exercise 5: Concatenate All
// Concatenates multiple slices into a single slice.
func ConcatenateAll(slices ...[]int) []int {
	// Calculate total length for efficient allocation
	totalLen := 0
	for _, slice := range slices {
		totalLen += len(slice)
	}
	
	// Pre-allocate result with exact capacity needed
	result := make([]int, 0, totalLen)
	
	// Append all elements from all slices
	for _, slice := range slices {
		result = append(result, slice...)
	}
	
	return result
}

func main() {
	fmt.Println("=== Exercise 1: Sum ===")
	fmt.Println("Sum(1, 2, 3):", Sum(1, 2, 3))
	fmt.Println("Sum(10, 20):", Sum(10, 20))
	fmt.Println("Sum():", Sum())
	fmt.Println("Sum(100):", Sum(100))
	
	fmt.Println("\n=== Exercise 2: Max ===")
	fmt.Println("Max(1, 5, 3, 2):", Max(1, 5, 3, 2))
	fmt.Println("Max(-10, -5, -20):", Max(-10, -5, -20))
	fmt.Println("Max(42):", Max(42))
	fmt.Println("Max():", Max())
	
	fmt.Println("\n=== Exercise 3: Join With Separator ===")
	fmt.Println("Join with ', ':", JoinWithSeparator(", ", "apple", "banana", "cherry"))
	fmt.Println("Join with ' - ':", JoinWithSeparator(" - ", "Go", "is", "awesome"))
	fmt.Println("Join with '|':", JoinWithSeparator("|", "one"))
	fmt.Println("Join with empty words:", JoinWithSeparator(", "))
	
	fmt.Println("\n=== Exercise 4: Filter Evens ===")
	fmt.Println("Filter [1,2,3,4,5,6]:", FilterEvens(1, 2, 3, 4, 5, 6))
	fmt.Println("Filter [1,3,5,7]:", FilterEvens(1, 3, 5, 7))
	fmt.Println("Filter [2,4,6]:", FilterEvens(2, 4, 6))
	fmt.Println("Filter []:", FilterEvens())
	
	fmt.Println("\n=== Exercise 5: Concatenate All ===")
	result := ConcatenateAll([]int{1, 2}, []int{3, 4}, []int{5})
	fmt.Println("Concatenate [[1,2], [3,4], [5]]:", result)
	
	result2 := ConcatenateAll([]int{10, 20})
	fmt.Println("Concatenate [[10,20]]:", result2)
	
	result3 := ConcatenateAll()
	fmt.Println("Concatenate []:", result3)
	
	// Bonus: Demonstrating the spread operator
	nums := []int{5, 10, 15, 20}
	fmt.Println("\n=== Bonus: Spread Operator ===")
	fmt.Println("nums slice:", nums)
	fmt.Println("Sum(nums...):", Sum(nums...))
	fmt.Println("Max(nums...):", Max(nums...))
	
	// Combining direct args and slice
	fmt.Println("Sum(1, 2, nums...):", Sum(append([]int{1, 2}, nums...)...))
}
