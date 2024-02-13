package main

import (
	"errors"
	"fmt"
)

func main() {
	// ask user for the data
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	// store in a note structure
	// store the note into a file

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note content:")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)

	if userInput == "" {
		return "", errors.New("invalid input")
	}

	return userInput, nil
}
