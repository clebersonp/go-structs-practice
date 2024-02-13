package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note struct
type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// Display - method with receiver of note struct type
func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.title, n.content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName += ".json"
	data := fmt.Sprintf("{ \"title\": \"%v\", \"content\": \"%v\", \"createdAt\": \"%v\" }", n.title, n.content, n.createdAt)
	return os.WriteFile(fileName, []byte(data), 0644)
}

// New constructor function
func New(title, content string) (Note, error) {
	// validation
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
