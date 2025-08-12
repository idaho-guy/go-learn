package main

import (
	"fmt"
	"time"
)

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

	appUser := User{
		firstName: fn,
		lastName:  ln,
		birthDate: bd,
		createdAt: time.Now(),
	}

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
