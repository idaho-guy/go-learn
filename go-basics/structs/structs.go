package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	fn := getUserData("Please enter your first name: ")
	ln := getUserData("Please enter your last name: ")
	bd := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(fn, ln, bd)

	if err != nil {
		panic(err)
	}

	adminUser := user.NewAdmin("x@gmail.com", "pass")
	adminUser.User.OutputUserDetails()

	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
