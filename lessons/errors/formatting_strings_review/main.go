package main

import (
	"fmt"
)

// Assignment
// We need better error logs for our backend developers to help them debug their code.

// Complete the getSMSErrorString() function. It should return a string with this format:

// SMS that costs $COST to be sent to 'RECIPIENT' can not be sent

// COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
// RECIPIENT is the stringified representation of the recipient's phone number
// Be sure to include the $ symbol, and wrap the recipient in single quotes.

func getSMSErrorString(cost float64, recipient string) string {
	// ?
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
}
