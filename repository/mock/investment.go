package mock

import (
	"errors"

	"github.com/ehardi19/bitment-backend/domain"
)

type InvestmentMockRepository struct {
	db map[int64]domain.Investment
}

func NewInvestmentMockRepository() *InvestmentMockRepository {
	return &InvestmentMockRepository{
		db: map[int64]domain.Investment{},
	}
}

func (g *InvestmentMockRepository) GetAllInvestmentByUser(userID int64) []domain.Investment {
	var Investments []domain.Investment

	for _, Investment := range g.db {
		if Investment.UserID == userID {
			Investments = append(Investments, Investment)
		}
	}

	return Investments
}

func (g *InvestmentMockRepository) CreateInvestment(Investment domain.Investment) (domain.Investment, error) {
	_, exist := g.db[Investment.ID]
	if exist {
		return domain.Investment{}, errors.New("Investment with that id already existed")
	}

	g.db[Investment.ID] = Investment

	return Investment, nil
}

func (g *InvestmentMockRepository) GetInvestmentByID(id int64) (domain.Investment, error) {
	Investment, ok := g.db[id]
	if ok {
		return Investment, nil
	}

	return domain.Investment{}, errors.New("not found")
}
