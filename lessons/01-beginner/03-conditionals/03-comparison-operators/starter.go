package main

import "fmt"

func CompareIntegers(a, b int) map[string]bool {
	return map[string]bool{
		"equal":            a == b,
		"notEqual":         a != b,
		"lessThan":         a < b,
		"greaterThan":      a > b,
		"lessOrEqual":      a <= b,
		"greaterOrEqual":   a >= b,
	}
}

func IsValidAge(age int) bool {
	return age >= 0 && age <= 150
}

func IsWeekend(day string) bool {
	return day == "Saturday" || day == "Sunday"
}

func IsNotEmpty(s string) bool {
	return !(s == "")
}

func CanVote(age int, isCitizen bool) bool {
	return age >= 18 && isCitizen
}

func NeedsUmbrella(isRaining bool, isSnowing bool) bool {
	return isRaining || isSnowing
}

func IsValidPassword(password string, confirmPassword string, length int) bool {
	return password == confirmPassword && length >= 8 && length <= 128
}

func CheckAccessLevel(age int, isMember bool, isPremium bool) string {
	if isPremium || (isMember && age >= 18) {
		return "full access"
	}
	if isMember && age < 18 {
		return "limited access"
	}

	return "no access"
}

func IsOutsideRange(value int, min int, max int) bool {
	inRange := value >= min && value <= max
	return !inRange
}

func main() {
	fmt.Println(CompareIntegers(5, 3))
	fmt.Println(IsValidAge(25))
	fmt.Println(IsWeekend("Saturday"))
	fmt.Println(IsNotEmpty("go"))
	fmt.Println(CanVote(21, true))
	fmt.Println(NeedsUmbrella(true, false))
	fmt.Println(IsValidPassword("secret123", "secret123", 9))
	fmt.Println(CheckAccessLevel(16, true, false))
	fmt.Println(IsOutsideRange(15, 1, 10))
}
