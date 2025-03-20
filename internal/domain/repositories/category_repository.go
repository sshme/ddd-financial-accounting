package repositories

import (
	"ddd-financial-accounting/internal/domain/entities"
)

type CategoryRepository interface {
	Save(category *entities.Category) error
	GetByID(id string) (*entities.Category, error)
	Update(category *entities.Category) error
	Delete(id string) error
	GetAll() ([]*entities.Category, error)
}
