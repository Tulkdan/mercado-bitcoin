package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Tulkdan/central-limit-order-book/internal/repository"
	"github.com/Tulkdan/central-limit-order-book/internal/service"
	"github.com/Tulkdan/central-limit-order-book/internal/web"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	port := getEnv("PORT", "8000")

	repository := repository.New()

	orderService := service.NewOrderService(repository)
	accountService := service.NewAccountService(repository)

	server := web.NewServer(port, orderService, accountService)
	server.ConfigureRouter()

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- server.Start(ctx)
	}()

	select {
	case <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}

	server.Shutdown()
}
