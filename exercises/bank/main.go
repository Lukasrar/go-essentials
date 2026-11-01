package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type bankAction int

const (
	checkBalance bankAction = iota + 1
	depositMoney
	withdrawMoney
	exit
)

const ACCOUNT_BALANCE_FILE_NAME = "balance.txt"

func writeBalanceToFile(value float64) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(ACCOUNT_BALANCE_FILE_NAME, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, errRead := os.ReadFile(ACCOUNT_BALANCE_FILE_NAME)

	if errRead != nil {
		return 1000, errors.New("Failed to read file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 1000, errors.New("Failed to store balance value")
	}
	return balance, nil
}

func main() {

	var accountBalance, errorFromFile = getBalanceFromFile()

	if errorFromFile != nil {
		fmt.Println("Error")
		fmt.Println(errorFromFile)
		fmt.Println("-----")
		panic("Cant continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What would you like to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case int(checkBalance):
			fmt.Printf("Your balance is %v \n", accountBalance)

		case int(depositMoney):
			fmt.Println("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Value must be greater than zero")
				continue
			}

			accountBalance += deposit
			fmt.Println("Updated balance: ", accountBalance)

			writeBalanceToFile(accountBalance)
		case int(withdrawMoney):
			fmt.Println("Amount: ")
			var amount float64
			fmt.Scan(&amount)

			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if amount > accountBalance {
				fmt.Println("Invalid amount. You cant withdraw more than you have")
				continue
			}

			accountBalance -= amount
			fmt.Println("Updated balance: ", accountBalance)
			writeBalanceToFile(accountBalance)

		default:
			fmt.Println("Session finished")
			return
		}
	}
}
