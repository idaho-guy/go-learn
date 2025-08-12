package user

import (
	"errors"
	"fmt"
	"time"
)

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstName, lastName, and birthdate are required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func (u User) OutputUserDetails() {
	// don't need to de-reference for structs, sugar supplied by go
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	// don't need to de-reference for structs, sugar supplied by go
	u.FirstName = ""
	u.LastName = ""
}
