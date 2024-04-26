package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/pratice1/note"
	"example.com/pratice1/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display() error
}

func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		panic(err)
	}

	todo.Display()
	saveData(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

	userNote.Display()
	saveData(userNote)
}

func saveData(data saver) {
	err := data.Save()

	if err != nil {
		panic(err)
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(text string) string {
	fmt.Printf("%v ", text)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
