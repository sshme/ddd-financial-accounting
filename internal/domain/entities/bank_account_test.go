package entities

import (
	"testing"
)

func TestNewBankAccount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		errorMsg string
	}{
		{
			name:    "valid bank account",
			input:   "Checking Account",
			wantErr: false,
		},
		{
			name:     "empty name",
			input:    "",
			wantErr:  true,
			errorMsg: "bank account name cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba, err := NewBankAccount(tt.input)

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

			if ba.Name != tt.input {
				t.Errorf("expected name %q, got %q", tt.input, ba.Name)
			}

			if ba.Balance != 0 {
				t.Errorf("expected balance 0, got %f", ba.Balance)
			}

			if ba.ID == "" {
				t.Error("ID should not be empty")
			}
		})
	}
}

func TestBankAccount_SetBalance(t *testing.T) {
	tests := []struct {
		name     string
		balance  float64
		wantErr  bool
		errorMsg string
	}{
		{
			name:    "positive balance",
			balance: 100.50,
			wantErr: false,
		},
		{
			name:    "zero balance",
			balance: 0,
			wantErr: false,
		},
		{
			name:     "negative balance",
			balance:  -50.25,
			wantErr:  true,
			errorMsg: "insufficient funds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba, _ := NewBankAccount("Test Account")

			err := ba.SetBalance(tt.balance)

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

			if ba.Balance != tt.balance {
				t.Errorf("expected balance %f, got %f", tt.balance, ba.Balance)
			}
		})
	}
}

func TestBankAccount_UpdateBalance(t *testing.T) {
	tests := []struct {
		name           string
		initialAmount  float64
		updateAmount   float64
		expectedAmount float64
		wantErr        bool
		errorMsg       string
	}{
		{
			name:           "add funds",
			initialAmount:  100,
			updateAmount:   50,
			expectedAmount: 150,
			wantErr:        false,
		},
		{
			name:           "withdraw valid amount",
			initialAmount:  100,
			updateAmount:   -50,
			expectedAmount: 50,
			wantErr:        false,
		},
		{
			name:           "withdraw too much",
			initialAmount:  100,
			updateAmount:   -150,
			expectedAmount: 100,
			wantErr:        true,
			errorMsg:       "insufficient funds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba, _ := NewBankAccount("Test Account")
			_ = ba.SetBalance(tt.initialAmount)

			err := ba.UpdateBalance(tt.updateAmount)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if err.Error() != tt.errorMsg {
					t.Fatalf("expected error message %q, got %q", tt.errorMsg, err.Error())
				}
				if ba.Balance != tt.initialAmount {
					t.Errorf("balance should not change on error, expected %f, got %f",
						tt.initialAmount, ba.Balance)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if ba.Balance != tt.expectedAmount {
				t.Errorf("expected balance %f, got %f", tt.expectedAmount, ba.Balance)
			}
		})
	}
}

func TestBankAccount_UpdateName(t *testing.T) {
	tests := []struct {
		name     string
		newName  string
		wantErr  bool
		errorMsg string
	}{
		{
			name:    "valid name update",
			newName: "Updated Account",
			wantErr: false,
		},
		{
			name:     "empty name",
			newName:  "",
			wantErr:  true,
			errorMsg: "bank account name cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba, _ := NewBankAccount("Initial Account")
			originalName := ba.Name

			err := ba.UpdateName(tt.newName)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if err.Error() != tt.errorMsg {
					t.Fatalf("expected error message %q, got %q", tt.errorMsg, err.Error())
				}
				if ba.Name != originalName {
					t.Errorf("name should not change on error, expected %q, got %q",
						originalName, ba.Name)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if ba.Name != tt.newName {
				t.Errorf("expected name %q, got %q", tt.newName, ba.Name)
			}
		})
	}
}
