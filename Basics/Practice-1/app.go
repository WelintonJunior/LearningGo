package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/pratice1/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

	userNote.Display()
	err = userNote.Save()

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
