package entities

type TransactionType string

const (
	TransactionIncome TransactionType = "Income"
	TransactionExpense TransactionType = "Expense"
)

type Transaction struct {
	Amount float32
	Description string
	TransactionType TransactionType 
}
