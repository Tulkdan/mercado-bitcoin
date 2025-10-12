package web

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/Tulkdan/central-limit-order-book/internal/service"
	"github.com/Tulkdan/central-limit-order-book/internal/web/handler"
	"github.com/Tulkdan/central-limit-order-book/internal/web/middleware"
)

type Server struct {
	port   string
	router *http.ServeMux
	server *http.Server

	orderService   *service.OrderService
	accountService *service.AccountService
}

func NewServer(port string, orderService *service.OrderService, accountService *service.AccountService) *Server {
	return &Server{
		port:           port,
		orderService:   orderService,
		accountService: accountService,
	}
}

func (s *Server) ConfigureRouter() {
	mux := http.NewServeMux()

	orderHandler := handler.NewOrderHandler(s.orderService)
	accountHandler := handler.NewAccountHandler(s.accountService)

	mux.HandleFunc("POST /order", middleware.WithRequestId(orderHandler.Create))
	mux.HandleFunc("DELETE /order/{id}", middleware.WithRequestId(orderHandler.Cancel))
	mux.HandleFunc("GET /account/{id}", middleware.WithRequestId(accountHandler.GetBalance))

	s.router = mux
}

func (s *Server) Start(ctx context.Context) error {
	s.server = &http.Server{
		Addr:         ":" + s.port,
		Handler:      s.router,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.server.Shutdown(context.Background())
}
