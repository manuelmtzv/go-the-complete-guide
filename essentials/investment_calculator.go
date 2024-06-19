package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount := requestInputValue[float64]("Enter investment amount: ")
	returnRate := requestInputValue[float64](fmt.Sprintf("Enter the expected return rate (default: %.2f): ", inflationRate))
	years := requestInputValue[float64]("Enter the investment time (years): ")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, returnRate, years)

	printResults(futureValue, futureRealValue)
}

func requestInputValue[T any](message string) T {
	var value T
	fmt.Print(message)
	fmt.Scan(&value)
	return value
}

func calculateFutureValues(investmentAmount, returnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow((1+returnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	return futureValue, futureRealValue
}

func printResults(futureValue, futureRealValue float64) {
	formattedFutureValue := fmt.Sprintf("Future value is %.2f", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future real value is %.2f", futureRealValue)

	fmt.Println(formattedFutureValue)
	fmt.Println(formattedFutureRealValue)
}
