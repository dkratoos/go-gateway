package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrDuplicatedAPIKey = errors.New("duplicated api key")

	ErrInvoiceNotFound = errors.New("invoice not found")

	ErrUnauthorized = errors.New("unauthorized")

	ErrInvalidAmount = errors.New("invalid amount")

	ErrInvalidStatus = errors.New("invalid status")

	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
