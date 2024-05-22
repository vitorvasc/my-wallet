package domain

import "time"

type Transaction struct {
	ID           uint64
	Type         string
	UserID       uint64
	From         TransactionFrom
	To           TransactionTo
	Status       TransactionStatus
	StatusDetail string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TransactionStatus string

var (
	Approved TransactionStatus = "approved"
	Rejected TransactionStatus = "rejected"
	Reverted TransactionStatus = "reverted"
)

type TransactionFrom struct {
	UserID        uint64
	Amount        float64
	PaymentMethod string
}

type TransactionTo struct {
	UserID uint64
	Amount float64
}
