//go:build !starter
// +build !starter

package main

import "fmt"

// Exercise 1: CountToN returns a slice containing numbers from 1 to n (inclusive).
func CountToN(n int) []int {
	if n < 1 {
		return []int{}
	}
	
	result := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

// Exercise 2: MultiplicationTable generates a multiplication table of the given size.
func MultiplicationTable(size int) [][]int {
	table := make([][]int, size)
	
	for i := 0; i < size; i++ {
		table[i] = make([]int, size)
		for j := 0; j < size; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	
	return table
}

// Exercise 3: FindFirstDivisible finds the first number divisible by the divisor.
func FindFirstDivisible(numbers []int, divisor int) int {
	if divisor == 0 {
		return -1
	}
	
	for _, num := range numbers {
		if num%divisor == 0 {
			return num
		}
	}
	
	return -1
}

// Exercise 4: SumOddNumbers sums only the odd numbers in a slice.
func SumOddNumbers(numbers []int) int {
	sum := 0
	
	for _, num := range numbers {
		if num%2 == 0 {
			continue
		}
		sum += num
	}
	
	return sum
}

// Exercise 5: ReadUntilNegative reads numbers from a slice until a negative number is found.
func ReadUntilNegative(numbers []int) []int {
	result := []int{}
	
	for i := 0; ; i++ {
		if i >= len(numbers) {
			break
		}
		if numbers[i] < 0 {
			break
		}
		result = append(result, numbers[i])
	}
	
	return result
}

func main() {
	fmt.Println("Exercise 1: CountToN")
	fmt.Println(CountToN(5))
	fmt.Println(CountToN(0))
	fmt.Println()

	fmt.Println("Exercise 2: MultiplicationTable")
	table := MultiplicationTable(3)
	for _, row := range table {
		fmt.Println(row)
	}
	fmt.Println()

	fmt.Println("Exercise 3: FindFirstDivisible")
	fmt.Println(FindFirstDivisible([]int{3, 5, 10, 15}, 5))
	fmt.Println(FindFirstDivisible([]int{1, 2, 3}, 7))
	fmt.Println()

	fmt.Println("Exercise 4: SumOddNumbers")
	fmt.Println(SumOddNumbers([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumOddNumbers([]int{2, 4, 6}))
	fmt.Println()

	fmt.Println("Exercise 5: ReadUntilNegative")
	fmt.Println(ReadUntilNegative([]int{1, 2, -1, 3}))
	fmt.Println(ReadUntilNegative([]int{5, 10, 15}))
}
