package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs/user"
)

func main() {
	user, error := user.New(requestUserData())

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(user.ToString())
	user.Clear()
	fmt.Println(user.ToString())
}

func requestUserData() (string, string, string) {
	firstName, _ := requestStringInput("Enter your first name: ")
	lastName, _ := requestStringInput("Enter your last name: ")
	birthDate, _ := requestStringInput("Enter your birth date(MM/DD/YY): ")
	return firstName, lastName, birthDate
}

func requestStringInput(promt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(promt)
	data, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(data), nil
}
