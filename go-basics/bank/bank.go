package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank!")
	for {

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Deposit Amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Amount must be greater that 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated ", accountBalance)
		case 3:
			var withDrawlAmount float64
			fmt.Print("Withdrawl Amount: ")
			fmt.Scan(&withDrawlAmount)
			if withDrawlAmount <= 0 {
				fmt.Println("Amount must be greater than 0")
			} else if withDrawlAmount > accountBalance {
				fmt.Println("Insufficient balance")
			} else {
				accountBalance -= withDrawlAmount
				fmt.Println("Balance updated ", accountBalance)
			}
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for choosing our bank")
			return

		}
		/*
			if choice == 1 {
				fmt.Println("Your balance is ", accountBalance)
			} else if choice == 2 {
				var depositAmount float64
				fmt.Print("Deposit Amount: ")
				fmt.Scan(&depositAmount)
				if depositAmount <= 0 {
					fmt.Println("Amount must be greater that 0")
					continue
				}
				accountBalance += depositAmount
				fmt.Println("Balance updated ", accountBalance)

			} else if choice == 3 {
				var withDrawlAmount float64
				fmt.Print("Withdrawl Amount: ")
				fmt.Scan(&withDrawlAmount)
				if withDrawlAmount <= 0 {
					fmt.Println("Amount must be greater than 0")
				} else if withDrawlAmount > accountBalance {
					fmt.Println("Insufficient balance")
				} else {
					accountBalance -= withDrawlAmount
					fmt.Println("Balance updated ", accountBalance)
				}

			} else {
				fmt.Println("Goodbye")
				break
			}
		*/
	}

}
