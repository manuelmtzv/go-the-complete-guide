package main

import (
	"fmt"
)

func main() {
	revenue := requestInputValue[float64]("Enter revenue: ")
	expenses := requestInputValue[float64]("Enter expenses: ")
	taxRate := requestInputValue[float64]("Enter tax rate: ")

	profit, ebt, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Profit is %.2f\n", profit)
	fmt.Printf("EBT is %.2f\n", ebt)
	fmt.Printf("Profit margin is %.2f", ratio)
}

func requestInputValue[T any](message string) T {
	var value T
	fmt.Print(message)
	fmt.Scan(&value)
	return value
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return profit, ebt, ratio
}
