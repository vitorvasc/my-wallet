package domain

import "time"

type Transaction struct {
	ID           uint64
	UserID       uint64
	Description  string
	Type         string
	From         TransactionFrom
	To           TransactionTo
	Status       TransactionStatus
	StatusDetail TransactionStatusDetail
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TransactionStatus string

var (
	Approved TransactionStatus = "approved"
	Rejected TransactionStatus = "rejected"
	Reverted TransactionStatus = "reverted"
)

type TransactionStatusDetail string

var (
	Processed                    TransactionStatusDetail = "processed"
	Refunded                     TransactionStatusDetail = "refunded"
	InvalidAmount                TransactionStatusDetail = "invalid_amount"
	AccreditationProcessingError TransactionStatusDetail = "accreditation_processing_error"
	InsufficientFunds            TransactionStatusDetail = "insufficient_funds"
	UnexpectedError              TransactionStatusDetail = "unexpected_error"
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
