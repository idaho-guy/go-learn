package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	prices = append(prices, 5.99) // returns new array with 3 elements
	// remove element(s) by making new slices
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string
// 	productNames[0] = "Dishwashing Liquid"
// 	prices := [4]float64{10.0, 9.9, 20.0, 12.3}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	//featuredPrices := prices[1:3] // 9.9 20
// 	var featuredPrices = prices[:3] // 10 9.9 20
// 	featuredPrices[0] = 13.45
// 	fmt.Println(prices)                                   // 13.45 9.9 20 12.3, slice is ref to actual array so changing changes the original array
// 	fmt.Println(len(featuredPrices), cap(featuredPrices)) // 3 4, cap is based on size of array that the slice was made from
// 	highLighted := featuredPrices[:1]
// 	fmt.Println(len(highLighted), cap(highLighted)) // 1 4
// 	highLighted = highLighted[:3]
// 	fmt.Println(len(highLighted), cap(highLighted)) // 3 4

// }
