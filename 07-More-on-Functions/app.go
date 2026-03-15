package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, getTransformFn(2)) // --> passing function as an argument
	tripled := transformNumbers(&numbers, getTransformFn(3))

	quadrupled := transformNumbers(&numbers, createTransformer(4))

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(quadrupled)
}

type TTransformFn func(int) int

func transformNumbers(numbers *[]int, transformFn TTransformFn) []int { // --> taking functions as a parameter
	dNum := []int{}

	for _, val := range *numbers {
		dNum = append(dNum, transformFn(val))
	}

	return dNum;
}

func getTransformFn(num int) TTransformFn {
	if num == 2 {
		return func (num int) int {  // --> anonymous function
			return num * 2
		}
		// return doubleInt
	}
	return func (num int) int {  // --> anonymous function
		return num * 3
	}
	// return tripleInt
}

// func doubleInt(num int) int {
// 	return num * 2
// }

// func tripleInt(num int) int {
// 	return num * 3
// }

func createTransformer(factor int) TTransformFn { // --> factory function
	return func(number int) int {
		return number * factor; // --> closure, being able to use factor inside this anonymous function that's returned, factor –– the parameter of createTransformer func, is in the scope of this anon func
 	}
}
