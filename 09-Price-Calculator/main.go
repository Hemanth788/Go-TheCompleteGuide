package main

import (
	"fmt"

	// cmdmanager "go.com/price-calculator/cmdManager"
	filemanager "go.com/price-calculator/fileManager"
	"go.com/price-calculator/prices"
)

const PRICES_FILE_PATH = "prices.txt"
const TAXED_PRICES_FILE_PATH = "taxed_prices.json"

func main() {
	var taxRates = []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		errorChans[idx] = make(chan error)
		fm := filemanager.New(PRICES_FILE_PATH, fmt.Sprintf("%.0f%v%v", taxRate * 100, "_%", TAXED_PRICES_FILE_PATH))
		// cmdM := cmdmanager.New()
		priceJob := prices.NewSTaxedPricesJob(fm, taxRate)
		go priceJob.Process(doneChans[idx], errorChans[idx])
		// if err != nil {
		// 	fmt.Println("Failed to process the job")
		// 	fmt.Println(err)
		// }
	}

	for idx, taxRate := range taxRates {
		select {
			case err := <- errorChans[idx]:
				if err != nil{
					fmt.Println(err)
				}
			case <- doneChans[idx]:
				fmt.Println("Done processing: ", taxRate)
			// for one use case, that's taxRate here, we have 2 channels, hence these are interconnected
			// so with select, we wait for one case of one channel to have something in it for its use case and the other case of the other channel isn't waited for
		}
	}
}
