//go:build solution
// +build solution

package main

import (
	"bytes"
	"fmt"
	"io"
)

func WriteLine(w io.Writer, line string) error {
	_, err := fmt.Fprintln(w, line)
	return err
}

func EmitReport(w io.Writer, service string, count int) error {
	return WriteLine(w, fmt.Sprintf("service=%s count=%d", service, count))
}

func BufferReport(service string, count int) string {
	var buf bytes.Buffer
	_ = EmitReport(&buf, service, count)
	return buf.String()
}

func main() {
	fmt.Print(BufferReport("billing-api", 3))
}
