package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

}

func transformNumbers(numbers *[]int, trasnform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, trasnform(val))
	}
	return dNumbers
}

func getTransformerFunction(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
