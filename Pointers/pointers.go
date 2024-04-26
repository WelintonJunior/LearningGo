package main

import "fmt"

func main() {

	// & antes da variavel para retornar o endereço da memoria
	// * antes da variavel para retornar o valor da variavel pointer

	numOfApples := 20
	numOfApplesPointer := &numOfApples
	realNumOfApples := numOfApplesAfterEat5(numOfApplesPointer)
	fmt.Print(realNumOfApples)
}

func numOfApplesAfterEat5(numOfApples *int) int {
	return *numOfApples - 5
}
