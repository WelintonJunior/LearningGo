package main

import "fmt"

func main() {

	revenue := showAsnwer("What the Revenue: ")
	expenses := showAsnwer("What the Expenses: ")
	taxRate := showAsnwer("What the Tax Rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("The EBT is: %.1f\nYour profit is: %.1f\nYour Ratio is: %.3f", ebt, profit, ratio)
}

func showAsnwer(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebtCalc := revenue - expenses
	profitCalc := ebtCalc * (1 - taxRate/100)
	ratioCalc := ebtCalc / profitCalc

	return ebtCalc, profitCalc, ratioCalc

}
