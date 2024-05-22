package repository

type UserBalanceAccreditation struct {
	UserID uint64  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type AccreditationRepository interface {
	AccreditateUserBalance(userBalanceAccreditation UserBalanceAccreditation) error
}
