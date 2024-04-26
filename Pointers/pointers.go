package main

import "fmt"

func main() {

	// & antes da variavel para retornar o endereço da memoria
	// * antes da variavel para retornar o valor da variavel pointer

	numOfApples := 20
	numOfApplesPointer := &numOfApples
	numOfApplesAfterEat5(numOfApplesPointer)
	fmt.Print(numOfApples)
}

func numOfApplesAfterEat5(numOfApples *int) {
	// return *numOfApples - 5
	*numOfApples = *numOfApples - 5
}
