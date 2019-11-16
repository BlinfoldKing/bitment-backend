package repository

import (
	"github.com/ehardi19/bitment-backend/domain"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *Repository {
	DB.AutoMigrate(
		&domain.User{},
		&domain.AccountBalance{},
		&domain.AccountCredit{},
	)

	return &Repository{
		DB: DB,
	}
}
