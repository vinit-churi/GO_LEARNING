package main

import "testing"

func TestBuildModulePath(t *testing.T) {
	got := BuildModulePath("Acme Org", "Billing API")
	if got != "github.com/acme-org/billing-api" {
		t.Fatalf("BuildModulePath() = %q", got)
	}
}

func TestIsPublicPackage(t *testing.T) {
	if !IsPublicPackage("payments") {
		t.Fatalf("IsPublicPackage(payments) = false")
	}
	if IsPublicPackage("internal-payments") {
		t.Fatalf("IsPublicPackage(internal-payments) = true")
	}
}

func TestIsPublicPackageRejectsBlank(t *testing.T) {
	if IsPublicPackage("   ") {
		t.Fatalf("IsPublicPackage(blank) = true")
	}
}
