package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// factorial of 5 * 4 * 3, etc..

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
