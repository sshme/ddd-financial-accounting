package factories

import (
	"ddd-financial-accounting/internal/domain/entities"
	"log/slog"
)

type BankAccountFactory struct {
	logger *slog.Logger
}

func NewBankAccountFactory(logger *slog.Logger) *BankAccountFactory {
	return &BankAccountFactory{
		logger: logger,
	}
}

func (f *BankAccountFactory) CreateBankAccount(name string) (*entities.BankAccount, error) {
	account, err := entities.NewBankAccount(name)
	if err != nil {
		f.logger.Error("Failed to create bank account", "error", err, "name", name)
		return nil, err
	}
	return account, nil
}
