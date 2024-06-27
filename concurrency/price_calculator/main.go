package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

const pricesFilePath = "prices.txt"

func main() {
	taxRates := []float64{0.1, 0.2, 0.3}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)
		filemanager, err := filemanager.New(pricesFilePath, fmt.Sprintf("results/result_%.0f.json", taxRate*100))

		if err != nil {
			fmt.Printf("error while creating the filemanager")
			return
		}

		job := prices.NewTaxIncludedPriceJob(taxRate, filemanager)

		go job.Process(doneChannels[index], errorChannels[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("Done.")
		}
	}

}
