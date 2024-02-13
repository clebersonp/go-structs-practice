package note

import (
	"errors"
	"time"
)

// Note struct
type Note struct {
	title     string
	content   string
	createdAt time.Time
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
