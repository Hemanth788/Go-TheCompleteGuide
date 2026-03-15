package main

import "fmt"

func profitCalc() (float64, float64, float64) {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue, expenses and tax rate, space-separated: ")
	fmt.Scan(&revenue, &expenses, &taxRate);

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
