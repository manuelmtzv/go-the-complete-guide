package main

import (
	"essentials/bank/utility"
	"fmt"
)

const balanceFile = "balance.txt"

type Account struct {
	balance float64
}

func main() {
	savedBalance, err := utility.GetFloatFromFile(balanceFile)
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

		utility.WriteFloatToFile("balance.txt", float64(account.balance))
	}
}
