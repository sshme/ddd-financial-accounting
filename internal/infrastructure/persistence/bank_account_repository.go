package persistence

import (
	"ddd-financial-accounting/internal/domain/entities"
	"ddd-financial-accounting/internal/domain/repositories"
	"errors"
	"sync"
)

var _ repositories.BankAccountRepository = (*BankAccountRepository)(nil)

type BankAccountRepository struct {
	accounts map[string]*entities.BankAccount
	mutex    sync.RWMutex
}

func NewInMemoryBankAccountRepository() *BankAccountRepository {
	return &BankAccountRepository{
		accounts: make(map[string]*entities.BankAccount),
	}
}

func (r *BankAccountRepository) Save(account *entities.BankAccount) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.accounts[account.ID]; exists {
		return errors.New("account with this ID already exists")
	}

	r.accounts[account.ID] = account
	return nil
}

func (r *BankAccountRepository) GetByID(id string) (*entities.BankAccount, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	account, exists := r.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}

func (r *BankAccountRepository) Update(account *entities.BankAccount) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.accounts[account.ID]; !exists {
		return errors.New("account not found")
	}

	r.accounts[account.ID] = account
	return nil
}

func (r *BankAccountRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.accounts[id]; !exists {
		return errors.New("account not found")
	}

	delete(r.accounts, id)
	return nil
}

func (r *BankAccountRepository) GetAll() ([]*entities.BankAccount, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	accounts := make([]*entities.BankAccount, 0, len(r.accounts))
	for _, account := range r.accounts {
		accounts = append(accounts, account)
	}

	return accounts, nil
}
