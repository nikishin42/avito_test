package application

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"server/server/config"
	"server/server/http/middleware"
	"server/server/http/router"
	"server/server/logger"
	"server/server/services"
	"server/server/services/wallet/api"
	"syscall"
)

type Application struct {
	server  *http.Server
	service services.Service
	l       logger.Logger
}

func New(config *config.Config, service *api.WalletService, l logger.Logger) *Application {
	return &Application{
		server: &http.Server{
			Addr:    ":" + config.ServerPort,
			Handler: router.Router(middleware.New(l), service),
		},
		service: service,
		l:       l,
	}
}

func (app *Application) Start() {
	app.service.Start()
	listenErr := make(chan error, 1)
	go func() {
		listenErr <- app.server.ListenAndServe()
	}()
	app.l.InfoLog.Println("http server started at port", app.server.Addr)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-listenErr:
		if err != nil {
			app.l.ErrorLog.Fatalln(err)
		}
	case s := <-osSignals:
		app.l.InfoLog.Println("SIGNAL:", s.String())
		app.l.InfoLog.Println("Stopping Wallet service")
		if err := app.server.Shutdown(context.Background()); err != nil {
			app.l.ErrorLog.Fatalln(err)
		}
	}
	close(listenErr)
	close(osSignals)
	app.l.InfoLog.Println("Wallet service stopped")
}
