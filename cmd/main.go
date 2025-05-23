package main

import (
	"context"
	"os"
	"os/signal"
	"step1_simple_api/internal/server"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	srv, err := server.New()

	if err != nil {
		logrus.Error(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := srv.Run(ctx); err != nil {
			logrus.Error(err)
		}
	}()

	logrus.Info("Server started")

	<-ctx.Done()
	srv.Shutdown()
}
