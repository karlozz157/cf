package wallet

import (
	"errors"
	"github.com/karlozz157/cf/entities"
)

const (
	income string = "Depósito"
	expense string = "Retiro"
)

type Wallet struct {
	balance float32
	transactions []entities.Transaction
}

func (wallet *Wallet) registerTransaction (transaction entities.Transaction) {
	wallet.transactions = append(wallet.transactions, transaction)
}

func (wallet *Wallet) Deposit (transaction entities.Transaction) {
	wallet.balance += transaction.Amount

	transaction.Description = income
	transaction.TransactionType = entities.TransactionIncome

	wallet.registerTransaction(transaction)
}

func (wallet *Wallet) GetBalance () float32 {
	return wallet.balance
}

func (wallet *Wallet) GetTransactions () []entities.Transaction {
	return wallet.transactions
}

func (wallet *Wallet) Withdrawal (transaction entities.Transaction) error {
	if wallet.balance < transaction.Amount {
		return errors.New("cannot withdraw, insufficient funds")
	}

	wallet.balance -= transaction.Amount

	transaction.Description = expense
	transaction.TransactionType = entities.TransactionExpense

	wallet.registerTransaction(transaction)

	return nil
} 
