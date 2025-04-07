package service

import (
	"github.com/dkratoos/go-gateway/internal/domain"
	"github.com/dkratoos/go-gateway/internal/dto"
)

type AccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) Create(input dto.CreateAccountRequest) (*dto.CreateAccountResponse, error) {
	account := dto.ToDomain(input)

	existingAccount, err := s.accountRepository.GetByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}

	err = s.accountRepository.Create(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomain(account)
	return &output, nil
}

func (s *AccountService) GetByAPIKey(apiKey string) (*dto.CreateAccountResponse, error) {
	account, err := s.accountRepository.GetByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomain(account)
	return &output, nil
}

func (s *AccountService) GetByID(id string) (*dto.CreateAccountResponse, error) {
	account, err := s.accountRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomain(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, balance float64) (*dto.CreateAccountResponse, error) {
	account, err := s.accountRepository.GetByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(balance)

	err = s.accountRepository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromDomain(account)
	return &output, nil
}
