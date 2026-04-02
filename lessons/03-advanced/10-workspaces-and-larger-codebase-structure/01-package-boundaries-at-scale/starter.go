//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"strings"

	"github.com/vinit-churi/learning-go/lessons/03-advanced/10-workspaces-and-larger-codebase-structure/01-package-boundaries-at-scale/workspaceplan"
)

func RootModuleLabel(owner, repo string) string {
	return workspaceplan.ModulePath(owner, repo)
}

func IsInternalPath(path string) bool {
	return strings.Contains(path, "/internal/") || strings.HasPrefix(path, "internal/")
}

func PackageSummary(path string) string {
	return fmt.Sprintf("pkg=%s scope=%s", path, workspaceplan.Scope(path))
}

func main() {
	fmt.Println(RootModuleLabel("acme", "study-guide"))
}
