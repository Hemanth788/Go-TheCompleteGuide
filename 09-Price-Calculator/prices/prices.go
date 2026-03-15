package prices

import (
	"fmt"

	"go.com/price-calculator/conversion"
	fileManager "go.com/price-calculator/fileManager"
)

type STaxedPricesJob struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices map[string]string
}

const PRICES_FILE_PATH = "prices.txt"
const TAXED_PRICES_FILE_PATH = "taxed_prices.json"

func (tpJob *STaxedPricesJob) LoadData() {
	lines, err := fileManager.ReadLinesFromFile(PRICES_FILE_PATH)
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	tpJob.Prices = prices
}

func (tpJob *STaxedPricesJob) Process() {
	tpJob.LoadData()
	var result = make(map[string]string)

	for _, price := range tpJob.Prices {
		taxedPrice := price * (1 + tpJob.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	tpJob.TaxedPrices = result
	fileManager.WriteJSONToFile(tpJob, fmt.Sprintf("%.0f%v%v", tpJob.TaxRate * 100, "%_", TAXED_PRICES_FILE_PATH))
}

func NewSTaxedPricesJob(taxRate float64) *STaxedPricesJob {
	return &STaxedPricesJob{
		Prices: []float64{},
		TaxRate: taxRate,
	}
}
