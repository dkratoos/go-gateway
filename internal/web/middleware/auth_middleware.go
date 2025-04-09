package middleware

import (
	"context"
	"net/http"

	"github.com/dkratoos/go-gateway/internal/service"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{accountService: accountService}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey == "" {
			http.Error(w, "X-API-Key header is required", http.StatusUnauthorized)
			return
		}

		account, err := m.accountService.GetByAPIKey(apiKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "account", account)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
