package entities

import (
	"testing"
)

func TestNewCategory(t *testing.T) {
	tests := []struct {
		name      string
		groupType GroupByEnum
		catName   string
		wantErr   bool
		errorMsg  string
	}{
		{
			name:      "valid income category",
			groupType: Income,
			catName:   "Salary",
			wantErr:   false,
		},
		{
			name:      "valid expense category",
			groupType: Expense,
			catName:   "Groceries",
			wantErr:   false,
		},
		{
			name:      "empty name",
			groupType: Expense,
			catName:   "",
			wantErr:   true,
			errorMsg:  "category name cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cat, err := NewCategory(tt.groupType, tt.catName)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if err.Error() != tt.errorMsg {
					t.Fatalf("expected error message %q, got %q", tt.errorMsg, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if cat.Name != tt.catName {
				t.Errorf("expected name %q, got %q", tt.catName, cat.Name)
			}

			if cat.GroupType != tt.groupType {
				t.Errorf("expected group type %q, got %q", tt.groupType, cat.GroupType)
			}

			if cat.ID == "" {
				t.Error("ID should not be empty")
			}
		})
	}
}

func TestCategory_UpdateName(t *testing.T) {
	tests := []struct {
		name     string
		newName  string
		wantErr  bool
		errorMsg string
	}{
		{
			name:    "valid name update",
			newName: "Updated Category",
			wantErr: false,
		},
		{
			name:     "empty name",
			newName:  "",
			wantErr:  true,
			errorMsg: "category name cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cat, _ := NewCategory(Income, "Initial Category")
			originalName := cat.Name

			err := cat.UpdateName(tt.newName)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if err.Error() != tt.errorMsg {
					t.Fatalf("expected error message %q, got %q", tt.errorMsg, err.Error())
				}
				// Check that name wasn't changed
				if cat.Name != originalName {
					t.Errorf("name should not change on error, expected %q, got %q",
						originalName, cat.Name)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if cat.Name != tt.newName {
				t.Errorf("expected name %q, got %q", tt.newName, cat.Name)
			}
		})
	}
}
