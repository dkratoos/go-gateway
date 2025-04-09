package server

import (
	"net/http"

	"github.com/dkratoos/go-gateway/internal/service"
	"github.com/dkratoos/go-gateway/internal/web/handlers"
	"github.com/dkratoos/go-gateway/internal/web/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	authMiddleware *middleware.AuthMiddleware
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Route("/accounts", func(r chi.Router) {
		r.Post("/", accountHandler.CreateAccount)
		r.Get("/", accountHandler.GetAccount)
	})

	s.router.Route("/invoices", func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/", invoiceHandler.CreateInvoice)
		r.Get("/{id}", invoiceHandler.GetInvoice)
		r.Get("/account/{accountId}", invoiceHandler.GetInvoices)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
