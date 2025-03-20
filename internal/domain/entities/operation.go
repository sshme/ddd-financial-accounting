package entities

import (
	"errors"
	"time"
)

type Operation struct {
	ID            string
	Type          GroupByEnum
	BankAccountId string
	Amount        float64
	Date          time.Time
	Description   string
	CategoryId    string
}

func NewOperation(operationType GroupByEnum, bankAccountId string, amount float64, description string, categoryId string) (*Operation, error) {
	if amount < 0 {
		return nil, errors.New("operation amount must be greater or equal than zero")
	}

	return &Operation{
		ID:            generateUUID(),
		Type:          operationType,
		BankAccountId: bankAccountId,
		Amount:        amount,
		Date:          time.Now(),
		Description:   description,
		CategoryId:    categoryId,
	}, nil
}
