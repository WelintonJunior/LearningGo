package main

import "fmt"

func main() {

	var isRunning int = 0
	var accountBalance = 3000.0

	for isRunning == 0 {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
		default:
			isRunning = 1
			fmt.Println("GoodBye!")
		}
	}

}
