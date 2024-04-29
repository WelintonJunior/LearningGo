package lists

import "fmt"

// Dynamic arrays
func main() {

	prices := []float64{10.99, 8.99}
	//adidicionar valor para o prices
	prices = append(prices, 5)

	// criar outro array com o novo valor
	// updatedPrices := append(prices, 5)

	// prices = prices[:1]

	// fmt.Print( /*updatedPrices, */ prices)

	discountPrices := []float64{101.99, 80.88, 20.59}
	prices = append(prices, discountPrices...)

	fmt.Print( /*updatedPrices, */ prices)

}

// func main() {

// 	var numbers [5]int = [5]int{3, 2, 4, 6, 6}

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Print("Number: ", numbers[i], "\n")
// 	}

// 	// featuredNumbers := numbers[1:3]
// 	fmt.Print(numbers, cap(numbers), len(numbers))

// }
