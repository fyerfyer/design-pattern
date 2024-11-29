package repository

import (
	"database/sql"
	"wallet_system/domain"
)

type TransactionRepository interface {
	Save(transaction domain.Transaction) error
}

type SQLTransactionRepository struct {
	db *sql.DB
}

func (t *SQLTransactionRepository) Save(transaction domain.Transaction) error {
	query := `INSERT INTO transactions (from_wallet_id, to_wallet_id, amount, type, create_time) 
	VALUES (?, ?, ?, ?, ?)`
	_, err := t.db.Exec(query,
		transaction.FromWalletID,
		transaction.ToWalletID,
		transaction.Amount,
		transaction.TransactionType,
		transaction.CreatedTime)
	return err
}
