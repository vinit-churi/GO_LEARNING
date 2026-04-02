package main

import "testing"

func TestRootModuleLabel(t *testing.T) {
	if got := RootModuleLabel("acme", "study-guide"); got != "github.com/acme/study-guide" {
		t.Fatalf("RootModuleLabel() = %q", got)
	}
}

func TestIsInternalPath(t *testing.T) {
	if !IsInternalPath("pkg/internal/cache") {
		t.Fatal("IsInternalPath() = false")
	}
	if IsInternalPath("pkg/public/cache") {
		t.Fatal("IsInternalPath(public) = true")
	}
}

func TestPackageSummary(t *testing.T) {
	if got := PackageSummary("pkg/internal/cache"); got != "pkg=pkg/internal/cache scope=internal" {
		t.Fatalf("PackageSummary() = %q", got)
	}
}
