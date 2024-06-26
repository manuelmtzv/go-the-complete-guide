package main

import "fmt"

func main() {
	prices := []float64{100, 200, 300}
	taxRates := []float64{0.1, 0.2, 0.3}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for i, price := range prices {
			taxIncludedPrices[i] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}
