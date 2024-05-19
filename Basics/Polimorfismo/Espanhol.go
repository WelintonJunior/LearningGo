package main

import "fmt"

type Espanhol struct{}

func (e Espanhol) Saudar(text string) {
	fmt.Println(text)
}
