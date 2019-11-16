package services

import (
	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) GetAllItem() ([]domain.Item, error) {
	return s.ItemRepo.GetAllItem(), nil
}

func (s *Services) GetItemByID(id int64) (domain.Item, error) {
	return s.ItemRepo.GetItemByID(id)
}
