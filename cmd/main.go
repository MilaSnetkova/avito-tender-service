package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tender-service/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		cancel()
	}()

	a, err := app.New(ctx)
	if err != nil {
		fmt.Println("failed to initialize app", err)
		os.Exit(1)
	}

	if err = a.Run(); err != nil {
		defer os.Exit(1)
		fmt.Println("failed to run app", err)
	}
}
