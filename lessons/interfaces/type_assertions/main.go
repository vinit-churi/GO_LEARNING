package main


// Assignment
// Implement the getExpenseReport function.

// If the expense is an email, return the email's toAddress and the cost of the email.
// If the expense is an sms, return the sms's toPhoneNumber and its cost.
// If the expense has any other underlying type, return an empty string and 0.0 for the cost.

func getExpenseReport(e expense) (string, float64) {
	ex , emailOk := e.(email)
	if emailOk {
		return ex.toAddress, ex.cost()
	} 
	sms , smsOk := e.(sms)
	if smsOk {
		return  sms.toPhoneNumber, sms.cost()
	}
	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
