package main

import (
	"fmt"
	"math"
)

func investmentCalc() (string, string) {
	// Explicit type assignment
	// only declaration, no assignment, go assigns default "null" value which for float is 0.0
	var investmentAmount float64
	var years float64
	const inflationRate = 2.5

	// Declaring multiple variables and assigning them values
	// investmentAmount, years := 1000.0, 10.0
	// var investmentAmount, years = 1000.0, 10.0

	var expectedReturnRate = 5.5

	// Recommended to use assignment operator if you want type to be inferred
	// expectedReturnRate := 5.5

	fmt.Print("Enter investment amount and the number of years: ")
	fmt.Scan(&investmentAmount, &years)
	fmt.Println(investmentAmount)
	fmt.Println(years)
	

	var futureValue = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	var futureRealValue = futureValue / math.Pow((1 + inflationRate / 100), years)

	// Type conversion
	// var investmentAmount = 1000
	// var expectedReturnRate = 5.5
	// var years = 10

	// var futureValue = float(investmentAmount) * math.Pow((1 + expectedReturnRate / 100), float(years))
	fmt.Println("Future value: ", futureValue)
	fmt.Println("Inflation-adjusted future value: ", futureRealValue)
	
	// fmt.Printf("Future value: %v\nInflation-adjusted future value: %v", futureValue, futureRealValue)
	fmt.Printf("\nFuture value: %9.2f\nInflation-adjusted future value: %9.2f\n", futureValue, futureRealValue)
	
	// Multiline strings
	fmt.Printf(`
Future value: %9.2f
Inflation-adjusted future value: %9.2f

`,futureValue, futureRealValue)
	
	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Inflation-adjusted future value: %.2f\n", futureRealValue)
	return formattedFV, formattedRFV
}