package workspaceplan

import "fmt"

func ModulePath(owner, repo string) string {
	return fmt.Sprintf("github.com/%s/%s", owner, repo)
}

func Scope(path string) string {
	if len(path) >= len("internal") && (path == "internal" || path[:len("internal/")] == "internal/") {
		return "internal"
	}
	if containsInternal(path) {
		return "internal"
	}
	return "public"
}

func containsInternal(path string) bool {
	for i := 0; i+len("/internal/") <= len(path); i++ {
		if path[i:i+len("/internal/")] == "/internal/" {
			return true
		}
	}
	return false
}
