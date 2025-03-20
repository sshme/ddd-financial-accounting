package entities

import (
	"errors"
	"github.com/google/uuid"
)

type BankAccount struct {
	ID      string
	Name    string
	Balance float64
}

func NewBankAccount(name string) (*BankAccount, error) {
	if name == "" {
		return nil, errors.New("bank account name cannot be empty")
	}

	return &BankAccount{
		ID:      generateUUID(),
		Name:    name,
		Balance: 0,
	}, nil
}

func (ba *BankAccount) SetBalance(balance float64) error {
	if balance < 0 {
		return errors.New("insufficient funds")
	}
	ba.Balance = balance
	return nil
}

func (ba *BankAccount) UpdateBalance(amount float64) error {
	if amount < 0 && ba.Balance+amount < 0 {
		return errors.New("insufficient funds")
	}
	ba.Balance += amount
	return nil
}

func (ba *BankAccount) UpdateName(name string) error {
	if name == "" {
		return errors.New("bank account name cannot be empty")
	}
	ba.Name = name
	return nil
}

func generateUUID() string {
	return uuid.NewString()
}
