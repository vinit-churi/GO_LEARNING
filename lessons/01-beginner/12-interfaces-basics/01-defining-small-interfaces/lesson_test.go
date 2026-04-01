package main

import "testing"

func TestEmailPreview(t *testing.T) {
	if got := (EmailAlert{Subject: "Daily report"}).Preview(); got != "email: Daily report" {
		t.Fatalf("Email Preview returned %q", got)
	}
}

func TestSMSPreview(t *testing.T) {
	if got := (SMSAlert{Message: "Build passed"}).Preview(); got != "sms: Build passed" {
		t.Fatalf("SMS Preview returned %q", got)
	}
}

func TestSendPreview(t *testing.T) {
	got := SendPreview(EmailAlert{Subject: "Daily report"})
	if got != "sending email: Daily report" {
		t.Fatalf("SendPreview returned %q", got)
	}
}
