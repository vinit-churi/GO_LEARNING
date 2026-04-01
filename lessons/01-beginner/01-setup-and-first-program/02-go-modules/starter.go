//go:build starter
// +build starter

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
	// TODO: Return a valid module path (e.g., "github.com/username/project")
	return ""
}

// Exercise 2: Return package information
func GetPackageInfo() string {
	// TODO: Return formatted package information
	// Should include package name and purpose
	return ""
}

// Exercise 3: Return the module version
func GetModuleVersion() string {
	// TODO: Return a semantic version string (e.g., "v1.0.0")
	return ""
}

// Exercise 4: Print a complete module report
func PrintModuleReport() {
	// TODO: Print a formatted report combining all the above information
	// Use the functions above to get the data
}
