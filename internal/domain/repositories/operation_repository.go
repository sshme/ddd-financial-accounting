package repositories

import (
	"ddd-financial-accounting/internal/domain/entities"
)

type OperationRepository interface {
	Save(operation *entities.Operation) error
	Delete(id string) error
	GetAll() ([]*entities.Operation, error)
}
