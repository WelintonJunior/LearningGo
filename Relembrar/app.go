package main

import (
	"fmt"
	"math"
)

func main() {

	const INFLATION_RATE = 2.5
	investment := 0.0
	years := 0.0
	expected := 0.0

	fmt.Print("What is your investment: ")
	fmt.Scan(&investment)
	fmt.Print("What is your Year: ")
	fmt.Scan(&years)
	fmt.Print("What is your expected: ")
	fmt.Scan(&expected)

	futureValue := investment * math.Pow(1+expected/100, years)
	futureRealValue := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

}
