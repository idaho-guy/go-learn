package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) // array created with initial size
	userNames[0] = "Julie"            // can't do with usual dynamic array
	userNames[1] = "Baby"
	// userNames[2] = "Fail" // fails because initial size of 2 is allocated
	userNames = append(userNames, "Max", "Mark", "Clara", "Tommy", "Emma")
	// when adding Tommy and Emma, all new array is created
	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3) // pre-allocate map for 3 elements
	courseRatings["go"] = 4.6
	courseRatings["react"] = 4.7
}
