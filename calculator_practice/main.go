package main

import (
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 01, 0.15}
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Calculate()
		priceJob.LoadData()
	}

}
