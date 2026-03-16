// Update Users
// We need a way to differentiate between standard and premium users. When a new user is created, they need a membership type, and that type will determine the message character limit.

// Assignment
// Create a new struct called Membership, it should have:
// A Type string field
// A MessageCharLimit integer field
// Update the User struct to embed a Membership.
// Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is a "premium" member, the MessageCharLimit should be 1000, otherwise, it should only be 100.

package main
type Membership struct {
	Type string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}



func newUser(name string, membershipType string) User {
	var value int = 100
	if membershipType == "premium" {
		value = 1000
	}

	membership := Membership {
		Type: membershipType,
		MessageCharLimit: value,
	}

	new_user := User {
		Name: name,
		Membership: membership,
	}

	return new_user
}
