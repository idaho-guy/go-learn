package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// factorial of 5 * 4 * 3, etc..

func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}
