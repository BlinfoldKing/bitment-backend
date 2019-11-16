package domain

type Goal struct {
	ID             int64   `json:"id"`
	UserID         int64   `json:"user_id"`
	Title          string  `json:"title"`
	MonthlyDeposit float64 `json:"monthly_deposit"`
	TargetBalance  float64 `json:"target_balance"`
	TargetYear     int     `json:"target_year"`
}
