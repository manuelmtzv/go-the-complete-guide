package main

import (
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate)
		job.Process()
	}
}
