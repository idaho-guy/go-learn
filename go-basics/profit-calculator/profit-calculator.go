package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// fmt.Println("ebt:", ebt)
	// fmt.Println("profit:", profit)
	// fmt.Println("ratio:", ratio)
	fmt.Printf("ebt: %v\nprofit: %.2f\nratio: %.2f\n", ebt, profit, ratio)

	// formattedData := fmt.Sprintf("ebt: %v\nprofit: %.2f\nratio: %.2f\n", ebt, profit, ratio)
	// fmt.Println(formattedData)

	// fmt.Printf(`
	// ebt: %v
	// profit: %.2f
	// ratio: %.2f`, ebt, profit, ratio)
	// fmt.Println()
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(prompt string) float64 {
	var num float64
	fmt.Print(prompt)
	fmt.Scan(&num)
	return num
}
