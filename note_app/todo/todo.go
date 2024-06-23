package todo

import (
	"errors"
	"fmt"
	"note-app/utility"
)

type Todo struct {
	Text string
	Done bool
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("text cannot be empty")
	}

	return &Todo{
		Text: text,
		Done: false,
	}, nil
}

func (todo *Todo) Display() string {
	return fmt.Sprintf("[%t] %s", todo.Done, todo.Text)
}

func (todo *Todo) Save() error {
	fileName := utility.ComposeFileName("todo", ".json")

	err := utility.SaveToJson(fileName, todo)

	if err != nil {
		return err
	}

	return nil
}
