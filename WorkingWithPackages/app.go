package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	var isRunning int = 0

	for isRunning == 0 {
		var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
		if err != nil {
			panic(err)
		}

		fmt.Println("Welcome to Go Bank!")
		presentOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your Balace is: %.1f", accountBalance)
		case 2:
			fmt.Println("How much?: ")
			depositedMoney := 0.0
			fmt.Scan(&depositedMoney)

			if depositedMoney <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositedMoney
			fmt.Printf("Balance Updated! New Amount: %.1f", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("How much?: ")
			withdrawMoney := 0.0
			fmt.Scan(&withdrawMoney)

			if withdrawMoney <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawMoney > accountBalance {
				fmt.Println("You don't have this money in your balance!")
				continue
			}

			accountBalance -= withdrawMoney
			fmt.Printf("Balance Updated! New Amount: %.1f", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			isRunning = 1
			fmt.Println("GoodBye!")
			return
		}
	}

}
