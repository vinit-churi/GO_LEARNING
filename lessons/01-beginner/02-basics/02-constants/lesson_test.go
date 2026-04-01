package constants

import (
	"math"
	"testing"
)

// Exercise 1 Tests: Basic Constants
func TestExercise1_Constants(t *testing.T) {
	tests := []struct {
		name     string
		constant interface{}
		expected interface{}
	}{
		{"Pi value", Pi, 3.14159},
		{"MaxRetries value", MaxRetries, 5},
		{"AppName value", AppName, "GoLearning"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, tt.constant)
			}
		})
	}
}

func TestExercise1_CalculateCircleArea(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected float64
		delta    float64
	}{
		{"radius 2.0", 2.0, 12.56636, 0.00001},
		{"radius 5.0", 5.0, 78.53975, 0.00001},
		{"radius 0.0", 0.0, 0.0, 0.00001},
		{"radius 1.0", 1.0, 3.14159, 0.00001},
		{"radius 10.0", 10.0, 314.159, 0.001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateCircleArea(tt.radius)
			if math.Abs(result-tt.expected) > tt.delta {
				t.Errorf("CalculateCircleArea(%v) = %v, expected %v", tt.radius, result, tt.expected)
			}
		})
	}
}

// Exercise 2 Tests: Typed vs Untyped Constants
func TestExercise2_ConstantValues(t *testing.T) {
	if UntypedNumber != 42 {
		t.Errorf("UntypedNumber should be 42, got %v", UntypedNumber)
	}
	if TypedNumber != 42 {
		t.Errorf("TypedNumber should be 42, got %v", TypedNumber)
	}
}

func TestExercise2_DemonstrateConstantTypes(t *testing.T) {
	result := DemonstrateConstantTypes()
	if !result {
		t.Error("DemonstrateConstantTypes() should return true")
	}
}

func TestExercise2_UntypedFlexibility(t *testing.T) {
	var i int = UntypedNumber
	var f float64 = UntypedNumber
	var i64 int64 = UntypedNumber

	if i != 42 {
		t.Errorf("UntypedNumber as int should be 42, got %v", i)
	}
	if f != 42.0 {
		t.Errorf("UntypedNumber as float64 should be 42.0, got %v", f)
	}
	if i64 != 42 {
		t.Errorf("UntypedNumber as int64 should be 42, got %v", i64)
	}
}

// Exercise 3 Tests: Const Blocks
func TestExercise3_StatusConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"StatusOK", StatusOK, 200},
		{"StatusCreated", StatusCreated, 201},
		{"StatusBadRequest", StatusBadRequest, 400},
		{"StatusUnauthorized", StatusUnauthorized, 401},
		{"StatusNotFound", StatusNotFound, 404},
		{"StatusInternalServerError", StatusInternalServerError, 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %v, expected %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestExercise3_GetStatusMessage(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected string
	}{
		{"OK", StatusOK, "OK"},
		{"Created", StatusCreated, "Created"},
		{"Bad Request", StatusBadRequest, "Bad Request"},
		{"Unauthorized", StatusUnauthorized, "Unauthorized"},
		{"Not Found", StatusNotFound, "Not Found"},
		{"Internal Server Error", StatusInternalServerError, "Internal Server Error"},
		{"Unknown Status", 999, "Unknown Status"},
		{"Another Unknown", 100, "Unknown Status"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStatusMessage(tt.code)
			if result != tt.expected {
				t.Errorf("GetStatusMessage(%v) = %q, expected %q", tt.code, result, tt.expected)
			}
		})
	}
}

