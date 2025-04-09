package dto

import (
	"time"

	"github.com/dkratoos/go-gateway/internal/domain"
)

const (
	StatusPending = string(domain.StatusPending)
	StatusPaid    = string(domain.StatusPaid)
	StatusFailed  = string(domain.StatusFailed)
)

type CreateInvoiceRequest struct {
	APIKey        string  `json:"api_key"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	PaymentMethod string  `json:"payment_method"`
	CreditCard    string  `json:"credit_card"`
	CVV           string  `json:"cvv"`
	ExpiryMonth   int     `json:"expiry_month"`
	ExpiryYear    int     `json:"expiry_year"`
	HolderName    string  `json:"holder_name"`
}

type CreateInvoiceResponse struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentMethod  string    `json:"payment_method"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
}

func InvoiceToDomain(input CreateInvoiceRequest, accountID string) (*domain.Invoice, error) {
	creditCard := &domain.CreditCard{
		Number:      input.CreditCard,
		CVV:         input.CVV,
		ExpiryMonth: input.ExpiryMonth,
		ExpiryYear:  input.ExpiryYear,
		HolderName:  input.HolderName,
	}
	return domain.NewInvoice(accountID, input.Amount, input.Description, input.PaymentMethod, creditCard)
}

func InvoiceFromDomain(invoice *domain.Invoice) CreateInvoiceResponse {
	return CreateInvoiceResponse{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentMethod:  invoice.PaymentMethod,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
	}
}
