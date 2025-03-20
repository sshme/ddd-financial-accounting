package persistence

import (
	"ddd-financial-accounting/internal/domain/entities"
	"ddd-financial-accounting/internal/domain/repositories"
	"errors"
	"sync"
)

var _ repositories.CategoryRepository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	categories map[string]*entities.Category
	mutex      sync.RWMutex
}

func NewInMemoryCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: make(map[string]*entities.Category),
	}
}

func (r *CategoryRepository) Save(category *entities.Category) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.categories[category.ID]; exists {
		return errors.New("category with this ID already exists")
	}

	r.categories[category.ID] = category
	return nil
}

func (r *CategoryRepository) GetByID(id string) (*entities.Category, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	account, exists := r.categories[id]
	if !exists {
		return nil, errors.New("category not found")
	}

	return account, nil
}

func (r *CategoryRepository) Update(category *entities.Category) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.categories[category.ID]; !exists {
		return errors.New("category not found")
	}

	r.categories[category.ID] = category
	return nil
}

func (r *CategoryRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.categories[id]; !exists {
		return errors.New("category not found")
	}

	delete(r.categories, id)
	return nil
}

func (r *CategoryRepository) GetAll() ([]*entities.Category, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	accounts := make([]*entities.Category, 0, len(r.categories))
	for _, account := range r.categories {
		accounts = append(accounts, account)
	}

	return accounts, nil
}
