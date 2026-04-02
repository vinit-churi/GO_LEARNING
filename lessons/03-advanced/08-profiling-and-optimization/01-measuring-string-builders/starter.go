//go:build !solution
// +build !solution

package main

import (
	"strings"
)

func JoinWithBuilder(parts []string) string {
	var b strings.Builder
	for i, part := range parts {
		if i > 0 {
			b.WriteString("/")
		}
		b.WriteString(part)
	}
	return b.String()
}

func JoinWithPlus(parts []string) string {
	out := ""
	for i, part := range parts {
		if i > 0 {
			out += "/"
		}
		out += part
	}
	return out
}

func ReportLine(parts []string) string {
	return "path=" + JoinWithBuilder(parts)
}

func main() {}
