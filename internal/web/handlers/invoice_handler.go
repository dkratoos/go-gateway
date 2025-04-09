package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dkratoos/go-gateway/internal/domain"
	"github.com/dkratoos/go-gateway/internal/dto"
	"github.com/dkratoos/go-gateway/internal/service"
	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = r.Header.Get("X-API-Key")
	if input.APIKey == "" {
		http.Error(w, "X-API-Key header is required", http.StatusBadRequest)
		return
	}

	invoice, err := h.invoiceService.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
	w.WriteHeader(http.StatusCreated)
}

func (h *InvoiceHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "X-API-Key header is required", http.StatusBadRequest)
		return
	}

	invoice, err := h.invoiceService.GetByID(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusUnauthorized)
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "accountId")

	if accountID == "" {
		http.Error(w, "account_id is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "X-API-Key header is required", http.StatusBadRequest)
		return
	}

	invoices, err := h.invoiceService.GetByAccountID(accountID)
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
