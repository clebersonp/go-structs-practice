package main

import (
	"bufio"
	"example.com/structs-practice/note"
	"example.com/structs-practice/report"
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

	// todoData
	err = outputData(todoData)
	if err != nil {
		return
	}

	// userNote
	err = outputData(userNote)
	if err != nil {
		return
	}

	fmt.Println("---------------------------")
	printSomething(1)
	printSomething(5.2)
	printSomething("text")
	printSomething(todoData)
	printSomething(userNote)
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

func outputData(data report.Outputtable) error {
	data.Display()
	// store the note into a file
	// using interface
	return saveData(data)
}

func saveData(data report.Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving failed")
		return err
	}
	fmt.Println("Saving succeeded!")
	return nil
}

// printSomething to explaining use of interface{} or just its alias 'any'
// type any = interface{}
// any type is allowed when using 'any' or 'interface{}' type
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64", value)
	case string:
		fmt.Println("string:", value)
	case todo.Todo:
		fmt.Println("todo.Todo:", value)
	case note.Note:
		fmt.Println("note.Note:", value)
	}
}
