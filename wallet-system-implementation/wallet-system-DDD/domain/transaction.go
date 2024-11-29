package domain

import "time"

type TransactionType string

const (
	TransactionTypeDebit    TransactionType = "DEBIT"
	TransactionTypeCredit   TransactionType = "CREDIT"
	TransactionTypeTransfer TransactionType = "TRANSFER"
)

type Transaction struct {
	ID              int64
	FromWalletID    int64
	ToWalletID      *int64
	Amount          Money
	TransactionType TransactionType
	CreatedTime     time.Time
}
