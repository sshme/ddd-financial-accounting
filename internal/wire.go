//go:build wireinject
// +build wireinject

package internal

import (
	"ddd-financial-accounting/internal/application/services"
	"ddd-financial-accounting/internal/domain/factories"
	"ddd-financial-accounting/internal/domain/repositories"
	"ddd-financial-accounting/internal/infrastructure/persistence"
	"ddd-financial-accounting/pkg/logger"
	"github.com/google/wire"
	"log/slog"
)

// Application represents the main application with all dependencies
type Application struct {
	AccountFacade   *services.AccountFacade
	CategoryFacade  *services.CategoryFacade
	OperationFacade *services.OperationFacade
	Logger          *slog.Logger
}

func provideLogger() *slog.Logger {
	return logger.SetupLogger()
}

func provideBankAccountRepository() repositories.BankAccountRepository {
	return persistence.NewInMemoryBankAccountRepository()
}

func provideCategoryRepository() repositories.CategoryRepository {
	return persistence.NewInMemoryCategoryRepository()
}

func provideOperationRepository() repositories.OperationRepository {
	return persistence.NewInMemoryOperationRepository()
}

func provideBankAccountFactory(logger *slog.Logger) *factories.BankAccountFactory {
	return factories.NewBankAccountFactory(logger)
}

func provideCategoryFactory(logger *slog.Logger) *factories.CategoryFactory {
	return factories.NewCategoryFactory(logger)
}

func provideOperationFactory(logger *slog.Logger) *factories.OperationFactory {
	return factories.NewOperationFactory(logger)
}

var RepositorySet = wire.NewSet(
	provideBankAccountRepository,
	provideOperationRepository,
	provideCategoryRepository,
)

var FactorySet = wire.NewSet(
	provideBankAccountFactory,
	provideOperationFactory,
	provideCategoryFactory,
)

var ApplicationSet = wire.NewSet(
	wire.Struct(new(Application), "*"),
	services.NewOperationFacade,
	services.NewCategoryFacade,
	services.NewAccountFacade,
	provideLogger,
	RepositorySet,
	FactorySet,
)

// InitializeApp is the function that will be implemented by wire
func InitializeApp() (*Application, error) {
	wire.Build(ApplicationSet)
	return &Application{}, nil
}
