package main

import (
	"benjamin/tax-calculator/filemanager"
	"benjamin/tax-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("data/prices.txt", fmt.Sprintf("data/result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Could not process job [%v]", err)
		}
	}
}
