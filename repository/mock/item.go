package mock

import (
	"errors"
	"time"

	"github.com/ehardi19/bitment-backend/domain"
)

type ItemMockRepository struct {
	db map[int64]domain.Item
}

func NewItemMockRepository() *ItemMockRepository {
	pizzaHut := domain.Item{
		ID:                 time.Now().UnixNano(),
		MerchantName:       "Pizza Hut",
		ItemName:           "Double Box Jumbo",
		Price:              200000,
		Description:        "-",
		CashbackPercentage: 0.5,
	}
	return &ItemMockRepository{
		db: map[int64]domain.Item{
			pizzaHut.ID: pizzaHut,
		},
	}
}

func (g *ItemMockRepository) GetAllItem() []domain.Item {
	var Items []domain.Item

	for _, Item := range g.db {
		Items = append(Items, Item)
	}

	return Items
}

func (g *ItemMockRepository) GetItemByID(id int64) (domain.Item, error) {
	Item, ok := g.db[id]
	if ok {
		return Item, nil
	}

	return domain.Item{}, errors.New("not found")
}
