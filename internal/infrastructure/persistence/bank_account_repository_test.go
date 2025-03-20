package persistence

import (
	"ddd-financial-accounting/internal/domain/entities"
	"testing"
)

func TestBankAccountRepository_Save(t *testing.T) {
	repo := NewInMemoryBankAccountRepository()
	account, err := entities.NewBankAccount("Test Account")
	if err != nil {
		t.Fatalf("Failed to create test account: %v", err)
	}

	err = repo.Save(account)
	if err != nil {
		t.Errorf("Failed to save account: %v", err)
	}

	err = repo.Save(account)
	if err == nil {
		t.Error("Expected error when saving account with duplicate ID, got nil")
	}
}

func TestBankAccountRepository_GetByID(t *testing.T) {
	repo := NewInMemoryBankAccountRepository()
	account, _ := entities.NewBankAccount("Test Account")
	_ = repo.Save(account)

	retrieved, err := repo.GetByID(account.ID)
	if err != nil {
		t.Errorf("Failed to get account by ID: %v", err)
	}
	if retrieved.ID != account.ID || retrieved.Name != account.Name {
		t.Errorf("Retrieved account does not match original")
	}

	_, err = repo.GetByID("non-existent-id")
	if err == nil {
		t.Error("Expected error when getting non-existent account, got nil")
	}
}

func TestBankAccountRepository_Update(t *testing.T) {
	repo := NewInMemoryBankAccountRepository()
	account, _ := entities.NewBankAccount("Test Account")
	_ = repo.Save(account)

	account.Name = "Updated Account"
	_ = account.SetBalance(100.0)

	err := repo.Update(account)
	if err != nil {
		t.Errorf("Failed to update account: %v", err)
	}

	retrieved, _ := repo.GetByID(account.ID)
	if retrieved.Name != "Updated Account" || retrieved.Balance != 100.0 {
		t.Errorf("Account update did not persist")
	}

	nonExistentAccount, _ := entities.NewBankAccount("Non-existent")
	err = repo.Update(nonExistentAccount)
	if err == nil {
		t.Error("Expected error when updating non-existent account, got nil")
	}
}

func TestBankAccountRepository_Delete(t *testing.T) {
	repo := NewInMemoryBankAccountRepository()
	account, _ := entities.NewBankAccount("Test Account")
	_ = repo.Save(account)

	err := repo.Delete(account.ID)
	if err != nil {
		t.Errorf("Failed to delete account: %v", err)
	}

	_, err = repo.GetByID(account.ID)
	if err == nil {
		t.Error("Expected error after deletion, account still exists")
	}

	err = repo.Delete(account.ID)
	if err == nil {
		t.Error("Expected error when deleting non-existent account, got nil")
	}
}

func TestBankAccountRepository_GetAll(t *testing.T) {
	repo := NewInMemoryBankAccountRepository()

	accounts, err := repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all accounts: %v", err)
	}
	if len(accounts) != 0 {
		t.Errorf("Expected empty account list, got %d accounts", len(accounts))
	}

	account1, _ := entities.NewBankAccount("Account 1")
	account2, _ := entities.NewBankAccount("Account 2")
	_ = repo.Save(account1)
	_ = repo.Save(account2)

	accounts, err = repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all accounts: %v", err)
	}
	if len(accounts) != 2 {
		t.Errorf("Expected 2 accounts, got %d", len(accounts))
	}

	foundAccount1 := false
	foundAccount2 := false
	for _, a := range accounts {
		if a.ID == account1.ID {
			foundAccount1 = true
		}
		if a.ID == account2.ID {
			foundAccount2 = true
		}
	}
	if !foundAccount1 || !foundAccount2 {
		t.Error("GetAll did not return all expected accounts")
	}
}
