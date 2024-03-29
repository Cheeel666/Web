package main

// Swag: export PATH=$(go env GOPATH)/bin:$PATH
import (
	"os"
	"time"
	"web/config"
	"web/internal/auth"
	"web/internal/handlers"
	"web/internal/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const configPath = "config/config.json"

// @title SkOpen API
// @version 1.0
// @description Swagger API for Golang agregator project.

// @contact.name Ilya Chelyadinov
// @contact.email ilchel1992@gmail.com

// @host localhost:5005
// @BasePath /api/v1
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
		[]byte(cfg.SignKey),
		20*time.Minute,
	)
	logrus.Info("Connected to DB")
	listenErr := make(chan error, 1)

	go server.SetupServer(cfg, service).Run(":5007")
	select {
	case err = <-listenErr:
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

	}
}
