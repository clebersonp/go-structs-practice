package main

import (
	"bufio"
	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ask user for the data
	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todoData, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// store in a note structure
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoData.Display()

	// store the note into a file
	err = todoData.Save()
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}
	fmt.Println("Saving the todo succeeded!")

	userNote.Display()

	// store the note into a file
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Saving the note succeeded!")
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
