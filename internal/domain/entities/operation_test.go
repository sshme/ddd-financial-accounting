package entities

import (
	"testing"
	"time"
)

func TestNewOperation(t *testing.T) {
	tests := []struct {
		name          string
		operationType GroupByEnum
		bankAccountId string
		amount        float64
		description   string
		categoryId    string
		wantErr       bool
		errorMsg      string
	}{
		{
			name:          "valid income operation",
			operationType: Income,
			bankAccountId: "bank123",
			amount:        100.50,
			description:   "Salary payment",
			categoryId:    "cat123",
			wantErr:       false,
		},
		{
			name:          "valid expense operation",
			operationType: Expense,
			bankAccountId: "bank123",
			amount:        50.25,
			description:   "Grocery shopping",
			categoryId:    "cat456",
			wantErr:       false,
		},
		{
			name:          "invalid negative amount",
			operationType: Income,
			bankAccountId: "bank123",
			amount:        -10.0,
			description:   "Negative amount",
			categoryId:    "cat123",
			wantErr:       true,
			errorMsg:      "operation amount must be greater or equal than zero",
		},
		{
			name:          "valid zero amount",
			operationType: Expense,
			bankAccountId: "bank123",
			amount:        0,
			description:   "Zero cost item",
			categoryId:    "cat456",
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := time.Now()

			op, err := NewOperation(tt.operationType, tt.bankAccountId, tt.amount, tt.description, tt.categoryId)

			after := time.Now()

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

			if op.Type != tt.operationType {
				t.Errorf("expected operation type %q, got %q", tt.operationType, op.Type)
			}

			if op.BankAccountId != tt.bankAccountId {
				t.Errorf("expected bank account ID %q, got %q", tt.bankAccountId, op.BankAccountId)
			}

			if op.Amount != tt.amount {
				t.Errorf("expected amount %f, got %f", tt.amount, op.Amount)
			}

			if op.Description != tt.description {
				t.Errorf("expected description %q, got %q", tt.description, op.Description)
			}

			if op.CategoryId != tt.categoryId {
				t.Errorf("expected category ID %q, got %q", tt.categoryId, op.CategoryId)
			}

			if op.ID == "" {
				t.Error("ID should not be empty")
			}

			if op.Date.Before(before) || op.Date.After(after) {
				t.Errorf("expected date between %v and %v, got %v", before, after, op.Date)
			}
		})
	}
}
