package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 5}
	doubled := transFormNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transFormNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transFormNumbers(numbers *[]int, transform func(int) int) []int {
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
