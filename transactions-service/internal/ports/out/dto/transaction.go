package dto

type TransactionResponse struct {
	ID           string `json:"id"`
	UserID       uint64 `json:"user_id"`
	Status       string `json:"status"`
	StatusDetail string `json:"status_detail"`
}
