package main

import "fmt"

func main() {
	prices := []float64{5, 10, 20}

	discountPrices := []float64{10.5, 20.0, 30.5}

	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}
