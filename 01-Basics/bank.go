package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"go.com/basics/util" // module followed by the package/dir name
)

const BALANCE_FILE = "balance.txt"

func bank() {
	var accountBalance, err = util.ReadFloatFromFile(BALANCE_FILE)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("____________________")
		panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Phone No.: ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
			case 1:
				fmt.Println("Your balance is", accountBalance)
			case 2:
				fmt.Print("Your deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					// return
					continue
				}

				accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
				util.WriteFloatToFile(accountBalance, BALANCE_FILE)
				fmt.Println("Balance updated! New amount:", accountBalance)
			case 3 :
				fmt.Print("Withdrawal amount: ")
				var withdrawalAmount float64
				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					return
				}

				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					return
				}

				accountBalance -= withdrawalAmount // accountBalance = accountBalance + depositAmount
				util.WriteFloatToFile(accountBalance, BALANCE_FILE)
				fmt.Println("Balance updated! New amount:", accountBalance)
			default :  
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for choosing our bank")
				return
		}
	}

}