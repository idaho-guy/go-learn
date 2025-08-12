package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age               // address of age, 0xc000090020
	fmt.Println("Age:", *agePointer) // * to de-reference pointer, display 32
	adultYears := getAdultYears((age))
	fmt.Println(adultYears)
}
func getAdultYears(age int) int {
	return age - 18
}
