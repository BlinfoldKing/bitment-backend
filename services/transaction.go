package services

import (
	"strconv"
	"time"

	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) GetAllTransactionByUser(userID int64) ([]domain.Transaction, error) {
	return s.TransactionRepo.GetAllTransactionByUser(userID), nil
}

func (s *Services) CreateTransaction(Transaction domain.Transaction) (domain.Transaction, int, error) {
	Item, err := s.ItemRepo.GetItemByID(Transaction.ItemID)
	if err != nil {
		return domain.Transaction{}, 400, err
	}

	Transaction.ID = time.Now().UnixNano()
	Transaction.Amount = Item.Price * float64(Transaction.NumberOfItem)
	Transaction.Cashback = Transaction.Amount * Item.CashbackPercentage
	Transaction.Date = time.Now()
	Transaction.Hash, err = s.Authentication.HashPassword(strconv.FormatInt(Transaction.ID, 10))
	if err != nil {
		return domain.Transaction{}, 400, err
	}

	Transaction, err = s.TransactionRepo.CreateTransaction(Transaction)
	if err != nil {
		return domain.Transaction{}, 400, err
	}

	return Transaction, 201, nil
}

func (s *Services) GetTransactionByID(id int64) (domain.Transaction, error) {
	Transaction, err := s.TransactionRepo.GetTransactionByID(id)
	if err != nil {
		return domain.Transaction{}, err
	}

	ok := s.Authentication.CheckPasswordHash(Transaction.Hash, strconv.FormatInt(Transaction.ID, 10))
	if !ok {
		return domain.Transaction{}, err
	}

	return Transaction, nil
}
