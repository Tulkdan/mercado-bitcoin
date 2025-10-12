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

	orderService := service.NewOrderService(repository.New())
	server := web.NewServer(port, orderService)
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
