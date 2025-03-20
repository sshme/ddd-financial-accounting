package factories

import (
	"ddd-financial-accounting/internal/domain/entities"
	"errors"
	"log/slog"
)

type OperationFactory struct {
	logger *slog.Logger
}

func NewOperationFactory(logger *slog.Logger) *OperationFactory {
	return &OperationFactory{
		logger: logger,
	}
}

func (f *OperationFactory) CreateOperation(operationType string, bankAccountId string, amount float64, description string, categoryId string) (*entities.Operation, error) {
	switch operationType {
	case string(entities.Income), string(entities.Expense):
		operation, err := entities.NewOperation(entities.GroupByEnum(operationType), bankAccountId, amount, description, categoryId)
		if err != nil {
			f.logger.Error("Failed to create operation", "error", err)
			return nil, err
		}
		return operation, nil
	default:
		return nil, errors.New("invalid operation type")
	}
}
