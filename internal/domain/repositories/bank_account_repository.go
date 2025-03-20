package repositories

import (
	"ddd-financial-accounting/internal/domain/entities"
)

type BankAccountRepository interface {
	Save(account *entities.BankAccount) error
	GetByID(id string) (*entities.BankAccount, error)
	Update(account *entities.BankAccount) error
	Delete(id string) error
	GetAll() ([]*entities.BankAccount, error)
}
