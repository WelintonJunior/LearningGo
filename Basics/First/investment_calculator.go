package main

import (
	"fmt"
	"math"
)

const INFLATION_RATE = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Insira o valor do investimento: ", "")
	fmt.Scan(&investmentAmount)
	outputText("Insira a quantidade de anos: ", "")
	fmt.Scan(&years)
	outputText("Insira a expectativa do valor de volta: ", "")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedDV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value: %.2f", futureRealValue)
	outputText(formattedDV, formattedRFV)

}

func outputText(text string, text2 string) {
	fmt.Print(text, text2)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+INFLATION_RATE/100, years)

	return fv, frv
}
