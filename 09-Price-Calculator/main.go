package main

import (
	"fmt"

	cmdManager "go.com/price-calculator/cmdManager"
	"go.com/price-calculator/prices"
)

const PRICES_FILE_PATH = "prices.txt"
const TAXED_PRICES_FILE_PATH = "taxed_prices.json"

func main() {
	var taxRates = []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := fileManager.New(PRICES_FILE_PATH, fmt.Sprintf("%.0f%v%v", taxRate * 100, "_%", TAXED_PRICES_FILE_PATH))
		cmdM := cmdManager.New()
		priceJob := prices.NewSTaxedPricesJob(cmdM, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Failed to process the job")
			fmt.Println(err)
		}
	}
}
