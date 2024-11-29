package domain

import (
	"errors"
	"time"
)

type Money struct {
	Amount float64
}

func (m Money) Add(other Money) Money {
	return Money{Amount: m.Amount + other.Amount}
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.Amount < other.Amount {
		return Money{}, errors.New("insufficient funds")
	}
	return Money{Amount: m.Amount - other.Amount}, nil
}

type Wallet struct {
	ID          int64
	Balance     Money
	CreatedTime time.Time
}

func NewVirtualWallet(id int64) *Wallet {
	return &Wallet{
		ID:          id,
		Balance:     Money{Amount: 0},
		CreatedTime: time.Now(),
	}
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance.Amount
}

func (w *Wallet) Credit(amount Money) error {
	if amount.Amount <= 0 {
		return errors.New("invalid credit amount")
	}
	w.Balance = w.Balance.Add(amount)
	return nil
}

func (w *Wallet) Debit(amount Money) error {
	newBalance, err := w.Balance.Subtract(amount)
	if err != nil {
		return err
	}
	w.Balance = newBalance
	return nil
}
