package persistence

import (
	"ddd-financial-accounting/internal/domain/entities"
	"testing"
)

func TestOperationRepository_Save(t *testing.T) {
	repo := NewInMemoryOperationRepository()
	operation, err := entities.NewOperation(entities.Income, "bank123", 100.50, "Test Operation", "cat123")
	if err != nil {
		t.Fatalf("Failed to create test operation: %v", err)
	}

	err = repo.Save(operation)
	if err != nil {
		t.Errorf("Failed to save operation: %v", err)
	}

	err = repo.Save(operation)
	if err == nil {
		t.Error("Expected error when saving operation with duplicate ID, got nil")
	}
}

func TestOperationRepository_Delete(t *testing.T) {
	repo := NewInMemoryOperationRepository()
	operation, _ := entities.NewOperation(entities.Income, "bank123", 100.0, "Test Operation", "cat123")
	_ = repo.Save(operation)

	err := repo.Delete(operation.ID)
	if err != nil {
		t.Errorf("Failed to delete operation: %v", err)
	}

	err = repo.Delete(operation.ID)
	if err == nil {
		t.Error("Expected error when deleting non-existent operation, got nil")
	}
}

func TestOperationRepository_GetAll(t *testing.T) {
	repo := NewInMemoryOperationRepository()

	operations, err := repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all operations: %v", err)
	}
	if len(operations) != 0 {
		t.Errorf("Expected empty operation list, got %d operations", len(operations))
	}

	op1, _ := entities.NewOperation(entities.Income, "bank123", 100.0, "Operation 1", "cat123")
	op2, _ := entities.NewOperation(entities.Expense, "bank456", 50.0, "Operation 2", "cat456")
	_ = repo.Save(op1)
	_ = repo.Save(op2)

	operations, err = repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all operations: %v", err)
	}
	if len(operations) != 2 {
		t.Errorf("Expected 2 operations, got %d", len(operations))
	}

	foundOp1 := false
	foundOp2 := false
	for _, op := range operations {
		if op.ID == op1.ID {
			foundOp1 = true
		}
		if op.ID == op2.ID {
			foundOp2 = true
		}
	}
	if !foundOp1 || !foundOp2 {
		t.Error("GetAll did not return all expected operations")
	}
}
