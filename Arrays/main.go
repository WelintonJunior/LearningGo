package main

import "fmt"

func main() {
	userNames := []string{}

	userNames = append(userNames, "Welinton")
	userNames = append(userNames, "Yasmin")

	fmt.Println(userNames)

	for index, userName := range userNames {
		fmt.Printf("Index: %v\nNome: %v\n", index, userName)
	}

}
