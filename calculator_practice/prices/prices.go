package prices

import (
	"bufio"
	"example.com/calculator/conversion"
	"fmt"
	"log"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (t *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal("Error: ", scanner.Err())
	}

	prices, err := conversion.StringToFloats(lines)
	if err != nil {
		log.Fatal("Error: ", err)
		file.Close()
		return
	}
	t.InputPrices = prices
	file.Close()

}
func (t *TaxIncludedPriceJob) Calculate() {
	t.LoadData()
	t.TaxIncludedPrices = make(map[string]string)
	for _, price := range t.InputPrices {
		t.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%2.f", price*(1+t.TaxRate))
	}
	fmt.Println(t.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{100, 200, 300},
		TaxRate:     taxRate,
	}
}
