//go:build starter

package variables

// Exercise1 demonstrates basic variable declaration with the var keyword and explicit types.
// TODO: Declare a variable 'name' of type string and assign it a value
// TODO: Declare a variable 'age' of type int and assign it a value
// TODO: Return both variables
func Exercise1() (string, int) {
	// Your code here
	return "", 0
}

// Exercise2 demonstrates short variable declaration using the := operator.
// TODO: Use := to declare 'temperature' with value 98.6
// TODO: Use := to declare 'isRaining' with value false
// TODO: Use := to declare 'city' with value "San Francisco"
// TODO: Return all three variables
func Exercise2() (float64, bool, string) {
	// Your code here
	return 0.0, false, ""
}

// Exercise3 demonstrates Go's zero values for different types.
// TODO: Declare an int variable without initialization
// TODO: Declare a float64 variable without initialization
// TODO: Declare a string variable without initialization
// TODO: Declare a bool variable without initialization
// TODO: Return all four variables (they will have their zero values)
func Exercise3() (int, float64, string, bool) {
	// Your code here
	return 0, 0.0, "", false
}

// Exercise4 demonstrates declaring multiple variables using a var block.
// TODO: Use a var () block to declare the following variables:
//   - productName (string) = "Laptop"
//   - price (float64) = 999.99
//   - inStock (bool) = true
//   - quantity (int) = 5
// TODO: Return all four variables
func Exercise4() (string, float64, bool, int) {
	// Your code here
	return "", 0.0, false, 0
}
