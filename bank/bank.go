package main

import (
	"fmt"
)

type Account struct {
	balance float64
}

func main() {
	savedBalance, err := GetFloatFromFile("balance.txt")
	if err != nil {
		panic(err)
	}

	account := Account{balance: savedBalance}

	for {
		action := RequestAction(1, 4)

		switch action {
		case 1:
			account.Deposit()
		case 2:
			account.Withdraw()
		case 3:
			account.CheckBalance()
		case 4:
			fmt.Println("Thank you for using our service")
			return
		default:
			fmt.Println("Invalid action")
			return
		}

		WriteFloatToFile("balance.txt", float64(account.balance))
	}
}
