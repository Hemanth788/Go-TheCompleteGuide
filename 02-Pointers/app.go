package main

import "fmt"

func main() {
	age := 43

	ageAddr := &age // addressOf age value is stored

	fmt.Println("Age before: ", age)
	// adultYears := getAdultYears(ageAddr)
	// fmt.Println("Adult age: ", adultYears)
	getAdultYears(ageAddr)
	fmt.Println("Age after: ", age)
}

// no arithmetic on pointers/addresses

func getAdultYears(ageAddr *int) {
	// return *ageAddr - 18 // valueAt the given address is used, de-referencing
	*ageAddr -= 18
}

