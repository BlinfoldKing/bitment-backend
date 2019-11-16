package domain

type Item struct {
	ID                 int64   `json:"id"`
	MerchantName       string  `json:"merchant_name"`
	ItemName           string  `json:"item_name"`
	Price              float64 `json:"price"`
	Description        string  `json:"description"`
	CashbackPercentage float64 `json:"cashback_percentage"`
}
