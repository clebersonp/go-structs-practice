package main

import (
	"bufio"
	"example.com/structs-practice/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ask user for the data
	title, content := getNoteData()

	// store in a note structure
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	// store the note into a file
	err = userNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	userInput = strings.TrimSuffix(userInput, "\n")
	return strings.TrimSuffix(userInput, "\r")
}
