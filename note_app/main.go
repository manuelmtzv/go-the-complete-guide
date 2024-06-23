package main

import (
	"bufio"
	"errors"
	"fmt"
	"note-app/note"
	"note-app/todo"
	"note-app/utility"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	// Create a new note

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = utility.OutputData(note)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new todo

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = utility.OutputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getTodoData() string {
	fmt.Println("Enter todo details.")
	text, err := requestStringInput("Text: ")

	if err != nil {
		panic(err)
	}

	return text
}

func getNoteData() (string, string) {
	fmt.Println("Enter note details.")
	title, err := requestStringInput("Title: ")

	if err != nil {
		panic(err)
	}

	content, err := requestStringInput("Content: ")

	if err != nil {
		panic(err)
	}

	return title, content
}

func requestStringInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	trimmedInput := strings.TrimSpace(input)
	if len([]rune(trimmedInput)) == 0 {
		return "", errors.New("input cannot be empty")
	}

	return trimmedInput, nil
}
