package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
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
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json, 0644)

	return nil
}
