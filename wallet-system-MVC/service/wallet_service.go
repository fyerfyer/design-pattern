package service

import (
	"database/sql"
	"errors"
	"time"
	"wallet-system/model"
	"wallet-system/repository"
)

type VirtualWalletService struct {
	walletRepo      *repository.WalletRepository
	transactionRepo *repository.TransactionRepository
	db              *sql.DB
}

func NewVirtualWalletService(walletRepo *repository.WalletRepository, transactionRepo *repository.TransactionRepository, db *sql.DB) *VirtualWalletService {
	return &VirtualWalletService{walletRepo: walletRepo, transactionRepo: transactionRepo, db: db}
}

func (s *VirtualWalletService) GetBalance(walletID int64) (float64, error) {
	wallet, err := s.walletRepo.GetWalletEntity(walletID)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}

func (s *VirtualWalletService) Debit(walletID int64, amount float64) error {
	wallet, err := s.walletRepo.GetWalletEntity(walletID)
	if err != nil {
		return err
	}

	if wallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	transaction := &model.Transaction{
		FromWalletID:    walletID,
		Amount:          amount,
		TransactionType: model.Debit,
		CreatedTime:     time.Now(),
	}
	if err := s.transactionRepo.SaveTransaction(transaction); err != nil {
		tx.Rollback()
		return err
	}

	newBalance := wallet.Balance - amount
	if err := s.walletRepo.UpdateBalance(walletID, newBalance); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *VirtualWalletService) Credit(walletID int64, amount float64) error {
	wallet, err := s.walletRepo.GetWalletEntity(walletID)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	transaction := &model.Transaction{
		FromWalletID:    walletID,
		Amount:          amount,
		TransactionType: model.Credit,
		CreatedTime:     time.Now(),
	}
	if err := s.transactionRepo.SaveTransaction(transaction); err != nil {
		tx.Rollback()
		return err
	}

	newBalance := wallet.Balance + amount
	if err := s.walletRepo.UpdateBalance(walletID, newBalance); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *VirtualWalletService) Transfer(fromWalletID, toWalletID int64, amount float64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if err := s.Debit(fromWalletID, amount); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.Credit(toWalletID, amount); err != nil {
		tx.Rollback()
		return err
	}

	transaction := &model.Transaction{
		FromWalletID:    fromWalletID,
		ToWalletID:      toWalletID,
		Amount:          amount,
		TransactionType: model.Transfer,
		CreatedTime:     time.Now(),
	}
	if err := s.transactionRepo.SaveTransaction(transaction); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
