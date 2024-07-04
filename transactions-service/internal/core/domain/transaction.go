package domain

import "time"

type Transaction struct {
	ID           string                  `bson:"-"`
	UserID       uint64                  `bson:"user_id"`
	Description  string                  `bson:"description"`
	Type         string                  `bson:"type"`
	From         TransactionFrom         `bson:"from"`
	To           TransactionTo           `bson:"to"`
	Status       TransactionStatus       `bson:"status"`
	StatusDetail TransactionStatusDetail `bson:"status_detail"`
	CreatedAt    time.Time               `bson:"created_at"`
	UpdatedAt    time.Time               `bson:"updated_at"`
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
	UserID        uint64  `bson:"user_id"`
	Amount        float64 `bson:"amount"`
	PaymentMethod string  `bson:"payment_method"`
}

type TransactionTo struct {
	UserID uint64  `bson:"user_id"`
	Amount float64 `bson:"amount"`
}
