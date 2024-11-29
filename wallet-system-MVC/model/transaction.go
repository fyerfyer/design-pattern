package model

import "time"

type Transaction struct {
	ID              int64
	FromWalletID    int64
	ToWalletID      int64
	Amount          float64
	TransactionType TransactionType
	CreatedTime     time.Time
}
