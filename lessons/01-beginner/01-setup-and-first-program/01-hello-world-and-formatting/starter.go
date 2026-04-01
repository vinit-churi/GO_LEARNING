package main

import "fmt"

func greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func learnerProfile(name string, lessonsCompleted int) string {
	return fmt.Sprintf("%s has completed %d lessons.", name, lessonsCompleted)
}

func isReadyForNextLesson(completed int) bool {
	return completed >= 3
}

func main() {
	fmt.Println(greeting("Gopher"))
	fmt.Println(learnerProfile("Ava", 2))
	fmt.Println(isReadyForNextLesson(3))
}
