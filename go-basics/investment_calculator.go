package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64
	var x float64

	prompt("X: ", &x)
	fmt.Println("x=", x)
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	// futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	futureValue, futureRealValue := calculateFutureValueAlt(investmentAmount, expectedReturnRate, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func prompt(text string, value *float64) {
	fmt.Print(text)
	fmt.Scan(value)
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fvWithInflation := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, fvWithInflation
}

func calculateFutureValueAlt(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
