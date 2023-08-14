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
	// srv := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router.Router(),
	// }

	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }

	mux := http.NewServeMux()
	mux.Handle("/time/", handlers.NewTimeHandler(time.RFC1123))
	mux.Handle("/users/", router.NewUserHandler())
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	wrapper := middlewares.NewRequestLogger(mux)

	log.Printf("starting server")
	log.Fatal(http.ListenAndServe(":8080", wrapper))
}
