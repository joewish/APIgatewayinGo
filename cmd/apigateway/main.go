package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vjoewish/APIGateway/internal/config"
)

func main() {

	logrus.Info("Loading configuration from config file")
	configs, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error lodding the config file: %w", err))
	}
	serverError := http.ListenAndServe(configs.Server.Port, nil)
	if serverError != nil {
		panic(fmt.Errorf("fatal error in sarting the server %w", serverError))
	}
	logrus.Info("server started")
}
