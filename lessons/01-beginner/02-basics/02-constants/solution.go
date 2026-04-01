//go:build !starter

package constants

// Exercise 1: Basic Constants
const Pi = 3.14159
const MaxRetries int = 5
const AppName = "GoLearning"

func CalculateCircleArea(radius float64) float64 {
	return Pi * radius * radius
}

// Exercise 2: Typed vs Untyped Constants
const UntypedNumber = 42
const TypedNumber int64 = 42

func DemonstrateConstantTypes() bool {
	var i int = UntypedNumber
	var f float64 = UntypedNumber
	var i64 int64 = UntypedNumber

	return i == 42 && f == 42.0 && i64 == 42
}

// Exercise 3: Const Blocks
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

func GetStatusMessage(code int) string {
	switch code {
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusNotFound:
		return "Not Found"
	case StatusInternalServerError:
		return "Internal Server Error"
	default:
		return "Unknown Status"
	}
}

// Exercise 4: Using iota for Enumerations
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	ReadPermission = 1 << iota
	WritePermission
	ExecutePermission
)

func IsWeekend(day int) bool {
	return day == Sunday || day == Saturday
}

func IsWeekday(day int) bool {
	return day >= Monday && day <= Friday
}

func HasPermission(userPermissions, requiredPermission int) bool {
	return userPermissions&requiredPermission != 0
}
