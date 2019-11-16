package services

import (
	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) GetAllVendor() ([]domain.Vendor, error) {
	return s.VendorRepo.GetAllVendor(), nil
}

func (s *Services) GetVendorByID(id int64) (domain.Vendor, error) {
	return s.VendorRepo.GetVendorByID(id)
}
