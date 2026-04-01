//go:build !solution
// +build !solution

package main

import "fmt"

func SumFirstThree(numbers [3]int) int {
	return numbers[0] + numbers[1] + numbers[2]
}

func AddTask(tasks []string, task string) []string {
	return append(tasks, task)
}

func WeekendWindow(days []string) []string {
	return days[5:7]
}

func main() {
	fmt.Println(SumFirstThree([3]int{3, 4, 5}))
	fmt.Println(AddTask([]string{"plan", "code"}, "test"))
	fmt.Println(WeekendWindow([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}))
}
