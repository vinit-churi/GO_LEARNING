//go:build !windows

package main

import "runtime"

func taggedPlatform() string {
	return runtime.GOOS
}
