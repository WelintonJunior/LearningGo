package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("What the revenue?: ")
	fmt.Scan(&revenue)
	fmt.Print("What the expenses?: ")
	fmt.Scan(&expenses)
	fmt.Print("What the tax rate?: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("The EBT is: ", ebt)
	fmt.Println("Your Profit is: ", profit)
	fmt.Println("The Ratio is: ", ratio)

}
