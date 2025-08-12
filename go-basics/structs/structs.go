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
	outputUserDetails(&appUser)

	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)
}

func outputUserDetails(user *User) {
	// don't need to de-reference for structs, sugar supplied by go
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
