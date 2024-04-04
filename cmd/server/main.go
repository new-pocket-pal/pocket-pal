package main

import (
	"pocket-pal/src/app"
	"pocket-pal/src/utils/wrapper"
)

func main() {
	app, err := app.Initialize()
	if err != nil {
		panic(err)
	}
	go app.Run()
	// Graceful
	wrapper.GracefulShutdown(app.Echo)
}
