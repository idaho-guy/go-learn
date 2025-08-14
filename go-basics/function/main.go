package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 5}
	doubled := transFormNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transFormNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transFormNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
