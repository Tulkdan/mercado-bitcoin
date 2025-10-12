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

	orderService *service.OrderService
}

func NewServer(port string, orderService *service.OrderService) *Server {
	return &Server{
		port:         port,
		orderService: orderService,
	}
}

func (s *Server) ConfigureRouter() {
	mux := http.NewServeMux()

	paymentsHandler := handler.NewOrderHandler(s.orderService)

	mux.HandleFunc("POST /order", middleware.WithRequestId(paymentsHandler.Create))
	// r.HandleFunc("POST /refunds", func(http.ResponseWriter, *http.Request) {})
	// r.HandleFunc("GET /payments/{id}", func(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	// })

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
