package slugutil

import "strings"

func Slug(input string) string {
	trimmed := strings.TrimSpace(strings.ToLower(input))
	fields := strings.Fields(trimmed)
	return strings.Join(fields, "-")
}
