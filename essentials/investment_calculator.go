package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	investmentAmount := requestInputValue[float64]("Enter investment amount: ")
	returnRate := requestInputValue[float64]("Enter the expected return rate (default: %.2f): ")
	years := requestInputValue[float64]("Enter the investment time (years): ")

	fmt.Printf("Enter the expected return rate (default: %.2f): ", returnRate)
	fmt.Scan(&returnRate)

	fmt.Print("Enter the investment time (years): ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+returnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	formattedFutureValue := fmt.Sprintf("Future value is %.2f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Future real value is %.2f\n", futureRealValue)

	fmt.Println(formattedFutureValue, formattedFutureRealValue)
}

func requestInputValue[T any](message string) T {
	var value T
	fmt.Print(message)
	fmt.Scan(&value)
	return value
}
