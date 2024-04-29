package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	fmt.Println("-----------")
	arr := []string{"Programar", "Pescar", "Yasmin"}
	fmt.Println(arr)
	fmt.Println("-----------\n")

	fmt.Println("-----------")
	fmt.Println(arr[0])
	newArr := arr[1:]
	fmt.Println(newArr)
	fmt.Println("-----------\n")

	fmt.Println("-----------")
	newSlice := arr[:2]
	fmt.Println(newSlice)
	fmt.Println("-----------\n")

	fmt.Println("-----------")
	newSlice = arr[:3]
	fmt.Println(newSlice)
	fmt.Println("-----------\n")

	fmt.Println("-----------")
	goals := []string{"Rico", "Casar"}
	goals[1] = "Familia"
	goals = append(goals, "Filhos")
	fmt.Println(goals)
	fmt.Println("-----------\n")

	products := []Product{
		{
			title: "Terra",
			id:    1,
			price: 26.5,
		},
		{
			title: "Vidro",
			id:    2,
			price: 12.5,
		},
	}

	products = append(products, Product{
		title: "Fornalha",
		id:    3,
		price: 100.2,
	})

	fmt.Println(products)

}
