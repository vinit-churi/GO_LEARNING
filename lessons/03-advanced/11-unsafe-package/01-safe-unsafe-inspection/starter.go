//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"unsafe"
)

type Header struct {
	Version uint16
	Payload uint64
}

func HeaderSize() uintptr {
	return unsafe.Sizeof(Header{})
}

func HeaderAlignment() uintptr {
	return unsafe.Alignof(Header{})
}

func PayloadOffset() uintptr {
	return unsafe.Offsetof(Header{}.Payload)
}

func main() {
	fmt.Println(HeaderSize(), HeaderAlignment(), PayloadOffset())
}
