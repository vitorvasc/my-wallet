package http

type UpdateTransactionRequest struct {
	ID     uint64 `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}
