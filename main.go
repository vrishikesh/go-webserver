package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/vrishikesh/go-webserver/app"
	"github.com/vrishikesh/go-webserver/middlewares"
	"github.com/vrishikesh/go-webserver/router"
)

//go:embed configs
var embedFS embed.FS

func main() {
	var err error
	app.EmbedFS = embedFS

	app.Logger, err = app.NewLogger()
	if err != nil {
		log.Fatalf("could not init logger: %s", err)
	}

	app.Config, err = app.NewConfig(app.EmbedFS, app.Logger)
	if err != nil {
		app.Logger.Fatalf("could not init config: %s", err)
	}

	rr := router.SetupRoutes()
	loggerMiddleware := middlewares.NewRequestLogger(rr)

	app.Logger.Printf("starting server")
	app.Logger.Fatal(http.ListenAndServe(app.Config.AppHost, loggerMiddleware))
}
