package domain

type AccountRepository interface {
	Create(account *Account) error
	GetByAPIKey(apiKey string) (*Account, error)
	GetByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

type InvoiceRepository interface {
	Create(invoice *Invoice) error
	GetByID(id string) (*Invoice, error)
	GetByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
