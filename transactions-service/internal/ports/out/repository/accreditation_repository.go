package repository

import "transactions-service/internal/ports/out/dto"

type AccreditationRepository interface {
	CreateUserBalanceCredit(accountCreditCreation dto.AccountCreditCreation) error
}
