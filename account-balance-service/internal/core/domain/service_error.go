package domain

import (
	"net/http"
)

type ServiceError struct {
	error
	Code    int
	Message string
}

var (
	ErrUserNotFound            = newServiceError(http.StatusNotFound, "invalid user")
	ErrInsufficientFunds       = newServiceError(http.StatusBadRequest, "insufficient funds")
	ErrObtainingAccountBalance = newServiceError(http.StatusInternalServerError, "error obtaining account balance")
	ErrAccreditingValue        = newServiceError(http.StatusInternalServerError, "unexpected accreditation error")
	ErrDebitingValue           = newServiceError(http.StatusInternalServerError, "unexpected debiting error")
)

func newServiceError(code int, message string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}
