package model

import "time"

type Wallet struct {
	ID          int64
	CreatedTime time.Time
	Balance     float64
}

type TransactionType string

const (
	Debit    TransactionType = "DEBIT"
	Credit   TransactionType = "CREDIT"
	Transfer TransactionType = "TRANSFER"
)
