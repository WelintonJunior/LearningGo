package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	accountBalance, err := os.ReadFile(accountBalanceFile)

	// Nil quer dizer que não tem erro, se for diferente de nil tem erro
	if err != nil {
		return 0, errors.New("Failed to find balance file")
	}

	balanceText := string(accountBalance)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	var isRunning int = 0

	for isRunning == 0 {
		var accountBalance, err = getBalanceFromFile()

		if err != nil {
			panic(err)
		}

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			isRunning = 1
			fmt.Println("GoodBye!")
			return
		}
	}

}
