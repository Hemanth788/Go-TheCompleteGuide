package main

import "fmt"

type Product struct {
	id    string
	name  string
	price float64
}

func main() {
	// 1
	arr := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(arr)

	// 2
	fmt.Println(arr[0])
	fmt.Println(arr[1:])

	// 3
	slice := arr[:2]
	fmt.Println(slice)

	// 4
	fmt.Println(cap(slice), slice[1:3]) // explicitly tell the :end if "re-slicing" a slice

	// 5
	goals := []string{"Learn", "Go"}
	fmt.Println(goals)

	// 6
	goals[1] = "Java"
	goals = append(goals, "as well")

	// 7
	products := []Product{
		{
			"1",
			"Phone1",
			0.01,
		},
		{
			"2",
			"Phone2",
			0.02,
		},
		{
			"3",
			"Phone3",
			0.03,
		},
	}

	newProduct := Product{
		"4",
		"Phone4",
		0.04,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
