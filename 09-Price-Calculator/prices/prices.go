package prices

import (
	"fmt"

	"go.com/price-calculator/conversion"
	ioManager "go.com/price-calculator/ioManager"
)

type STaxedPricesJob struct {
	TaxRate     float64 `json:"tax_rate"`
	Prices      []float64 `json:"prices"`
	TaxedPrices map[string]string `json:"taxed_prices"`
	IOManager ioManager.IIOManager `json:"-"`
}

func (tpJob *STaxedPricesJob) LoadData() error {
	lines, err := tpJob.IOManager.ReadResultFromInput()
	if err != nil {
		return err
	}
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	tpJob.Prices = prices
	return nil
}

func (tpJob *STaxedPricesJob) Process(doneChan chan bool, errorChan chan error) {
// error 
	err := tpJob.LoadData()
	if err != nil {
		errorChan <- err
		return
	}

	var result = make(map[string]string)

	for _, price := range tpJob.Prices {
		taxedPrice := price * (1 + tpJob.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	tpJob.TaxedPrices = result
	// return 
	tpJob.IOManager.WriteResultToOutput(tpJob)
	doneChan <- true
}

func NewSTaxedPricesJob(ioM ioManager.IIOManager, taxRate float64) *STaxedPricesJob {
	return &STaxedPricesJob{
		IOManager: ioM,
		Prices: []float64{},
		TaxRate: taxRate,
	}
}
