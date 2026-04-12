//go:build starter

package constants

// Exercise 1: Basic Constants
// TODO: Declare Pi as an untyped constant with value 3.14159
const Pi = 3.14159
// TODO: Declare MaxRetries as a typed int constant with value 5
const MaxRetries int = 5
// TODO: Declare AppName as a string constant with value "GoLearning"
const AppName string = "GoLearning"

// TODO: Implement CalculateCircleArea to calculate the area of a circle
// Formula: area = π × radius²
func CalculateCircleArea(radius float64) float64 {
	// TODO: Implement this function
	return Pi * radius * radius
}

// Exercise 2: Typed vs Untyped Constants
// TODO: Declare UntypedNumber as an untyped constant with value 42
// TODO: Declare TypedNumber as an int64 constant with value 42
const UntypedNumber = 42
const TypedNumber int64 = 42

// TODO: Implement DemonstrateConstantTypes to show untyped constant flexibility
func DemonstrateConstantTypes() bool {
	// TODO: Assign UntypedNumber to int, float64, and int64 variables
	// TODO: Return true if all assignments work (they should)
	var isInt int = UntypedNumber
	var isFloat64 float64 = UntypedNumber
	var isInt64 int64 = UntypedNumber
	_ = isInt
	_ = isFloat64
	_ = isInt64
	return true
}

// Exercise 3: Const Blocks
// TODO: Create a const block for HTTP status codes:
// - StatusOK = 200
// - StatusCreated = 201
// - StatusBadRequest = 400
// - StatusUnauthorized = 401
// - StatusNotFound = 404
// - StatusInternalServerError = 500
const StatusOK, StatusCreated, StatusBadRequest, StatusUnauthorized, StatusNotFound, StatusInternalServerError = 200, 201, 400, 401, 404, 500

// TODO: Implement GetStatusMessage to return human-readable status messages
func GetStatusMessage(code int) string {
	// TODO: Use a switch statement to return appropriate messages
	// Return "Unknown Status" for codes not in the list
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown Status"
	}
	return ""
}

// Exercise 4: Using iota for Enumerations
// TODO: Create a const block using iota for days of the week:
// - Sunday = 0
// - Monday = 1
// - Tuesday = 2
// - Wednesday = 3
// - Thursday = 4
// - Friday = 5
// - Saturday = 6

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// TODO: Create a const block for file permissions using bit shifting:
// - ReadPermission = 1 << iota (equals 1)
// - WritePermission (equals 2)
// - ExecutePermission (equals 4)

const (
	ReadPermission = 1 << iota
	WritePermission
	ExecutePermission
)

// TODO: Implement IsWeekend to check if a day is Saturday or Sunday
func IsWeekend(day int) bool {
	// TODO: Implement this function
	if day == Sunday || day == Saturday {
		return true
	}
	return false
}

// TODO: Implement IsWeekday to check if a day is Monday through Friday
func IsWeekday(day int) bool {
	// TODO: Implement this function
	if day != Sunday && day != Saturday{
		return true
	}
	return false
}

// TODO: Implement HasPermission to check if user has required permission
// Use bitwise AND to check if the permission bit is set
func HasPermission(userPermissions, requiredPermission int) bool {
	// TODO: Implement this function
	if (userPermissions & requiredPermission) == requiredPermission {
		return true
	}
	return false
}
