package domain

type AccountRepository interface {
	Create(account *Account) error
	GetByAPIKey(apiKey string) (*Account, error)
	GetByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}
