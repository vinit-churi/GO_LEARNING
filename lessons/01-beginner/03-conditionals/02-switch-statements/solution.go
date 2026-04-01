//go:build !starter
// +build !starter

package main

import "fmt"

// Exercise 1: Day of the Week
// Returns "Weekday", "Weekend", or "Invalid day"
func GetDayType(day string) string {
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "Weekday"
	case "Saturday", "Sunday":
		return "Weekend"
	default:
		return "Invalid day"
	}
}

// Exercise 2: Grade Evaluator
// Returns appropriate message for each grade
func GetGradeMessage(grade string) string {
	switch grade {
	case "A", "B":
		return "Excellent work!"
	case "C":
		return "Good job!"
	case "D":
		return "You passed, but consider studying more."
	case "F":
		return "You need to retake this course."
	default:
		return "Invalid grade"
	}
}

// Exercise 3: Number Classifier
// Uses expression-less switch to classify numbers
func ClassifyNumber(n int) string {
	switch {
	case n < 0:
		return "Negative"
	case n == 0:
		return "Zero"
	case n > 0 && n <= 10:
		return "Small positive"
	case n > 10 && n <= 100:
		return "Medium positive"
	case n > 100:
		return "Large positive"
	}
	return "" // Unreachable, but satisfies compiler
}

// Exercise 4: Season Describer
// Demonstrates fallthrough keyword
func DescribeSeason(season string) string {
	result := ""
	switch season {
	case "Spring":
		return "Flowers bloom"
	case "Summer":
		result = "Hot and sunny\n"
		fallthrough
	case "Fall", "Autumn":
		result += "Leaves change color"
		return result
	case "Winter":
		return "Cold and snowy"
	default:
		return "Unknown season"
	}
}

// Exercise 5: HTTP Status Classifier
// Classifies HTTP status codes by range
func ClassifyHTTPStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return "Success"
	case code >= 300 && code <= 399:
		return "Redirection"
	case code >= 400 && code <= 499:
		return "Client Error"
	case code >= 500 && code <= 599:
		return "Server Error"
	default:
		return "Unknown Status"
	}
}

func main() {
	// Exercise 1 examples
	fmt.Println("Exercise 1: Day of the Week")
	fmt.Println("Monday:", GetDayType("Monday"))
	fmt.Println("Saturday:", GetDayType("Saturday"))
	fmt.Println("Funday:", GetDayType("Funday"))
	fmt.Println()

	// Exercise 2 examples
	fmt.Println("Exercise 2: Grade Evaluator")
	fmt.Println("Grade A:", GetGradeMessage("A"))
	fmt.Println("Grade B:", GetGradeMessage("B"))
	fmt.Println("Grade C:", GetGradeMessage("C"))
	fmt.Println("Grade F:", GetGradeMessage("F"))
	fmt.Println()

	// Exercise 3 examples
	fmt.Println("Exercise 3: Number Classifier")
	fmt.Println("-5:", ClassifyNumber(-5))
	fmt.Println("0:", ClassifyNumber(0))
	fmt.Println("5:", ClassifyNumber(5))
	fmt.Println("50:", ClassifyNumber(50))
	fmt.Println("200:", ClassifyNumber(200))
	fmt.Println()

	// Exercise 4 examples
	fmt.Println("Exercise 4: Season Describer")
	fmt.Println("Spring:", DescribeSeason("Spring"))
	fmt.Println("Summer:", DescribeSeason("Summer"))
	fmt.Println("Fall:", DescribeSeason("Fall"))
	fmt.Println("Winter:", DescribeSeason("Winter"))
	fmt.Println()

	// Exercise 5 examples
	fmt.Println("Exercise 5: HTTP Status Classifier")
	fmt.Println("200:", ClassifyHTTPStatus(200))
	fmt.Println("301:", ClassifyHTTPStatus(301))
	fmt.Println("404:", ClassifyHTTPStatus(404))
	fmt.Println("500:", ClassifyHTTPStatus(500))
	fmt.Println("999:", ClassifyHTTPStatus(999))
}
