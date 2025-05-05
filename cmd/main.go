package main

import (
	"step1_simple_api/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {
	srv, err := server.New()

	if err != nil {
		logrus.Error(err)
	}

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Error(err)
		}
	}()

	logrus.Info("Server started")
	select {}
}
