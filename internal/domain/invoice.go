package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusPaid    Status = "paid"
	StatusFailed  Status = "failed"
)

type Invoice struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Status         Status    `json:"status"`
	Description    string    `json:"description"`
	PaymentMethod  string    `json:"payment_method"`
	Amount         float64   `json:"amount"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreditCard struct {
	Number      string `json:"number"`
	CVV         string `json:"cvv"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	HolderName  string `json:"holder_name"`
}

func NewInvoice(accountID string, amount float64, description string, paymentMethod string, card *CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Status:         StatusPending,
		Description:    description,
		Amount:         amount,
		PaymentMethod:  paymentMethod,
		CardLastDigits: card.Number[len(card.Number)-4:],
		CreatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Pay() error {
	if i.Amount > 1000 {
		return nil
	}

	if rand.New(rand.NewSource(time.Now().UnixNano())).Float64() < 0.7 {
		i.Status = StatusPaid
	} else {
		i.Status = StatusFailed
	}

	return nil
}

func (i *Invoice) UpdateStatus(status Status) error {
	if status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = status

	return nil
}
