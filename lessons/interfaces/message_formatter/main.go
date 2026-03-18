package main

import (
	"fmt"
)

// Assignment
// Implement the formatter interface with a method format that returns a formatted string.
// Define structs that satisfy the formatter interface: plainText, bold, code.
// The structs must all have a message field of type string
// plainText should return the message as is.
// bold should wrap the message in two asterisks (**) to simulate bold text (e.g., **message**).
// code should wrap the message in a single backtick (`) to simulate code block (e.g., `message`)

type formatter interface {
	format() string
}

type plainText struct {
	message string
}
type bold struct {
	message string
}
type code struct {
	message string
}


func (p plainText) format () string {
	return p.message
}

func (p bold) format () string {
	return fmt.Sprintf("**%v**", p.message)
}
func (p code) format () string {
	return fmt.Sprintf("`%v`", p.message)
}



// Don't Touch below this line
func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
