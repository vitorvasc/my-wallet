package dto

import (
	"time"
	"transactions-service/internal/core/domain"
)

type TransactionResponse struct {
	ID           string                   `json:"id"`
	UserID       uint64                   `json:"user_id"`
	Description  string                   `json:"description,omitempty"`
	Type         string                   `json:"type"`
	From         *TransactionResponseFrom `json:"from,omitempty"`
	To           *TransactionResponseTo   `json:"to,omitempty"`
	Status       string                   `json:"status"`
	StatusDetail string                   `json:"status_detail"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
}

type TransactionResponseFrom struct {
	UserID        uint64  `json:"user_id,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	PaymentMethod string  `json:"payment_method,omitempty"`
}

type TransactionResponseTo struct {
	UserID uint64  `json:"user_id,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

func TransactionResponseFromDomain(transaction *domain.Transaction) *TransactionResponse {
	return &TransactionResponse{
		ID:          transaction.ID.Hex(),
		UserID:      transaction.UserID,
		Description: transaction.Description,
		Type:        string(transaction.Type),
		From: &TransactionResponseFrom{
			UserID:        transaction.From.UserID,
			Amount:        transaction.From.Amount,
			PaymentMethod: transaction.From.PaymentMethod,
		},
		To: &TransactionResponseTo{
			UserID: transaction.To.UserID,
			Amount: transaction.To.Amount,
		},
		Status:       string(transaction.Status),
		StatusDetail: string(transaction.StatusDetail),
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,
	}
}
