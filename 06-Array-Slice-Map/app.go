package main

import "fmt"

type SSMap map[string]string

func (ssm SSMap) log() {
	fmt.Println(ssm)
}

func main() {
	prices := []float64{10.99, 8.99}
	prices = append(prices, 5.99)

	pricesToAppend := []float64{1.99, 2.99}
	prices = append(prices, pricesToAppend...) // spread operator from JS, but at the end, not beginning, looks bad

	fmt.Println(prices)

	websites := map[string]string{
		"google": "google.com",
		"amazon": "amazon.com",
	}
	fmt.Println(websites["google"])

	userNames := make([]string, 2, 5)
	//   [[size]], min, max	
	
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Maxmillian")
	// _ _ Max Maxmillian

	courses := make(map[string]string, 5)
	fmt.Println(courses)

	ssm := SSMap{"bla": "bli", "blu": "!"}
	ssm.log()

	// looping through arrays, slices
	for idx, val := range userNames {
		fmt.Println(idx, ": ", val)
	}
	
	fmt.Println()
	
	// looping through maps
	for key, val := range ssm {
		fmt.Println(key, ": ", val)
	}
}

// func main() {
// 	prices := [4]float64{1.31, 2.14, 3.11, 4.99}
// 	end := 4
// 	featurePrices := prices[1:end]
// 	featurePrices[0] = 12.22 // changes the original array
// 	fmt.Println(prices, featurePrices,
// 		len(featurePrices), // number of elements in an array
// 		cap(featurePrices)) // what all number of elements can be chosen towards the end of the original array
// 		/*
// 		i 0 1 2 3
// 		x 1 2 3 4
// 		x [1:3]
// 		y 2 4 len: 2, cap: 3
// 		y [:4]
// 		z 2 3 4 len: 3, cap: 3 to its max. capacity
// 		slice is a window into the original array
// 		*/
// }
