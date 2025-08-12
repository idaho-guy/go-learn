package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("You can reach us at", randomdata.PhoneNumber())
	for {
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
				fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
