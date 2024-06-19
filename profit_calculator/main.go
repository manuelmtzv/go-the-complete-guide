package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	revenue := requestFloatValue("Enter revenue: ")
	expenses := requestFloatValue("Enter expenses: ")
	taxRate := requestFloatValue("Enter tax rate: ")

	profit, ebt, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Profit is %.2f\n", profit)
	fmt.Printf("EBT is %.2f\n", ebt)
	fmt.Printf("Profit margin is %.2f", ratio)

	writeResultsInFile(profit, ebt, ratio)
}

func requestFloatValue(message string) float64 {
	var value string
	fmt.Print(message)
	fmt.Scan(&value)

	parsedValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		fmt.Println("Invalid value")
		return requestFloatValue(message)
	}

	return parsedValue
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return profit, ebt, ratio
}

func writeResultsInFile(profit, ebt, ratio float64) {
	file, err := os.Create("results.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf(`
		Profit: %.2f
		EBT: %.2f
		Profit margin: %.2f
	`, profit, ebt, ratio))
}
