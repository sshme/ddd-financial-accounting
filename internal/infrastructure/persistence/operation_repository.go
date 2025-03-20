package persistence

import (
	"ddd-financial-accounting/internal/domain/entities"
	"ddd-financial-accounting/internal/domain/repositories"
	"errors"
	"sync"
)

var _ repositories.OperationRepository = (*OperationRepository)(nil)

type OperationRepository struct {
	operations map[string]*entities.Operation
	mutex      sync.RWMutex
}

func (r *OperationRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.operations[id]; !exists {
		return errors.New("operation not found")
	}

	delete(r.operations, id)
	return nil
}

func NewInMemoryOperationRepository() *OperationRepository {
	return &OperationRepository{
		operations: make(map[string]*entities.Operation),
	}
}

func (r *OperationRepository) Save(category *entities.Operation) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.operations[category.ID]; exists {
		return errors.New("operation with this ID already exists")
	}

	r.operations[category.ID] = category
	return nil
}

func (r *OperationRepository) GetAll() ([]*entities.Operation, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var operations []*entities.Operation
	for _, operation := range r.operations {
		operations = append(operations, operation)
	}

	return operations, nil
}
