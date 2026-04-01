//go:build !starter
// +build !starter

package main

import "fmt"

func main() {
	// Test Exercise 1
	fmt.Println("Module Path:", GetModulePath())
	fmt.Println()
	
	// Test Exercise 2
	fmt.Println(GetPackageInfo())
	fmt.Println()
	
	// Test Exercise 3
	fmt.Println("Version:", GetModuleVersion())
	fmt.Println()
	
	// Test Exercise 4
	PrintModuleReport()
}

// Exercise 1: Return the module path
func GetModulePath() string {
	return "github.com/yourname/go-learning"
}

// Exercise 2: Return package information
func GetPackageInfo() string {
	return "Package: main\nPurpose: Learning Go modules and packages"
}

// Exercise 3: Return the module version
func GetModuleVersion() string {
	return "v1.0.0"
}

// Exercise 4: Print a complete module report
func PrintModuleReport() {
	fmt.Println("=== Module Information ===")
	fmt.Println("Module Path:", GetModulePath())
	fmt.Println("Package: main")
	fmt.Println("Version:", GetModuleVersion())
	fmt.Println("Purpose: Learning Go modules and packages")
	fmt.Println("========================")
}
