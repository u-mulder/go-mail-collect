package app

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"mailcollector/internal/web"
	"mailcollector/internal/logger"
)

const defaultPort = ":80"

type App struct {
	// instance that supports incomig web-requests
	e *echo.Echo
	l zerolog.Logger

	// TODO
	// mf *Mailfetch
	// t *Transport
}

func New() *App {
	return &App{
		e: echo.New(),
		l: logger.New(),
	}
}

func (a *App) Start() error {
	var err error

	a.l.Info().Msg("Starting app...")

	// init api endpoints
	if err = web.RegisterEndpoints(a.e); err != nil {
		return err
	}
	a.e.Logger.Fatal(a.e.Start(defaultPort))

	//a.mf.Init()
	//a.t.Init()

	// TODO remove
	log.Printf("app.Started")

	return nil
}

func (a *App) Stop() {
	// TODO do we need it?
	//log.Printf("app.Stop")
	
}
