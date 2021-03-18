package main

import (
	"fmt"
	"github.com/karlozz157/cf/entities"
	"github.com/karlozz157/cf/wallet"
	"os"
)

func getAmount(message string) float32 {
	var amount float32
	fmt.Print(message)
	fmt.Scanf("%f", &amount)
	return amount
}

func createTransaction (myWallet *wallet.Wallet, amount float32, isIncome bool) {
	transaction := entities.Transaction { Amount: amount }

	if isIncome {
		myWallet.Deposit(transaction)
	} else {
		err := myWallet.Withdrawal(transaction)

		if err != nil {
			fmt.Print(err)
		}
	}

	fmt.Printf("Your current balance is %f", myWallet.GetBalance())
}

func getTransactions (myWallet *wallet.Wallet) {
	transactions := myWallet.GetTransactions()
	isEmpty := len(transactions) == 0

	fmt.Println(transactions)

	if isEmpty {
		fmt.Println("You dont have transactions")
		return
	}

	for _, transaction := range transactions {
		fmt.Println("======================")
		fmt.Println("Amount:      ", transaction.Amount)
		fmt.Println("Description: ", transaction.Description)
	}

	fmt.Printf("\nYour current balance is %f", myWallet.GetBalance())
}

func program (myWallet *wallet.Wallet) {
	var option int

	fmt.Println("\n\nMenu")
	fmt.Println("1) Make a deposit")
	fmt.Println("2) Withdrawal")
	fmt.Println("3) My transactions")
	fmt.Println("4) Exit")
	fmt.Print("\nChoose an option: ")
	fmt.Scanf("%d", &option)
	fmt.Println("")

	switch option {
	case 1:
		amount := getAmount("How much do you want to deposit? ")
		createTransaction(myWallet, amount, true)
	case 2:
		amount := getAmount("How much do you want to withdraw? ")
		createTransaction(myWallet, amount, false)
	case 3:
		getTransactions(myWallet)
	case 4:
		os.Exit(1)
	default:
		fmt.Println("You have to choose a valid option")
	}
}

func main () {
	myWallet := wallet.Wallet {}

	for true {
		program(&myWallet)
	}
}
