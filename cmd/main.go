package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	appPkg "mailcollector/internal/app"
)

func main() {
	app := appPkg.New()

	if err := app.Start(); err != nil {
		panic(fmt.Sprintf("Error starting mailcollector: %s", err))
	}

	appStopCh := make(chan os.Signal, 1)

	signal.Notify(appStopCh)
	for sig := range appStopCh {
		switch sig {
		case os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM:
			app.Stop()
			return
		default:
		}
	}
}