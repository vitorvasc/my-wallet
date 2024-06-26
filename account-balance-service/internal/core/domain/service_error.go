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
	ErrUserNotFound            = newServiceError(http.StatusNotFound, "invalid user")
	ErrInsufficientFunds       = newServiceError(http.StatusBadRequest, "insufficient funds")
	ErrObtainingAccountBalance = newServiceError(http.StatusInternalServerError, "error obtaining account balance")
	ErrAccreditingValue        = newServiceError(http.StatusInternalServerError, "unexpected accreditation error")
	ErrDebitingValue           = newServiceError(http.StatusInternalServerError, "unexpected debiting error")
	ErrObtainingUser           = newServiceError(http.StatusInternalServerError, "error obtaining user")
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
