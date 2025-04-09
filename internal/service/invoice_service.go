package service

import (
	"github.com/dkratoos/go-gateway/internal/domain"
	"github.com/dkratoos/go-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{invoiceRepository: invoiceRepository, accountService: accountService}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceRequest) (*dto.CreateInvoiceResponse, error) {
	account, err := s.accountService.GetByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.InvoiceToDomain(input, account.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Pay(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusPaid {
		if _, err := s.accountService.UpdateBalance(account.APIKey, invoice.Amount); err != nil {
			return nil, err
		}
	}

	err = s.invoiceRepository.Create(invoice)
	if err != nil {
		return nil, err
	}

	output := dto.InvoiceFromDomain(invoice)
	return &output, nil
}

func (s *InvoiceService) GetByID(id string, apiKey string) (*dto.CreateInvoiceResponse, error) {
	invoice, err := s.invoiceRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	account, err := s.accountService.GetByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != account.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	output := dto.InvoiceFromDomain(invoice)
	return &output, nil
}

func (s *InvoiceService) GetByAccountID(accountID string) ([]*dto.CreateInvoiceResponse, error) {
	invoices, err := s.invoiceRepository.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.CreateInvoiceResponse, len(invoices))
	for i, invoice := range invoices {
		temp := dto.InvoiceFromDomain(invoice)
		output[i] = &temp
	}
	return output, nil
}

func (s *InvoiceService) GetByAPIKey(apiKey string) ([]*dto.CreateInvoiceResponse, error) {
	account, err := s.accountService.GetByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.GetByAccountID(account.ID)
}
