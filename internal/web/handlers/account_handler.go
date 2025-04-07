package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dkratoos/go-gateway/internal/dto"
	"github.com/dkratoos/go-gateway/internal/service"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
	w.WriteHeader(http.StatusCreated)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "X-API-Key header is required", http.StatusBadRequest)
		return
	}

	account, err := h.accountService.GetByAPIKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "X-API-Key header is required", http.StatusBadRequest)
		return
	}

	var input dto.UpdateBalanceRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.UpdateBalance(apiKey, input.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
