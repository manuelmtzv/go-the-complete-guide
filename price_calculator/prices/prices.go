package prices

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedTaxes map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {
	file, err := os.Open("prices.txt")

	if err != nil {
		return errors.New("error opening file")
	}

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return errors.New("error reading file")
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		price, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return errors.New("error parsing price")
		}

		prices[lineIndex] = price
	}

	job.InputPrices = prices

	defer file.Close()
	return nil
}

func (job *TaxIncludedPriceJob) Process() {
	err := job.LoadData()

	if err != nil {
		fmt.Println(err)
		return
	}

	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}
