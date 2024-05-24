package prices

import (
	"example.com/calculator/conversion"
	"example.com/calculator/filemanager"
	"fmt"
	"log"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (t *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	prices, err := conversion.StringToFloats(lines)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}
	t.InputPrices = prices
}
func (t *TaxIncludedPriceJob) Calculate() {
	t.LoadData()
	t.TaxIncludedPrices = make(map[string]string)
	for _, price := range t.InputPrices {
		t.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%2.f", price*(1+t.TaxRate))
	}
	err := filemanager.WriteJson(fmt.Sprintf("tax_included_prices_%v.json", t.TaxRate), t)
	if err != nil {
		log.Fatal("Error: ", err)

	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{100, 200, 300},
		TaxRate:     taxRate,
	}
}
