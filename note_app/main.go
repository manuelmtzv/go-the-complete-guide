package main

import (
	"bufio"
	"errors"
	"fmt"
	"note-app/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note created successfully!")
	fmt.Println(note.Display())

	err = note.Save()

	if err != nil {
		fmt.Println(err)
		return
	}
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
