//go:build solution
// +build solution

package main

import (
	"fmt"
	"strings"

	"github.com/vinit-churi/learning-go/lessons/02-intermediate/09-modules-and-package-design/01-package-layout-basics/slugutil"
)

func BuildModulePath(owner, repo string) string {
	return fmt.Sprintf("github.com/%s/%s", slugutil.Slug(owner), slugutil.Slug(repo))
}

func IsPublicPackage(name string) bool {
	cleaned := strings.TrimSpace(name)
	return cleaned != "" && !strings.HasPrefix(cleaned, "internal-")
}

func main() {
	fmt.Println(BuildModulePath("Acme Org", "Billing API"))
	fmt.Println(IsPublicPackage("payments"))
}
