package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vrishikesh/go-webserver/handlers"
	"github.com/vrishikesh/go-webserver/middlewares"
	"github.com/vrishikesh/go-webserver/router"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/time/", handlers.NewTimeHandler(time.RFC1123))
	mux.Handle("/users/", router.NewUserHandler())
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	requestLoggerMiddleware := middlewares.NewRequestLogger(mux)

	log.Printf("starting server")
	if err := http.ListenAndServe(":8080", requestLoggerMiddleware); err != nil {
		log.Fatal(err)
	}
}
