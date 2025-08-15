package cmdmanager

import "fmt"

type CMDManger struct {
}

func (cmd CMDManger) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)

	}

	return prices, nil
}

func (cmd CMDManger) WriteResult(data interface{}) error {
	fmt.Printf(fmt.Sprintf("%v", data))
	return nil
}
