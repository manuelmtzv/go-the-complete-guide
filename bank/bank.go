package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Account struct {
	balance float64
}

func main() {
	savedBalance, err := readBalanceFromFile()
	if err != nil {
		panic(err)
	}

	account := Account{balance: savedBalance}

	for {
		action := requestAction()

		switch action {
		case 1:
			account.deposit()
		case 2:
			account.withdraw()
		case 3:
			account.checkBalance()
		case 4:
			fmt.Println("Thank you for using our service")
			return
		default:
			fmt.Println("Invalid action")
			return
		}

		writeBalanceToFile(float64(account.balance))
	}
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("failed to read balance")
	}
	parsedBalance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to parse balance")
	}
	return parsedBalance, nil
}

func writeBalanceToFile(balance float64) {
	os.WriteFile("balance.txt", []byte(fmt.Sprintf("%f", balance)), 0644)
}

func requestAction() int {
	var action int
	fmt.Println("\n1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check balance")
	fmt.Println("4. Exit")

	fmt.Print("Enter action: ")
	fmt.Scan(&action)

	return action
}

func (account *Account) deposit() {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("Invalid amount")
		return
	}
	account.balance += amount
	account.checkBalance()
}

func (account *Account) withdraw() {
	var amount string
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)

	parsedAmount, err := strconv.ParseFloat(amount, 64)

	if err != nil || parsedAmount <= 0 {
		fmt.Println("Invalid amount")
		return
	}

	if parsedAmount > account.balance {
		fmt.Println("Insufficient balance")
		return
	}
	account.balance -= parsedAmount
	account.checkBalance()
}

func (account *Account) checkBalance() {
	fmt.Println("Balance: ", account.balance)
}
