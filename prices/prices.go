package prices

import (
	"benjamin/tax-calculator/conversion"
	"benjamin/tax-calculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (j *TaxIncludedPriceJob) LoadData() error {
	lines, err := j.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	j.InputPrices = prices
	return nil
}

func (j *TaxIncludedPriceJob) Process() error {
	err := j.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range j.InputPrices {
		taxIncludedPrice := price * (1 + j.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	j.TaxIncludedPrices = result
	return j.IOManager.WriteResult(j)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
