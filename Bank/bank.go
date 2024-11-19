package main

import (
	"fmt"

	"example.com/bank/fileUtils"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileUtils.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("==============")
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("==============")
		panic("JUEPUTA SE CAYO EL RANCHO")
	}

	fmt.Println("Welcome to Go bank")
	for {
		presentOptions()

		var choice int

		print("Select an action: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is: ", accountBalance)
		case 2:
			var userDeposit float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&userDeposit)

			if userDeposit <= 0 {
				fmt.Println("Invalid deposit amount")
				continue
			}

			accountBalance += userDeposit

			fileUtils.WriteFloatToFile(accountBalance, accountBalanceFile)

			fmt.Println("New balance:", accountBalance)
		case 3:
			var withDrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withDrawAmount)

			if withDrawAmount <= 0 {
				fmt.Println("Invalid withdrawal amount")
				continue
			}

			if withDrawAmount > accountBalance {
				fmt.Println("The withdrawal ammount is greater than the balace account")
				continue
			}

			accountBalance -= withDrawAmount

			fileUtils.WriteFloatToFile(accountBalance, accountBalanceFile)

			fmt.Println("New balance:", accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}

	}
}
