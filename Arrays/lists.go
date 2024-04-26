package main

import (
	"fmt"
)

func main() {

	var numbers [5]int = [5]int{3, 2, 4, 6, 6}

	for i := 0; i < len(numbers); i++ {
		fmt.Print("Number: ", numbers[i], "\n")
	}

	featuredNumbers := numbers[1:3]
	fmt.Print(featuredNumbers)

}
