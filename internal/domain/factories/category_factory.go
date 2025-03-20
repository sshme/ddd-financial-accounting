package factories

import (
	"ddd-financial-accounting/internal/domain/entities"
	"errors"
	"log/slog"
)

type CategoryFactory struct {
	logger *slog.Logger
}

func NewCategoryFactory(logger *slog.Logger) *CategoryFactory {
	return &CategoryFactory{
		logger: logger,
	}
}

func (f *CategoryFactory) CreateCategory(groupType string, name string) (*entities.Category, error) {
	switch groupType {
	case string(entities.Income), string(entities.Expense):
		account, err := entities.NewCategory(entities.GroupByEnum(groupType), name)
		if err != nil {
			f.logger.Error("Failed to create category", "error", err, "name", name)
			return nil, err
		}
		return account, nil
	default:
		return nil, errors.New("invalid group type")
	}
}
