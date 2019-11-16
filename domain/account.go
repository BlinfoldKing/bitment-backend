package domain

type AccountBalance struct {
	UserID  int64   `json:"user_id"`
	Balance float64 `json:"balance"`
}

type AccountCredit struct {
	UserID  int64   `json:"user_id"`
	Credits float64 `json:"credits"`
}
