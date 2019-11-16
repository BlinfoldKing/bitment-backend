package domain

type Vendor struct {
	ID              int64   `json:"id"`
	VendorName      string  `json:"vendor_name"`
	Description     string  `json:"description"`
	Peformance      float64 `json:"peformance"`
	MinimumPurchase float64 `json:"minimum_purchase"`
	TotalFunds      float64 `json:"total_funds"`
}
