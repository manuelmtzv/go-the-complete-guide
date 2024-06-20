package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RequestAction(min, max int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		println("\nWelcome! Choose an action:")
		println("1. Deposit")
		println("2. Withdraw")
		println("3. Check Balance")
		println("4. Exit")
		print("Choose an action: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			println("Invalid action")
			continue
		}

		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if (err != nil) || value < min || value > max {
			println("Invalid action")
			continue
		}

		return value
	}

}
