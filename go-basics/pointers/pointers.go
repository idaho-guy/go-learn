package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age               // address of age, 0xc000090020
	fmt.Println("Age:", *agePointer) // * to de-reference pointer, display 32
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}
func getAdultYears(age *int) int { // copy of age will not be sent to function
	return *age - 18
}
