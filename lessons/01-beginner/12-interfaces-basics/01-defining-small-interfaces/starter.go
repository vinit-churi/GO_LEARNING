//go:build !solution
// +build !solution

package main

import "fmt"

type Previewer interface {
	Preview() string
}

type EmailAlert struct {
	Subject string
}

type SMSAlert struct {
	Message string
}

func (e EmailAlert) Preview() string {
	return "email: " + e.Subject
}

func (s SMSAlert) Preview() string {
	return "sms: " + s.Message
}

func SendPreview(p Previewer) string {
	return "sending " + p.Preview()
}

func main() {
	fmt.Println(SendPreview(EmailAlert{Subject: "Daily report"}))
	fmt.Println(SendPreview(SMSAlert{Message: "Build passed"}))
}
