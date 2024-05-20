package db

type AccountBalanceRepository interface {
	GetAccountBalance(userID uint64) (float64, error)
	UpdateAccountBalance(userID uint64, amount float64) error
}
