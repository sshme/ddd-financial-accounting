package services

import (
	"ddd-financial-accounting/internal/domain/entities"
	"ddd-financial-accounting/internal/domain/factories"
	"ddd-financial-accounting/internal/domain/repositories"
	"log/slog"
)

type AccountFacade struct {
	repository repositories.BankAccountRepository
	factory    *factories.BankAccountFactory
	logger     *slog.Logger
}

func NewAccountFacade(
	repository repositories.BankAccountRepository,
	factory *factories.BankAccountFactory,
	logger *slog.Logger,
) *AccountFacade {
	return &AccountFacade{
		repository: repository,
		factory:    factory,
		logger:     logger,
	}
}

func (f *AccountFacade) Create(name string) (*entities.BankAccount, error) {
	account, err := f.factory.CreateBankAccount(name)
	if err != nil {
		return nil, err
	}

	if err := f.repository.Save(account); err != nil {
		return nil, err
	}

	return account, nil
}

func (f *AccountFacade) GetAccount(id string) (*entities.BankAccount, error) {
	return f.repository.GetByID(id)
}

func (f *AccountFacade) UpdateAccountName(id string, name string) error {
	account, err := f.repository.GetByID(id)
	if err != nil {
		return err
	}

	if err := account.UpdateName(name); err != nil {
		return err
	}

	return f.repository.Update(account)
}

func (f *AccountFacade) SetAccountBalance(id string, amount float64) error {
	account, err := f.repository.GetByID(id)
	if err != nil {
		return err
	}

	if err := account.SetBalance(amount); err != nil {
		return err
	}

	return f.repository.Update(account)
}

func (f *AccountFacade) DeleteAccount(id string) error {
	return f.repository.Delete(id)
}

func (f *AccountFacade) GetAllAccounts() ([]*entities.BankAccount, error) {
	return f.repository.GetAll()
}
