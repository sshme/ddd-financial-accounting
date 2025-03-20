package services

import (
	"ddd-financial-accounting/internal/domain/entities"
	"ddd-financial-accounting/internal/domain/factories"
	"ddd-financial-accounting/internal/domain/repositories"
	"errors"
	"log/slog"
)

type CategoryFacade struct {
	repository repositories.CategoryRepository
	factory    *factories.CategoryFactory
	logger     *slog.Logger
}

func NewCategoryFacade(
	repository repositories.CategoryRepository,
	factory *factories.CategoryFactory,
	logger *slog.Logger,
) *CategoryFacade {
	return &CategoryFacade{
		repository: repository,
		factory:    factory,
		logger:     logger,
	}
}

func (f *CategoryFacade) Create(name string) (*entities.Category, *entities.Category, error) {
	categoryIncome, err := f.factory.CreateCategory(string(entities.Income), name)
	if err != nil {
		return nil, nil, err
	}

	categoryExpense, err := f.factory.CreateCategory(string(entities.Expense), name)
	if err != nil {
		return nil, nil, err
	}

	if err := f.repository.Save(categoryIncome); err != nil {
		return nil, nil, err
	}

	if err := f.repository.Save(categoryExpense); err != nil {
		err := f.repository.Delete(categoryIncome.ID)
		if err != nil {
			return nil, nil, errors.New("failed to delete income category after failing to save expense category")
		}
		return nil, nil, err
	}

	return categoryIncome, categoryExpense, nil
}

func (f *CategoryFacade) Delete(id string) error {
	return f.repository.Delete(id)
}

func (f *CategoryFacade) GetAll() ([]*entities.Category, error) {
	return f.repository.GetAll()
}
