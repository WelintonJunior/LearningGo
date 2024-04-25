package main

import (
	"fmt"
	"math"
)

func main() {

	const INFLATION_RATE = 2.5
	investmentAmount := 0.0
	years := 0.0
	expectedReturnRate := 0.0

	fmt.Print("Insira o valor do investimento: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Insira a quantidade de anos: ")
	fmt.Scan(&years)
	fmt.Print("Insira a expectativa do valor de volta: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	fmt.Println("Future Value: ", math.Round(futureValue))
	fmt.Println("Future Real Value: ", math.Round(futureRealValue))

}
