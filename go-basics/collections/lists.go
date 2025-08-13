package main

import (
	"fmt"
)

func main() {
	var productNames [4]string
	productNames[0] = "Dishwashing Liquid"
	prices := [4]float64{10.0, 9.9, 20.0, 12.3}
	fmt.Println(prices)
	fmt.Println(productNames)
	//featuredPrices := prices[1:3] // 9.9 20
	var featuredPrices = prices[:3] // 10 9.9 20
	fmt.Println(featuredPrices)
}
