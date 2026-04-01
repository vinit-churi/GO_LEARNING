//go:build solution

package main

import "fmt"

const maxDailyPracticeMinutes = 90

type PracticeSession struct {
	Name      string
	Minutes   int
	Completed bool
}

func newPracticeSession(name string) PracticeSession {
	return PracticeSession{Name: name}
}

func remainingPracticeMinutes(minutesSpent int) int {
	remaining := maxDailyPracticeMinutes - minutesSpent
	if remaining < 0 {
		return 0
	}

	return remaining
}

func hasStarted(session PracticeSession) bool {
	return session.Minutes > 0 || session.Completed
}

func main() {
	session := newPracticeSession("variables")
	fmt.Println(session)
	fmt.Println(remainingPracticeMinutes(35))
	fmt.Println(hasStarted(session))
}
