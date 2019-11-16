package services

import (
	"github.com/ehardi19/bitment-backend/authentication"
	"github.com/ehardi19/bitment-backend/repository"
	"github.com/ehardi19/bitment-backend/repository/mock"
)

type Services struct {
	Repository *repository.Repository
	// Mock Repository
	GoalRepo        *mock.GoalMockRepository
	InvestmentRepo  *mock.InvestmentMockRepository
	ItemRepo        *mock.ItemMockRepository
	TransactionRepo *mock.TransactionMockRepository
	VendorRepo      *mock.VendorMockRepository

	Authentication *authentication.Authentication
}

func NewServices(repo *repository.Repository, auth *authentication.Authentication) *Services {
	return &Services{
		Repository:      repo,
		Authentication:  auth,
		GoalRepo:        mock.NewGoalMockRepository(),
		InvestmentRepo:  mock.NewInvestmentMockRepository(),
		ItemRepo:        mock.NewItemMockRepository(),
		TransactionRepo: mock.NewTransactionMockRepository(),
		VendorRepo:      mock.NewVendorMockRepository(),
	}
}
