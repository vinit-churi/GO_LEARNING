package labelutil

import "strings"

func Header(team string) string {
	return strings.ToUpper(strings.TrimSpace(team)) + " DESK"
}
