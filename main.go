package main

import (
	"log"
	"net/http"

	"github.com/vrishikesh/go-webserver/router"
)

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(router.Router),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
