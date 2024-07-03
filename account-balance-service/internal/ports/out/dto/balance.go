package dto

type BalanceResponse struct {
	UserID  uint64  `json:"user_id"`
	Balance float64 `json:"balance"`
}
