package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("invalid user id")
	ErrInsufficientFunds = errors.New("insufficient funds")
)
