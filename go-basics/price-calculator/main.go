package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, .07, .10, .15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		// fm := cmdmanager.CMDManger{}
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// 	break
		// }

	}
	for _, doneChan := range doneChans {
		<-doneChan
	}

}
