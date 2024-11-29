package repository

import (
	"database/sql"
	"wallet-system/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) SaveTransaction(transaction *model.Transaction) error {
	query := `INSERT INTO transactions (from_wallet_id, to_wallet_id, amount, type, create_time) 
              VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query,
		transaction.FromWalletID,
		transaction.ToWalletID,
		transaction.Amount,
		transaction.TransactionType,
		transaction.CreatedTime)
	return err
}
