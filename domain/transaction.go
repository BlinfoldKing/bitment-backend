package domain

import "time"

type Transaction struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	ItemID       int64     `json:"item_id"`
	NumberOfItem int       `json:"number_of_item"`
	Amount       float64   `json:"amount"`
	Cashback     float64   `json:"cashback"`
	Date         time.Time `json:"time"`
	Description  string    `json:"description"`
	Hash         string    `json:"hash"`
}
