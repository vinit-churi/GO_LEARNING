package main


// Assignment
// Complete the maxMessages function. The parameter thresh is your total budget in pennies. Return the maximum number of whole messages you can send without the total cost ever exceeding thresh.

// Each message costs 100 pennies, plus an additional fee. The fee structure is:

// 1st message: 100 + 0
// 2nd message: 100 + 1
// 3rd message: 100 + 2
// 4th message: 100 + 3

func maxMessages(thresh int) int {
	total :=0
	// for ; thresh > 
	for ; ; {
		cost := 100 + total
		if thresh < cost {
			break;
		}
		thresh -= cost
		total++
	}
	return total
}
