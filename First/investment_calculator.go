package main

import (
	"fmt"
	"math"
)

func main() {

	const INFLATION_RATE = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Insira o valor do investimento: ", "")
	fmt.Scan(&investmentAmount)
	outputText("Insira a quantidade de anos: ", "")
	fmt.Scan(&years)
	outputText("Insira a expectativa do valor de volta: ", "")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	formattedDV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value: %.2f", futureRealValue)
	outputText(formattedDV, formattedRFV)

}

func outputText(text string, text2 string) {
	fmt.Print(text, text2)
}
