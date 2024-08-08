package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	var limit int
	if membershipType == "premium" {
		limit = 1000
	} else {
		limit = 100
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: limit,
		},
	}
}
