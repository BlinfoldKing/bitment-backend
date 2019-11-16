package services

import (
	"time"

	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) GetTotalInvestmentProjectionByUser(userID int64) (float64, error) {
	investments := s.InvestmentRepo.GetAllInvestmentByUser(userID)

	var ret float64

	for _, investment := range investments {
		vendor, err := s.VendorRepo.GetVendorByID(investment.VendorID)
		if err != nil {
			return 0, err
		}

		ret += vendor.TotalFunds
	}

	return ret, nil
}

func (s *Services) GetAllInvestmentByUser(userID int64) ([]domain.Investment, error) {
	return s.InvestmentRepo.GetAllInvestmentByUser(userID), nil
}

func (s *Services) CreateInvestment(Investment domain.Investment) (domain.Investment, int, error) {
	Investment.ID = time.Now().UnixNano()

	Investment, err := s.InvestmentRepo.CreateInvestment(Investment)
	if err != nil {
		return domain.Investment{}, 400, err
	}

	return Investment, 201, nil
}

func (s *Services) GetInvestmentByID(id int64) (domain.Investment, error) {
	return s.InvestmentRepo.GetInvestmentByID(id)
}
