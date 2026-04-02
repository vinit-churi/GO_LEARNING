package main

import "testing"

func TestFullName(t *testing.T) {
	info := AuditInfo{Owner: "Ria", Team: "platform"}
	if got := info.FullName(); got != "Ria @ platform" {
		t.Fatalf("FullName() = %q", got)
	}
}

func TestSummary(t *testing.T) {
	service := DeployService{AuditInfo: AuditInfo{Owner: "Ria", Team: "platform"}, Name: "billing-api"}
	if got := service.Summary(); got != "billing-api owned by Ria @ platform" {
		t.Fatalf("Summary() = %q", got)
	}
}

func TestTouchUpdatesEmbeddedOwner(t *testing.T) {
	service := DeployService{AuditInfo: AuditInfo{Owner: "Ria", Team: "platform"}, Name: "billing-api"}
	service.Touch(" Riya ")
	if service.Owner != "Riya" {
		t.Fatalf("Touch() left owner as %q", service.Owner)
	}
}
