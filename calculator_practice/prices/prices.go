package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (t *TaxIncludedPriceJob) Calculate() {
	t.TaxIncludedPrices = make(map[string]float64)
	for _, price := range t.InputPrices {
		t.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + t.TaxRate)
	}
	fmt.Println(t.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{100, 200, 300},
		TaxRate:     taxRate,
	}
}
