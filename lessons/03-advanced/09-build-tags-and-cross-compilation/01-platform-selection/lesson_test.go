package main

import (
	"runtime"
	"testing"
)

func TestTaggedPlatform(t *testing.T) {
	if got := taggedPlatform(); got != runtime.GOOS {
		t.Fatalf("taggedPlatform() = %q, want %q", got, runtime.GOOS)
	}
}

func TestBuildArtifactName(t *testing.T) {
	got := BuildArtifactName("study-guide")
	if runtime.GOOS == "windows" {
		if got != "study-guide.exe" {
			t.Fatalf("BuildArtifactName() = %q", got)
		}
		return
	}
	if got != "study-guide" {
		t.Fatalf("BuildArtifactName() = %q", got)
	}
}

func TestTargetTriple(t *testing.T) {
	if got := TargetTriple("linux", "amd64"); got != "linux/amd64" {
		t.Fatalf("TargetTriple() = %q", got)
	}
}
