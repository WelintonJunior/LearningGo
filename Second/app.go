package main

import (
	"errors"
	"fmt"
	"os"
)

var resultsFile = "results.txt"

func main() {
	revenue, err1 := showAsnwer("What the Revenue: ")
	expenses, err2 := showAsnwer("What the Expenses: ")
	taxRate, err3 := showAsnwer("What the Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Print(err1)
		return
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("The EBT is: %.1f\nYour profit is: %.1f\nYour Ratio is: %.3f", ebt, profit, ratio)
}

func storeResults(ebtCalc, profitCalc, ratioCalc float64) {
	results := fmt.Sprintf("EBT: %.1f\n Profit: %.1f\n Ratio: %.3f", ebtCalc, profitCalc, ratioCalc)
	os.WriteFile(resultsFile, []byte(results), 0644)
}

func showAsnwer(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number")
	}
	return userInput, nil
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebtCalc := revenue - expenses
	profitCalc := ebtCalc * (1 - taxRate/100)
	ratioCalc := ebtCalc / profitCalc

	storeResults(ebtCalc, profitCalc, ratioCalc)

	return ebtCalc, profitCalc, ratioCalc

}
