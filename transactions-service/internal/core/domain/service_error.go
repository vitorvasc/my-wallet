package domain

import (
	"net/http"
)

type ServiceError interface {
	error
	GetCode() int
	GetMessage() string
}

type serviceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrObtainingTransaction = newServiceError(http.StatusInternalServerError, "error obtaining transaction")
	ErrParsingTransaction   = newServiceError(http.StatusInternalServerError, "error parsing transaction")
	ErrCreatingTransaction  = newServiceError(http.StatusInternalServerError, "error creating transaction")
	ErrUpdatingTransaction  = newServiceError(http.StatusInternalServerError, "error updating transaction")
	ErrInvalidUsersInvolved = newServiceError(http.StatusUnprocessableEntity, "invalid users involved")

	ErrObtainingUserByID   = newServiceError(http.StatusFailedDependency, "error obtaining user")
	ErrUserNotFound        = newServiceError(http.StatusNotFound, "user not found")
	ErrParsingUserResponse = newServiceError(http.StatusInternalServerError, "error parsing user response")

	ErrObtainingAccountBalance       = newServiceError(http.StatusFailedDependency, "error obtaining account balance")
	ErrParsingAccountBalanceResponse = newServiceError(http.StatusInternalServerError, "error parsing account balance response")

	ErrObtainingUserTransactions = newServiceError(http.StatusInternalServerError, "error obtaining user transactions")

	ErrProcessingTransactionStrategyNotFound = newServiceError(http.StatusUnprocessableEntity, "error processing transaction strategy not found")

	ErrInvalidAmount           = newServiceError(http.StatusBadRequest, "invalid amount")
	ErrInsufficientFunds       = newServiceError(http.StatusUnprocessableEntity, "insufficient funds")
	ErrProcessingAccreditation = newServiceError(http.StatusInternalServerError, "error processing accreditation")
)

func newServiceError(code int, message string) ServiceError {
	return &serviceError{
		Code:    code,
		Message: message,
	}
}

func (e *serviceError) Error() string {
	return e.Message
}

func (e *serviceError) GetCode() int {
	return e.Code
}

func (e *serviceError) GetMessage() string {
	return e.Message
}
