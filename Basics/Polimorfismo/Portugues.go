package main

import "fmt"

type Portugues struct{}

func (p Portugues) Saudar(text string) {
	fmt.Println(text)
}
