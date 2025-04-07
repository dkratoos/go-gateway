package repository

import (
	"database/sql"
	"time"

	"github.com/dkratoos/go-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(account *domain.Account) error {
	smtmt, err := r.db.Prepare("INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer smtmt.Close()

	_, err = smtmt.Exec(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) GetByAPIKey(apiKey string) (*domain.Account, error) {
	return r.getAccount("api_key", apiKey)
}

func (r *AccountRepository) GetByID(id string) (*domain.Account, error) {
	return r.getAccount("id", id)
}

func (r *AccountRepository) getAccount(field, value string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	row := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE "+field+" = $1", value).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)

	if row == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if row != nil {
		return nil, row
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", account.ID).Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = $1, updated_at = $2 WHERE id = $3", currentBalance+account.Balance, time.Now(), account.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
