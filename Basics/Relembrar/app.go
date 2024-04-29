package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Bem vindo ao LivroDosk")
	books := map[string]Books{}

	// manos := []int{2, 56, 2, 546, 2, 45, 3, 6, 567, 2, 456}

	// for _, mano := range manos {
	// 	fmt.Print(mano)
	// }

	for {
		showFirstMessages()
		choice := 0
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if len(books) > 0 {
				for _, book := range books {
					fmt.Printf("Nome: %v\nData de lançamento: %v\nPreço: %.2f\nQuantidade de folhas: %d", book.name, book.dataLancamento, book.preco, book.qtdFolhas)
				}
			} else {
				fmt.Println("Não há livros cadastrados!")
			}
		case 2:
			bookName := getUserInput("Nome do livro: ")
			bookData := getUserInput("Data de lançamento do livro: ")
			bookPreco, _ := strconv.ParseFloat(getUserInput("Preço do livro: "), 64)
			bookQtdFolhas, _ := strconv.ParseInt(getUserInput("Quantidade de folhas do livro: "), 10, 64)

			books[bookName] = Books{name: bookName, dataLancamento: bookData, qtdFolhas: bookQtdFolhas, preco: bookPreco}
			// books = append(books, Books{name: bookName, dataLancamento: bookData, preco: bookPreco, qtdFolhas: bookQtdFolhas})
		case 3:
			bookNameDelete := getUserInput("Digite o nome do livro a ser deletado: ")

			if _, exists := books[bookNameDelete]; exists {
				delete(books, bookNameDelete)
				fmt.Println("Livro deletado com sucesso!")
			} else {
				fmt.Println("Livro não encontrado!")
			}

		default:
			return
		}
	}

}

func showFirstMessages() {
	fmt.Println("\nO que você deseja fazer?")
	fmt.Println("1. Ver Livros")
	fmt.Println("2. Cadastrar Livro")
	fmt.Println("3. Deletar Livro")
	fmt.Println("4. Sair")
}

func getUserInput(text string) string {
	fmt.Println(text)
	var input string
	fmt.Scan(&input)
	return input
}
