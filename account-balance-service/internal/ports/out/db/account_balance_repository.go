package db

type AccountBalanceRepository interface {
	GetAccountBalance(userID uint64) (float64, error)
	Accredit(userID uint64, amount float64) error
	Debit(userID uint64, amount float64) error
}
