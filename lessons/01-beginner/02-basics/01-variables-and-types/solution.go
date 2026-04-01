//go:build !starter

package variables

// Exercise1 demonstrates basic variable declaration with the var keyword and explicit types.
func Exercise1() (string, int) {
	var name string = "Alice"
	var age int = 30
	return name, age
}

// Exercise2 demonstrates short variable declaration using the := operator.
func Exercise2() (float64, bool, string) {
	temperature := 98.6
	isRaining := false
	city := "San Francisco"
	return temperature, isRaining, city
}

// Exercise3 demonstrates Go's zero values for different types.
func Exercise3() (int, float64, string, bool) {
	var num int
	var decimal float64
	var text string
	var flag bool
	return num, decimal, text, flag
}

// Exercise4 demonstrates declaring multiple variables using a var block.
func Exercise4() (string, float64, bool, int) {
	var (
		productName string  = "Laptop"
		price       float64 = 999.99
		inStock     bool    = true
		quantity    int     = 5
	)
	return productName, price, inStock, quantity
}
