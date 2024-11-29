package service

import (
	"database/sql"
	"time"
	"wallet_system/domain"
	"wallet_system/repository"
)

type WalletService struct {
	WalletRepo      repository.WalletRepository
	TransactionRepo repository.TransactionRepository
	db              *sql.DB
}

func (s *WalletService) GetBalance(walletID int64) (float64, error) {
	wallet, err := s.WalletRepo.GetByID(walletID)
	if err != nil {
		return 0, err
	}
	return wallet.GetBalance(), nil
}

func (s *WalletService) Debit(walletID int64, amount float64) error {
	wallet, err := s.WalletRepo.GetByID(walletID)
	if err != nil {
		return err
	}

	money := domain.Money{Amount: amount}
	if err := wallet.Debit(money); err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if err := s.WalletRepo.Save(wallet); err != nil {
		tx.Rollback()
		return err
	}

	transaction := domain.Transaction{
		FromWalletID:    walletID,
		Amount:          money,
		TransactionType: domain.TransactionTypeDebit,
		CreatedTime:     time.Now(),
	}

	if err := s.TransactionRepo.Save(transaction); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *WalletService) Credit(walletID int64, amount float64) error {
	wallet, err := s.WalletRepo.GetByID(walletID)
	if err != nil {
		return err
	}

	money := domain.Money{Amount: amount}
	if err := wallet.Credit(money); err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if err := s.WalletRepo.Save(wallet); err != nil {
		tx.Rollback()
		return err
	}

	transaction := domain.Transaction{
		ToWalletID:      &walletID,
		Amount:          money,
		TransactionType: domain.TransactionTypeCredit,
		CreatedTime:     time.Now(),
	}

	if err := s.TransactionRepo.Save(transaction); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
