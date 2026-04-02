//go:build solution
// +build solution

package main

import "fmt"

func ClassifyScore(score int) string {
	switch {
	case score >= 90:
		return "high"
	case score >= 60:
		return "medium"
	default:
		return "low"
	}
}

func IsTerminal(status string) bool {
	return status == "done" || status == "failed"
}

func BuildStatusLine(name string, score int) string {
	return fmt.Sprintf("%s:%s", name, ClassifyScore(score))
}

func main() {
	fmt.Println(BuildStatusLine("job-42", 91))
}
