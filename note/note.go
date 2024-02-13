package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note struct
// `json:"title"` are called 'struct tags' - Metadata and golang feature
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Display - method with receiver of note struct type
func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName += ".json"
	// Marshal needs to upper case all field struct that will be saved into json format
	jsonData, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}

// New constructor function
func New(title, content string) (Note, error) {
	// validation
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
