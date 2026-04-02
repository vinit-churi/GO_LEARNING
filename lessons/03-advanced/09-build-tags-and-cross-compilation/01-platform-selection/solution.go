//go:build solution
// +build solution

package main

import "fmt"

func BuildArtifactName(base string) string {
	if taggedPlatform() == "windows" {
		return base + ".exe"
	}
	return base
}

func TargetTriple(goos, goarch string) string {
	return fmt.Sprintf("%s/%s", goos, goarch)
}

func main() {
	fmt.Println(taggedPlatform(), BuildArtifactName("study-guide"))
}
