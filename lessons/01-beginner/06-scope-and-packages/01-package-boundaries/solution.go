//go:build solution
// +build solution

package main

import (
	"fmt"
	"strings"

	"github.com/vinit-churi/learning-go/lessons/01-beginner/06-scope-and-packages/01-package-boundaries/labelutil"
)

func DeskHeader(team string) string {
	return labelutil.Header(team)
}

func BuildMemberCode(team, name string) string {
	cleanTeam := strings.ToUpper(strings.TrimSpace(team))
	cleanName := strings.ToUpper(strings.TrimSpace(name))
	return fmt.Sprintf("%s-%s", cleanTeam, cleanName)
}

func HasShortName(name string) bool {
	if trimmed := strings.TrimSpace(name); trimmed != "" {
		return len(trimmed) <= 4
	}
	return true
}

func main() {
	fmt.Println(DeskHeader(" support "))
	fmt.Println(BuildMemberCode("sales", " aria "))
	fmt.Println(HasShortName("Mia"))
}
