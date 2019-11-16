package mock

import (
	"errors"

	"github.com/ehardi19/bitment-backend/domain"
)

type TransactionMockRepository struct {
	db map[int64]domain.Transaction
}

func NewTransactionMockRepository() *TransactionMockRepository {
	return &TransactionMockRepository{
		db: map[int64]domain.Transaction{},
	}
}

func (g *TransactionMockRepository) GetAllTransactionByUser(userID int64) []domain.Transaction {
	var Transactions []domain.Transaction

	for _, Transaction := range g.db {
		if Transaction.UserID == userID {
			Transactions = append(Transactions, Transaction)
		}
	}

	return Transactions
}

func (g *TransactionMockRepository) CreateTransaction(Transaction domain.Transaction) (domain.Transaction, error) {
	_, exist := g.db[Transaction.ID]
	if exist {
		return domain.Transaction{}, errors.New("Transaction with that id already existed")
	}

	g.db[Transaction.ID] = Transaction

	return Transaction, nil
}

func (g *TransactionMockRepository) GetTransactionByID(id int64) (domain.Transaction, error) {
	Transaction, ok := g.db[id]
	if ok {
		return Transaction, nil
	}

	return domain.Transaction{}, errors.New("not found")
}
