package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	fmt.Println("Title:", title)
	fmt.Println("Content:", content)
}

func getNoteData() (string, string) {
	title, err := requestStringInput("Note title: ")

	if err != nil {
		panic(err)
	}

	content, err := requestStringInput("Note content: ")

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
