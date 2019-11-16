package services

import (
	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) GetAccountBalanceByUser(userID int64) (domain.AccountBalance, error) {
	balance, err := s.Repository.GetAccountBalanceByUser(userID)
	if err != nil {
		return domain.AccountBalance{}, err
	}

	transactions := s.TransactionRepo.GetAllTransactionByUser(balance.UserID)
	for _, transaction := range transactions {
		balance.Balance -= transaction.Amount
	}

	return balance, nil
}

func (s *Services) GetAccountCreditByUser(userID int64) (domain.AccountCredit, error) {
	Credit, err := s.Repository.GetAccountCreditByUser(userID)
	if err != nil {
		return domain.AccountCredit{}, err
	}

	transactions := s.TransactionRepo.GetAllTransactionByUser(Credit.UserID)
	for _, transaction := range transactions {
		Credit.Credits += transaction.Cashback
	}

	goals := s.GoalRepo.GetAllGoalByUser(Credit.UserID)
	for _, goal := range goals {
		Credit.Credits -= goal.MonthlyDeposit
	}

	return Credit, nil
}
