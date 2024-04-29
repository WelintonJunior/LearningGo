package main

import (
	"fmt"
	"strconv"

	"example.com/pratice3/copo"
)

type Copo copo.Copo

func main() {

	copos := map[string]Copo{}

	for {
		FirstWelcome()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:

			if len(copos) > 0 {
				for _, copo := range copos {
					fmt.Printf("Nome: %v\nDesenho: %v\nTamanho: %v\nMarca: %v\n\n", copo.Name, copo.Desenho, copo.Tamanho, copo.Marca)
				}
			} else {
				fmt.Println("Não há copos cadastrados!")
			}

		case 2:
			nomeValue := getUserInput("Nome: ")
			desenhoValue := getUserInput("Desenho: ")
			tamanhoValue, err := strconv.ParseFloat(getUserInput("Tamanho: "), 64)

			if err != nil {
				panic(err)
			}

			marcaValue := getUserInput("Marca: ")

			copos[nomeValue] = Copo{Name: nomeValue,
				Desenho: desenhoValue,
				Tamanho: float64(tamanhoValue),
				Marca:   marcaValue}
		case 3:
			nomeDeleteValue := getUserInput("Nome: ")
			if _, exists := copos[nomeDeleteValue]; exists {
				delete(copos, nomeDeleteValue)
			} else {
				println("Copo não existe")
			}

		default:
			return
		}

	}

}

func FirstWelcome() {
	fmt.Println("O que voce deseja fazer?\n")
	fmt.Println("1. Ver Copos")
	fmt.Println("2. Cadastrar Copos")
	fmt.Println("3. Deletar Copos")
	fmt.Println("4. Sair")
}

func getUserInput(txt string) string {
	fmt.Print(txt)
	var input string
	fmt.Scan(&input)
	return input
}
