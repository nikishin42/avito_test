package main

import (
	"server/server/application"
	"server/server/config"
	"server/server/logger"
	"server/server/provider"
	"server/server/repos"
	"server/server/services/wallet/api"
	"server/server/services/wallet/usecases"
)

func main() {
	conf := config.LoadConfig()
	l := logger.New()
	p := provider.New(conf, l)
	defer p.Close()

	repositories := repos.New(p, l)
	service := usecases.New(repositories, l)
	walletService := api.New(service, l)
	app := application.New(conf, walletService, l)
	app.Start()
}
