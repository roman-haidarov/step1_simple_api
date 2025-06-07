package main

import (
	"context"
	"os"
	"os/signal"
	"step1_simple_api/internal/logger"
	"step1_simple_api/internal/server"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.Init(zerolog.DebugLevel)
	srv, err := server.New()

	if err != nil {
		log.Fatal().Err(err).Msg("can not create server")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := srv.Run(ctx); err != nil {
			log.Fatal().Err(err).Msg("server error")
		}
	}()

	log.Info().Msg("Server started")

	<-ctx.Done()
	srv.Shutdown()
}