// Exercise 4 Tests: Using iota for Enumerations
func TestExercise4_WeekdayValues(t *testing.T) {
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"Sunday", Sunday, 0},
		{"Monday", Monday, 1},
		{"Tuesday", Tuesday, 2},
		{"Wednesday", Wednesday, 3},
		{"Thursday", Thursday, 4},
		{"Friday", Friday, 5},
		{"Saturday", Saturday, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %v, expected %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestExercise4_PermissionValues(t *testing.T) {
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"ReadPermission", ReadPermission, 1},
		{"WritePermission", WritePermission, 2},
		{"ExecutePermission", ExecutePermission, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %v, expected %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestExercise4_IsWeekend(t *testing.T) {
	tests := []struct {
		name     string
		day      int
		expected bool
	}{
		{"Sunday is weekend", Sunday, true},
		{"Monday is not weekend", Monday, false},
		{"Tuesday is not weekend", Tuesday, false},
		{"Wednesday is not weekend", Wednesday, false},
		{"Thursday is not weekend", Thursday, false},
		{"Friday is not weekend", Friday, false},
		{"Saturday is weekend", Saturday, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsWeekend(tt.day)
			if result != tt.expected {
				t.Errorf("IsWeekend(%v) = %v, expected %v", tt.day, result, tt.expected)
			}
		})
	}
}

func TestExercise4_IsWeekday(t *testing.T) {
	tests := []struct {
		name     string
		day      int
		expected bool
	}{
		{"Sunday is not weekday", Sunday, false},
		{"Monday is weekday", Monday, true},
		{"Tuesday is weekday", Tuesday, true},
		{"Wednesday is weekday", Wednesday, true},
		{"Thursday is weekday", Thursday, true},
		{"Friday is weekday", Friday, true},
		{"Saturday is not weekday", Saturday, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsWeekday(tt.day)
			if result != tt.expected {
				t.Errorf("IsWeekday(%v) = %v, expected %v", tt.day, result, tt.expected)
			}
		})
	}
}

func TestExercise4_HasPermission(t *testing.T) {
	tests := []struct {
		name               string
		userPermissions    int
		requiredPermission int
		expected           bool
	}{
		{"Read only - has read", ReadPermission, ReadPermission, true},
		{"Read only - no write", ReadPermission, WritePermission, false},
		{"Read only - no execute", ReadPermission, ExecutePermission, false},
		{"Write only - no read", WritePermission, ReadPermission, false},
		{"Write only - has write", WritePermission, WritePermission, true},
		{"Read+Write - has read", ReadPermission | WritePermission, ReadPermission, true},
		{"Read+Write - has write", ReadPermission | WritePermission, WritePermission, true},
		{"Read+Write - no execute", ReadPermission | WritePermission, ExecutePermission, false},
		{"All permissions - has read", ReadPermission | WritePermission | ExecutePermission, ReadPermission, true},
		{"All permissions - has write", ReadPermission | WritePermission | ExecutePermission, WritePermission, true},
		{"All permissions - has execute", ReadPermission | WritePermission | ExecutePermission, ExecutePermission, true},
		{"No permissions - no read", 0, ReadPermission, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasPermission(tt.userPermissions, tt.requiredPermission)
			if result != tt.expected {
				t.Errorf("HasPermission(%v, %v) = %v, expected %v",
					tt.userPermissions, tt.requiredPermission, result, tt.expected)
			}
		})
	}
}

// Additional integration tests
func TestIntegration_ConstantsInExpressions(t *testing.T) {
	// Test that Pi can be used in calculations with different numeric types
	var radius32 float32 = 2.0
	area32 := float32(Pi) * radius32 * radius32
	if math.Abs(float64(area32)-12.56636) > 0.01 {
		t.Errorf("Pi should work with float32, got area %v", area32)
	}

	// Test that UntypedNumber can be used in different contexts
	var sum = UntypedNumber + 10
	if sum != 52 {
		t.Errorf("UntypedNumber + 10 should be 52, got %v", sum)
	}
}

func TestIntegration_BitwisePermissions(t *testing.T) {
	// Combine multiple permissions
	fullAccess := ReadPermission | WritePermission | ExecutePermission

	if fullAccess != 7 {
		t.Errorf("Full access should be 7 (binary: 111), got %v", fullAccess)
	}

	// Check individual permissions in combined value
	if !HasPermission(fullAccess, ReadPermission) {
		t.Error("Full access should include read permission")
	}
	if !HasPermission(fullAccess, WritePermission) {
		t.Error("Full access should include write permission")
	}
	if !HasPermission(fullAccess, ExecutePermission) {
		t.Error("Full access should include execute permission")
	}

	// Remove a permission
	noExecute := fullAccess &^ ExecutePermission // bitwise AND NOT
	if HasPermission(noExecute, ExecutePermission) {
		t.Error("After removing execute, it should not be present")
	}
	if !HasPermission(noExecute, ReadPermission) || !HasPermission(noExecute, WritePermission) {
		t.Error("Read and write should still be present after removing execute")
	}
}
