package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	returnRate := 5.5
	years := 10.0

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+returnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Printf("Future value is %f\n", futureValue)
	fmt.Printf("Future real value is %f", futureRealValue)
}
