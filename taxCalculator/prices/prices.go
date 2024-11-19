package prices

import (
	"fmt"

	"example.com/tax-calculator/file"
	"example.com/tax-calculator/parse_utils"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices TaxIncludedPrices `json:"tax_included_prices"`
	IOManager         file.FileManager  `json:"-"`
}

type TaxIncludedPrices map[string]string

func NewTaxIncludedPriceJob(fm file.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		IOManager:   fm,
	}
}

func (j *TaxIncludedPriceJob) Process() {
	j.LoadData()
	result := make(TaxIncludedPrices)

	for _, price := range j.InputPrices {
		taxIncludedPrice := price * (1 + j.TaxRate)
		result[parse_utils.TransformFloatToString(price)] = parse_utils.TransformFloatToString(taxIncludedPrice)
	}

	j.TaxIncludedPrices = result
	j.IOManager.WriteJSON(j)
}

func (j *TaxIncludedPriceJob) LoadData() {

	pricesSlice, err := j.IOManager.ReadDataFile()

	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
		return
	}
	j.InputPrices = pricesSlice
}
