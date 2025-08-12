package main

import (
	"fmt"
	"time"
)

// func newUser(firstName, lastName, birthDate string) User {
// 	return User{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthDate: birthDate,
// 		createdAt: time.Now(),
// 	}
// }

func newUser(firstName, lastName, birthDate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) outputUserDetails() {
	// don't need to de-reference for structs, sugar supplied by go
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUserName() {
	// don't need to de-reference for structs, sugar supplied by go
	u.firstName = ""
	u.lastName = ""
}

func main() {

	fn := getUserData("Please enter your first name: ")
	ln := getUserData("Please enter your last name: ")
	bd := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(fn, ln, bd)

	// ... do something awesome with that gathered data!
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
