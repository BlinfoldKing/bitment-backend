package repository

import "github.com/ehardi19/bitment-backend/domain"

func (r *Repository) CreateAccountCredit(accountCredit domain.AccountCredit) (domain.AccountCredit, error) {
	if err := r.DB.Model(&accountCredit).Create(&accountCredit).Error; err != nil {
		return accountCredit, err
	}

	return accountCredit, nil
}

func (r *Repository) GetAccountCreditByUser(userID int64) (domain.AccountCredit, error) {
	var accountCredit domain.AccountCredit

	if err := r.DB.Model(&accountCredit).Where("user_id = ?", userID).Find(&accountCredit).Error; err != nil {
		return accountCredit, err
	}

	return accountCredit, nil
}
