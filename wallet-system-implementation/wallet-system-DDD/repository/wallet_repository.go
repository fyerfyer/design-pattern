package repository

import (
	"database/sql"
	"wallet_system/domain"
)

// The main role of the walletRepository interface is to
// decouple the domain layer from the data access layer.
// the domain layer is concerned only with how to operate
// through interfaces.

type WalletRepository interface {
	GetByID(walletID int64) (*domain.Wallet, error)
	Save(wallet *domain.Wallet) error
}

type SQLWalletRepository struct {
	db *sql.DB
}

func (w *SQLWalletRepository) GetByID(walletID int64) (*domain.Wallet, error) {
	var wallet domain.Wallet
	query := "SELECT id, created_time, balance FROM wallets WHERE id = ?"
	row := w.db.QueryRow(query, walletID)
	err := row.Scan(&wallet.ID,
		&wallet.CreatedTime,
		&wallet.Balance)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (w *SQLWalletRepository) Save(wallet *domain.Wallet) error {
	_, err := w.db.Exec("UPDATE wallets SET balance = ? WHERE id = ?", wallet.Balance.Amount, wallet.ID)
	return err
}
