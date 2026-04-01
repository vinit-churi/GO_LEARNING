//go:build starter
// +build starter

package main

import "fmt"

// Exercise 1: Day of the Week
// Write a function that returns "Weekday", "Weekend", or "Invalid day"
func GetDayType(day string) string {
	// TODO: Implement using a switch statement
	// Use multiple values in one case for weekdays (Monday-Friday)
	// Use multiple values in one case for weekend (Saturday-Sunday)
	// Use default for invalid input
	return ""
}

// Exercise 2: Grade Evaluator
// Write a function that returns grade messages
func GetGradeMessage(grade string) string {
	// TODO: Implement using a switch statement
	// "A" or "B" should return "Excellent work!"
	// "C" should return "Good job!"
	// "D" should return "You passed, but consider studying more."
	// "F" should return "You need to retake this course."
	// Use default for invalid grades
	return ""
}

// Exercise 3: Number Classifier
// Write a function using expression-less switch (switch without an expression)
func ClassifyNumber(n int) string {
	// TODO: Implement using switch {}  (no expression after switch)
	// n < 0 → "Negative"
	// n == 0 → "Zero"
	// n > 0 && n <= 10 → "Small positive"
	// n > 10 && n <= 100 → "Medium positive"
	// n > 100 → "Large positive"
	return ""
}

// Exercise 4: Season Describer
// Write a function that uses fallthrough for the Summer case
func DescribeSeason(season string) string {
	// TODO: Implement using switch with fallthrough
	// "Spring" → "Flowers bloom"
	// "Summer" → "Hot and sunny" + fallthrough to Fall case
	// "Fall" or "Autumn" → "Leaves change color"
	// "Winter" → "Cold and snowy"
	// Use default for unknown seasons
	// Hint: You'll need a variable to accumulate the result for Summer
	return ""
}

// Exercise 5: HTTP Status Classifier
// Write a function that classifies HTTP status codes by range
func ClassifyHTTPStatus(code int) string {
	// TODO: Implement using expression-less switch with ranges
	// 200-299 → "Success"
	// 300-399 → "Redirection"
	// 400-499 → "Client Error"
	// 500-599 → "Server Error"
	// Use default for unknown status codes
	return ""
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
