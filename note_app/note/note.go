package note

import (
	"errors"
	"fmt"
	"note-app/utility"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Display() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", note.Title, note.Content, note.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (note *Note) Save() error {
	fileName := utility.ComposeFileName(note.Title, ".json")

	err := utility.SaveToJson(fileName, note)

	if err != nil {
		return err
	}

	return nil
}
