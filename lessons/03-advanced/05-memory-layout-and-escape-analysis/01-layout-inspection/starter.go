//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"unsafe"
)

type Record struct {
	Flag  bool
	Value int64
	Code  byte
}

func RecordSize() uintptr {
	return unsafe.Sizeof(Record{})
}

func ValueOffset() uintptr {
	return unsafe.Offsetof(Record{}.Value)
}

func HasPadding() bool {
	return ValueOffset() > unsafe.Sizeof(Record{}.Flag)
}

func main() {
	fmt.Println(RecordSize(), ValueOffset(), HasPadding())
}
