package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) // array created with initial size
	userNames[0] = "Julie"            // can't do with usual dynamic array
	userNames[1] = "Baby"
	// userNames[2] = "Fail" // fails because initial size of 2 is allocated
	userNames = append(userNames, "Max", "Mark", "Clara", "Tommy", "Emma")
	fmt.Println(userNames)
}
