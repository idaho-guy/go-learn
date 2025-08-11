package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRation: %.3f\n", ebt, profit, ratio)
	err := os.WriteFile("results.txt", []byte(results), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// fmt.Println("ebt:", ebt)
	// fmt.Println("profit:", profit)
	// fmt.Println("ratio:", ratio)
	fmt.Printf("ebt: %v\nprofit: %.2f\nratio: %.2f\n", ebt, profit, ratio)
	storeResults(ebt, profit, ratio)
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

func getUserInput(prompt string) (float64, error) {
	var num float64
	fmt.Print(prompt)
	fmt.Scan(&num)
	if num <= 0 {
		return 0, errors.New("Value must be a positive number")
	}
	return num, nil
}
