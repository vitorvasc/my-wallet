package repository

type UserBalanceCredit struct {
	UserID uint64  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type AccreditationRepository interface {
	CreateUserBalanceCredit(userBalanceAccreditation UserBalanceCredit) error
}
