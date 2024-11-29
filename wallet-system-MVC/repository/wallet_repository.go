package repository

import (
	"database/sql"
	"wallet-system/model"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetWalletEntity(walletID int64) (*model.Wallet, error) {
	var wallet model.Wallet
	query := "SELECT id, created_time, balance FROM wallets WHERE id = ?"
	row := r.db.QueryRow(query, walletID)
	err := row.Scan(&wallet.ID,
		&wallet.CreatedTime,
		&wallet.Balance)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) UpdateBalance(walletID int64, balance float64) error {
	query := "UPDATE wallets SET balance = ? WHERE id = ?"
	_, err := r.db.Exec(query, balance, walletID)
	return err
}
