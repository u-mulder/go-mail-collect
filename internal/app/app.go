package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"mailcollector/internal/web"
)

const defaultPort = ":80"

type App struct {
	// instance that supports incomig web-requests
	e *echo.Echo

	// TODO
	// mf *Mailfetch
	// t *Transport
	// TODO add some cool looger
}

func New() *App {
	return &App{
		e: echo.New(),
	}
}

func (a *App) Start() error {
	var err error

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
	log.Printf("app.Stop")
	
}
