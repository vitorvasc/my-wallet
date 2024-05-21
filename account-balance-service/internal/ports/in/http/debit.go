package http

type DebitRequest struct {
	Amount float64 `json:"amount" binding:"required"`
	UserID uint64  `json:"user_id" binding:"required"`
}
