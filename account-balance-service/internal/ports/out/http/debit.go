package http

type DebitResponse struct {
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}
