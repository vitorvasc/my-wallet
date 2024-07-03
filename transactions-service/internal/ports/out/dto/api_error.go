package dto

import (
	"fmt"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApiError(code int, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}

func NewRequiredFieldApiError(code int, fieldName string) ApiError {
	return ApiError{
		Code:    code,
		Message: fmt.Sprintf("%s is required", fieldName),
	}
}

func NewInvalidFieldApiError(code int, fieldName string) ApiError {
	return ApiError{
		Code:    code,
		Message: fmt.Sprintf("%s must be valid", fieldName),
	}
}
