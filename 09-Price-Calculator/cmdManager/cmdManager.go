package cmdmanager

import "fmt"

type CMDManager struct {}

func (cmdM CMDManager) ReadResultFromInput() ([]string, error) {
	fmt.Println("Please enter your prices, Enter after each")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "-" {
			break
		}

		prices = append(prices, price)

	}
	return prices, nil
}

func (cmdM CMDManager) WriteResultToOutput(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}