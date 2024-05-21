package kafka

type CreditRequest struct {
	UserID uint64  `json:"user_id"`
	Amount float64 `json:"amount"`
}
