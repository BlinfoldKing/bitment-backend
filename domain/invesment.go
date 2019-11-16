package domain

type Investment struct {
	ID       int64 `json:"id"`
	UserID   int64 `json:"user_id"`
	VendorID int64 `json:"vendor_id"`
}
