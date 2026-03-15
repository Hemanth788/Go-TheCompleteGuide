package main

import (
	"fmt"

	"go.com/structs/user"
)

type S string

func (customStr S) log() {
	fmt.Println(customStr)
}

func main() {
	firstNameInput := getUserData("Please enter your first name: ")
	lastNameInput := getUserData("Please enter your last name: ")
	birthDateInput := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	var appUser *user.User

	var name S
	name = "Moong dal"
	name.log()
	
	// struct literal
	// can omit keys if order is same as the type def
	// can create an empty struct literal User{}
	
	appUser, err := user.New(firstNameInput, lastNameInput, birthDateInput)

	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}