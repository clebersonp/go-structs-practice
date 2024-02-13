package note

import (
	"errors"
	"fmt"
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
	fmt.Printf("Your n titled %v has the following content:\n%v\n", n.title, n.content)
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
