package services

import (
	"ddd-financial-accounting/internal/domain/entities"
	"ddd-financial-accounting/internal/domain/factories"
	"ddd-financial-accounting/internal/domain/repositories"
	"log/slog"
)

type OperationFacade struct {
	bankAccountRepository repositories.BankAccountRepository
	categoryRepository    repositories.CategoryRepository
	repository            repositories.OperationRepository
	factory               *factories.OperationFactory
	logger                *slog.Logger
}

func NewOperationFacade(
	bankAccountRepository repositories.BankAccountRepository,
	categoryRepository repositories.CategoryRepository,
	repository repositories.OperationRepository,
	factory *factories.OperationFactory,
	logger *slog.Logger,
) *OperationFacade {
	return &OperationFacade{
		bankAccountRepository: bankAccountRepository,
		categoryRepository:    categoryRepository,
		repository:            repository,
		factory:               factory,
		logger:                logger,
	}
}

func (f *OperationFacade) Create(operationType string, bankAccountId string, amount float64, description string, categoryId string) (*entities.Operation, error) {
	_, err := f.bankAccountRepository.GetByID(bankAccountId)
	if err != nil {
		return nil, err
	}

	_, err = f.categoryRepository.GetByID(categoryId)
	if err != nil {
		return nil, err
	}

	operation, err := f.factory.CreateOperation(operationType, bankAccountId, amount, description, categoryId)
	if err != nil {
		return nil, err
	}

	if err := f.repository.Save(operation); err != nil {
		return nil, err
	}

	return operation, nil
}

func (f *OperationFacade) Delete(id string) error {
	return f.repository.Delete(id)
}

func (f *OperationFacade) GetAll() ([]*entities.Operation, error) {
	return f.repository.GetAll()
}
