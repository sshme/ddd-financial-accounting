package persistence

import (
	"ddd-financial-accounting/internal/domain/entities"
	"testing"
)

func TestCategoryRepository_Save(t *testing.T) {
	repo := NewInMemoryCategoryRepository()
	category, err := entities.NewCategory(entities.Income, "Test Category")
	if err != nil {
		t.Fatalf("Failed to create test category: %v", err)
	}

	err = repo.Save(category)
	if err != nil {
		t.Errorf("Failed to save category: %v", err)
	}

	err = repo.Save(category)
	if err == nil {
		t.Error("Expected error when saving category with duplicate ID, got nil")
	}
}

func TestCategoryRepository_GetByID(t *testing.T) {
	repo := NewInMemoryCategoryRepository()
	category, _ := entities.NewCategory(entities.Expense, "Test Category")
	_ = repo.Save(category)

	retrieved, err := repo.GetByID(category.ID)
	if err != nil {
		t.Errorf("Failed to get category by ID: %v", err)
	}
	if retrieved.ID != category.ID || retrieved.Name != category.Name || retrieved.GroupType != category.GroupType {
		t.Errorf("Retrieved category does not match original")
	}

	_, err = repo.GetByID("non-existent-id")
	if err == nil {
		t.Error("Expected error when getting non-existent category, got nil")
	}
}

func TestCategoryRepository_Update(t *testing.T) {
	repo := NewInMemoryCategoryRepository()
	category, _ := entities.NewCategory(entities.Income, "Test Category")
	_ = repo.Save(category)

	_ = category.UpdateName("Updated Category")

	err := repo.Update(category)
	if err != nil {
		t.Errorf("Failed to update category: %v", err)
	}

	retrieved, _ := repo.GetByID(category.ID)
	if retrieved.Name != "Updated Category" {
		t.Errorf("Category update did not persist")
	}

	nonExistentCategory, _ := entities.NewCategory(entities.Expense, "Non-existent")
	err = repo.Update(nonExistentCategory)
	if err == nil {
		t.Error("Expected error when updating non-existent category, got nil")
	}
}

func TestCategoryRepository_Delete(t *testing.T) {
	repo := NewInMemoryCategoryRepository()
	category, _ := entities.NewCategory(entities.Income, "Test Category")
	_ = repo.Save(category)

	err := repo.Delete(category.ID)
	if err != nil {
		t.Errorf("Failed to delete category: %v", err)
	}

	_, err = repo.GetByID(category.ID)
	if err == nil {
		t.Error("Expected error after deletion, category still exists")
	}

	err = repo.Delete(category.ID)
	if err == nil {
		t.Error("Expected error when deleting non-existent category, got nil")
	}
}

func TestCategoryRepository_GetAll(t *testing.T) {
	repo := NewInMemoryCategoryRepository()

	categories, err := repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all categories: %v", err)
	}
	if len(categories) != 0 {
		t.Errorf("Expected empty category list, got %d categories", len(categories))
	}

	category1, _ := entities.NewCategory(entities.Income, "Category 1")
	category2, _ := entities.NewCategory(entities.Expense, "Category 2")
	_ = repo.Save(category1)
	_ = repo.Save(category2)

	categories, err = repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all categories: %v", err)
	}
	if len(categories) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(categories))
	}

	foundCategory1 := false
	foundCategory2 := false
	for _, c := range categories {
		if c.ID == category1.ID {
			foundCategory1 = true
		}
		if c.ID == category2.ID {
			foundCategory2 = true
		}
	}
	if !foundCategory1 || !foundCategory2 {
		t.Error("GetAll did not return all expected categories")
	}
}
