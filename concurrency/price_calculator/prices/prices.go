package prices

import (
	"errors"
	"fmt"
	"price-calculator/conversion"
	"price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	FileManager       iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"taxRate"`
	InputPrices       []float64           `json:"inputPrices"`
	TaxIncludedPrices map[string]string   `json:"taxIncludedPrices"`
}

func NewTaxIncludedPriceJob(taxRate float64, filemanager iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		FileManager: filemanager,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.FileManager.ReadLines()

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

func (job *TaxIncludedPriceJob) Process(channel chan bool, errorChannel chan error) {
	err := job.LoadData()

	if err != nil {
		errorChannel <- err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		resultPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", resultPrice)
	}

	job.TaxIncludedPrices = result

	err = job.FileManager.WriteJson(job)

	if err != nil {
		errorChannel <- err
	}

	channel <- true
}
