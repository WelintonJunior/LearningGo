package main

import "fmt"

func main() {
	userNames := make([]string)

	userNames = append(userNames, "Welinton")
	userNames = append(userNames, "Yasmin")

	fmt.Println(userNames)

}
