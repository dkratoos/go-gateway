package dto

import (
	"time"

	"github.com/dkratoos/go-gateway/internal/domain"
)

type CreateAccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateAccountResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	APIKey    string    `json:"api_key,omitempty"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomain(input CreateAccountRequest) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromDomain(account *domain.Account) CreateAccountResponse {
	return CreateAccountResponse{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		APIKey:    account.APIKey,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

type GetAccountResponse struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	APIKey  string  `json:"api_key,omitempty"`
	Balance float64 `json:"balance"`
}

type UpdateBalanceRequest struct {
	Balance float64 `json:"balance"`
}

type UpdateBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type UpdateBalanceError struct {
	Error string `json:"error"`
}
