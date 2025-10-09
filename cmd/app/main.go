package main

import (
	"context"
	"os"
	"os/signal"

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
	server := web.NewServer(port)
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
