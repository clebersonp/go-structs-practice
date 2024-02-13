package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Todo struct
// `json:"text"` are called 'struct tags' - Metadata and golang feature
type Todo struct {
	Text string `json:"text"`
}

// Display - method with receiver of note struct type
func (t Todo) Display() {
	fmt.Println(t.Text)
}

// Save - in Go we don't need to specify that some structs uses or implements an interface directly.
// We just need follow the same signature and Go automatically implement it
func (t Todo) Save() error {
	fileName := "todo.json"
	// Marshal needs to upper case all field struct that will be saved into json format
	jsonData, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

// New constructor function
func New(content string) (Todo, error) {
	// validation
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: content,
	}, nil
}
