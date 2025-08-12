package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age               // address of age, 0xc000090020
	fmt.Println("Age:", *agePointer) // * to de-reference pointer, display 32
	updateAgeToAdultYears(agePointer)
	fmt.Println(age)
}
func updateAgeToAdultYears(age *int) { // copy of age will not be sent to function
	*age = *age - 18 // override in memory
}
