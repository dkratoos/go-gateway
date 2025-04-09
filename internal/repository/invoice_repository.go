package repository

import (
	"database/sql"
	"time"

	"github.com/dkratoos/go-gateway/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(invoice *domain.Invoice) error {
	smtmt, err := r.db.Prepare("INSERT INTO invoices (id, account_id, status, description, payment_method, amount, card_last_digits, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	defer smtmt.Close()

	_, err = smtmt.Exec(invoice.ID, invoice.AccountID, invoice.Status, invoice.Description, invoice.PaymentMethod, invoice.Amount, invoice.CardLastDigits, invoice.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *InvoiceRepository) GetByID(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	var createdAt time.Time

	row := r.db.QueryRow("SELECT id, account_id, status, description, payment_method, amount, card_last_digits, created_at FROM invoices WHERE id = $1", id).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentMethod,
		&invoice.Amount,
		&invoice.CardLastDigits,
		&createdAt,
	)

	if row == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}

	if row != nil {
		return nil, row
	}

	invoice.CreatedAt = createdAt

	return &invoice, nil
}

func (r *InvoiceRepository) GetByAccountID(accountID string) ([]*domain.Invoice, error) {
	rows, err := r.db.Query("SELECT id, account_id, status, description, payment_method, amount, card_last_digits, created_at FROM invoices WHERE account_id = $1", accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := []*domain.Invoice{}

	for rows.Next() {
		var invoice domain.Invoice
		var createdAt time.Time

		err = rows.Scan(
			&invoice.ID,
			&invoice.AccountID,
			&invoice.Status,
			&invoice.Description,
			&invoice.PaymentMethod,
			&invoice.Amount,
			&invoice.CardLastDigits,
			&createdAt,
		)

		if err != nil {
			return nil, err
		}

		invoice.CreatedAt = createdAt
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func (r *InvoiceRepository) UpdateStatus(invoice *domain.Invoice) error {
	smtmt, err := r.db.Prepare("UPDATE invoices SET status = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer smtmt.Close()

	_, err = smtmt.Exec(invoice.Status, invoice.ID)
	if err != nil {
		return err
	}

	return nil
}
