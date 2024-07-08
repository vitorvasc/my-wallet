package repository

type AccountBalanceRepository interface {
	GetAccountBalance(userID uint64) (float64, error)
	Credit(userID uint64, amount float64) error
	Debit(userID uint64, amount float64) error
}
