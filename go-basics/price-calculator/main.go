package main

import "example.com/price-calculator/prices"

func main() {
	taxRates := []float64{0, .07, .10, .15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()

	}

}
