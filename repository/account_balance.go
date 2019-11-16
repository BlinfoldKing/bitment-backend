package repository

import "github.com/ehardi19/bitment-backend/domain"

func (r *Repository) CreateAccountBalance(accountBalance domain.AccountBalance) (domain.AccountBalance, error) {
	if err := r.DB.Model(&accountBalance).Create(&accountBalance).Error; err != nil {
		return accountBalance, err
	}

	return accountBalance, nil
}

func (r *Repository) GetAccountBalanceByUser(userID int64) (domain.AccountBalance, error) {
	var accountBalance domain.AccountBalance

	if err := r.DB.Model(&accountBalance).Where("user_id = ?", userID).Find(&accountBalance).Error; err != nil {
		return accountBalance, err
	}

	return accountBalance, nil
}
