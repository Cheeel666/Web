package main

import (
	"time"
	"web/config"
	"web/internal/auth"
	"web/internal/handlers"
	"web/internal/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const configPath = "config/config.json"

func main() {
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	var service handlers.Service
	err = service.DB.Connect(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	defer service.DB.Close()

	service.Auth = *auth.NewAuthorizer(
		service.DB,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl")*time.Second,
	)
	logrus.Info("Connected to DB")

	server.SetupServer(cfg, service).Run(cfg.Port)
}
