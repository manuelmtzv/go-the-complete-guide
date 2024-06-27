package prices

import (
	"errors"
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
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
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		return errors.New("error reading file")
	}

	prices, err := conversion.StringsToFloat(lines...)

	if err != nil {
		return errors.New("error converting strings to floats")
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process() {
	err := job.LoadData()

	if err != nil {
		fmt.Println(err)
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		resultPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", resultPrice)
	}
	fmt.Println(result)
}
