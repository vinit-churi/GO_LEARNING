package main

import (
	"strings"
	"testing"
)

func TestGetModulePath(t *testing.T) {
	path := GetModulePath()
	
	// Should not be empty
	if path == "" {
		t.Error("GetModulePath() should not return empty string")
	}
	
	// Should contain at least one slash (domain/path structure)
	if !strings.Contains(path, "/") {
		t.Errorf("Module path should contain '/', got: %s", path)
	}
	
	// Should not contain spaces
	if strings.Contains(path, " ") {
		t.Errorf("Module path should not contain spaces, got: %s", path)
	}
	
	// Should be lowercase (Go convention)
	if strings.ToLower(path) != path {
		t.Errorf("Module path should be lowercase, got: %s", path)
	}
}

func TestGetPackageInfo(t *testing.T) {
	info := GetPackageInfo()
	
	// Should not be empty
	if info == "" {
		t.Error("GetPackageInfo() should not return empty string")
	}
	
	// Should mention package
	if !strings.Contains(strings.ToLower(info), "package") {
		t.Error("Package info should mention 'package'")
	}
	
	// Should mention main package
	if !strings.Contains(strings.ToLower(info), "main") {
		t.Error("Package info should mention 'main'")
	}
	
	// Should have multiple lines or clear formatting
	if !strings.Contains(info, "\n") && !strings.Contains(info, ":") {
		t.Error("Package info should be formatted (contain newline or colon)")
	}
}

func TestGetModuleVersion(t *testing.T) {
	version := GetModuleVersion()
	
	// Should not be empty
	if version == "" {
		t.Error("GetModuleVersion() should not return empty string")
	}
	
	// Should start with 'v'
	if !strings.HasPrefix(version, "v") {
		t.Errorf("Version should start with 'v', got: %s", version)
	}
	
	// Should contain at least two dots (v1.0.0 format)
	if strings.Count(version, ".") < 2 {
		t.Errorf("Version should follow semantic versioning (v1.0.0), got: %s", version)
	}
	
	// Check that it's a valid semver format (basic check)
	parts := strings.Split(strings.TrimPrefix(version, "v"), ".")
	if len(parts) != 3 {
		t.Errorf("Version should have 3 parts (major.minor.patch), got: %s", version)
	}
}

func TestPrintModuleReport(t *testing.T) {
	// This test ensures the function can be called without panicking
	// We can't easily test printed output without refactoring, 
	// but we can ensure it doesn't crash
	
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PrintModuleReport() panicked: %v", r)
		}
	}()
	
	PrintModuleReport()
}

func TestModulePathIsValid(t *testing.T) {
	path := GetModulePath()
	
	// More strict validation
	invalidChars := []string{" ", "\t", "\n", "@", "!", "#"}
	for _, char := range invalidChars {
		if strings.Contains(path, char) {
			t.Errorf("Module path contains invalid character '%s': %s", char, path)
		}
	}
}

func TestPackageInfoCompleteness(t *testing.T) {
	info := GetPackageInfo()
	
	// Should mention purpose
	lowerInfo := strings.ToLower(info)
	if !strings.Contains(lowerInfo, "purpose") && !strings.Contains(lowerInfo, "learning") {
		t.Error("Package info should describe its purpose")
	}
}

func TestVersionFormat(t *testing.T) {
	version := GetModuleVersion()
	
	// Remove 'v' prefix and check each part
	versionNum := strings.TrimPrefix(version, "v")
	parts := strings.Split(versionNum, ".")
	
	if len(parts) != 3 {
		t.Fatalf("Expected 3 version parts, got %d", len(parts))
	}
	
	// Each part should be numeric (basic check - just ensure they're not empty)
	for i, part := range parts {
		if part == "" {
			t.Errorf("Version part %d is empty", i)
		}
	}
}
