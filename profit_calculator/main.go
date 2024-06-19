package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(message)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value <= 0 {
			fmt.Println("Invalid value. Value must be a positive number.")
			continue
		}

		return value
	}
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return profit, ebt, ratio
}

func writeResultsInFile(profit, ebt, ratio float64) {
	resultString := fmt.Sprintf(`
		Profit: %.2f
		EBT: %.2f
		Profit margin: %.2f
	`, profit, ebt, ratio)

	err := os.WriteFile("results.txt", []byte(resultString), 0644)
	if err != nil {
		panic(err)
	}
}
