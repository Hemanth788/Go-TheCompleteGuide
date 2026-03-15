package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, getTransformFn(2)) // --> passing function as an argument
	tripled := transformNumbers(&numbers, getTransformFn(3))

	quadrupled := transformNumbers(&numbers, createTransformer(4))
	fiveFactorial := factorial(5)

	sum3 := sumUp(1, 2, 3)
	sum4 := sumUp(1, 2, 3, 4)
	sum5 := sumUp(numbers[0], numbers[1:]...)

	fmt.Println("numbers: ", numbers)
	fmt.Println("doubled: ", doubled)
	fmt.Println("tripled: ", tripled)
	fmt.Println("quadrupled: ", quadrupled)
	fmt.Println("fiveFactorial: ", fiveFactorial)
	fmt.Println("sum3: ", sum3)
	fmt.Println("sum4: ", sum4)
	fmt.Println("sum5: ", sum5)
}

type TTransformFn func(int) int

func transformNumbers(numbers *[]int, transformFn TTransformFn) []int { // --> taking functions as a parameter
	dNum := []int{}

	for _, val := range *numbers {
		dNum = append(dNum, transformFn(val))
	}

	return dNum
}

func getTransformFn(num int) TTransformFn {
	if num == 2 {
		return func(num int) int { // --> anonymous function
			return num * 2
		}
		// return doubleInt
	}
	return func(num int) int { // --> anonymous function
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
		return number * factor // --> closure, being able to use factor inside this anonymous function that's returned, factor –– the parameter of createTransformer func, is in the scope of this anon func
	}
}

func factorial(number int) int {
	if number == 2 {
		return number
	}
	return number * factorial(number-1)
}

func sumUp(start int, numbers ...int) int { // --> variadic function
	sum := start

	for _, val := range numbers {
		sum += val
	}

	return sum
}
