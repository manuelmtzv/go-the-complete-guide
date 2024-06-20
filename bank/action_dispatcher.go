package main

import (
	"fmt"
	"strconv"
)

func (account *Account) Deposit() {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("Invalid amount")
		return
	}
	account.balance += amount
	account.CheckBalance()
}

func (account *Account) Withdraw() {
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
	account.CheckBalance()
}

func (account *Account) CheckBalance() {
	fmt.Println("Balance: ", account.balance)
}
