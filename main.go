package main

import (
	"example.com/structs-practice/note"
	"fmt"
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

	// store the note into a file
	fmt.Println(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
