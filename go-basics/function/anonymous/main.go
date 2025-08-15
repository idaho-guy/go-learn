package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 5
	})

	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	// using factor from funtion's outer scope, below is a closure
	// factor will be availble in double and triple always (locked in)
	return func(number int) int {
		return number * factor
	}

}
