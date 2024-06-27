package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

const pricesFilePath = "prices.txt"

func main() {
	taxRates := []float64{0.1, 0.2, 0.3}

	for _, taxRate := range taxRates {
		filemanager, err := filemanager.New(pricesFilePath, fmt.Sprintf("results/result_%.0f.json", taxRate*100))

		if err != nil {
			fmt.Printf("error while creating the filemanager")
			return
		}

		job := prices.NewTaxIncludedPriceJob(taxRate, filemanager)

		job.Process()
	}
}
