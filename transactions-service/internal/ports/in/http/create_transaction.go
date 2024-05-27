package http

type CreateTransactionRequest struct {
	Type        TransactionType `json:"type" binding:"required"`
	Description string          `json:"description"`
	From        TransactionFrom `json:"from" binding:"required"`
	To          TransactionTo   `json:"to" binding:"required"`
}

type TransactionType string

var (
	AccountTransfer TransactionType = "account_transfer"
	Deposit         TransactionType = "deposit"
	Withdrawal      TransactionType = "withdrawal"
	BillPayment     TransactionType = "bill_payment"
)

type PaymentMethod string

var (
	AccountBalance PaymentMethod = "account_balance"

	PureCash PaymentMethod = "pure_cash"
)

type TransactionFrom struct {
	UserID        uint64        `json:"user_id" binding:"required"`
	Amount        float64       `json:"amount" binding:"required"`
	PaymentMethod PaymentMethod `json:"payment_method" binding:"required"`
}

type TransactionTo struct {
	UserID uint64  `json:"user_id"`
	Amount float64 `json:"amount" binding:"required"`
}
