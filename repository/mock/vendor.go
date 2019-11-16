package mock

import (
	"errors"
	"time"

	"github.com/ehardi19/bitment-backend/domain"
)

type VendorMockRepository struct {
	db map[int64]domain.Vendor
}

func NewVendorMockRepository() *VendorMockRepository {
	makmurSanjaya := domain.Vendor{
		ID:              time.Now().UnixNano(),
		VendorName:      "PT. PHBD Makmur Sanjaya",
		Description:     "Digital Venture of Capital City",
		MinimumPurchase: 100000,
		Peformance:      0.576,
		TotalFunds:      300000000,
	}
	niagaCahaya := domain.Vendor{
		ID:              time.Now().UnixNano(),
		VendorName:      "Niaga Cahaya Makmur",
		Description:     "Capital Market",
		MinimumPurchase: 50000,
		Peformance:      0.424,
		TotalFunds:      510004340,
	}
	bareksa := domain.Vendor{
		ID:              time.Now().UnixNano(),
		VendorName:      "Bareksa TradeX",
		Description:     "Market Analysis Consultant",
		MinimumPurchase: 10000,
		Peformance:      0.324,
		TotalFunds:      352022210,
	}
	angkasaMitra := domain.Vendor{
		ID:              time.Now().UnixNano(),
		VendorName:      "Angkasa Mitra Bersama",
		Description:     "Airplane Factory",
		MinimumPurchase: 100000,
		Peformance:      0.611,
		TotalFunds:      763012261,
	}
	tradiaLegion := domain.Vendor{
		ID:              time.Now().UnixNano(),
		VendorName:      "Tradia Legion Corporation",
		Description:     "International Trading Corporation",
		MinimumPurchase: 100000,
		Peformance:      0.576,
		TotalFunds:      654112624,
	}
	return &VendorMockRepository{
		db: map[int64]domain.Vendor{
			makmurSanjaya.ID: makmurSanjaya,
			niagaCahaya.ID:   niagaCahaya,
			bareksa.ID:       bareksa,
			angkasaMitra.ID:  angkasaMitra,
			tradiaLegion.ID:  tradiaLegion,
		},
	}
}

func (g *VendorMockRepository) GetAllVendor() []domain.Vendor {
	var Vendors []domain.Vendor

	for _, Vendor := range g.db {
		Vendors = append(Vendors, Vendor)
	}

	return Vendors
}

func (g *VendorMockRepository) GetVendorByID(id int64) (domain.Vendor, error) {
	Vendor, ok := g.db[id]
	if ok {
		return Vendor, nil
	}

	return domain.Vendor{}, errors.New("not found")
}
