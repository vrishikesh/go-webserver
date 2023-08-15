package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vrishikesh/go-webserver/handlers"
	"github.com/vrishikesh/go-webserver/helpers"
	"github.com/vrishikesh/go-webserver/middlewares"
	"github.com/vrishikesh/go-webserver/router"
)

func main() {
	rr := InitRoutes()
	loggerMiddleware := middlewares.NewRequestLogger(rr)

	log.Printf("starting server")
	if err := http.ListenAndServe(":8080", loggerMiddleware); err != nil {
		log.Fatal(err)
	}
}

func InitRoutes() *router.RegexRouter {
	rr := router.NewRegexRouter()

	// random routes
	rr.Handler(helpers.TimeRouteRegex, http.MethodGet, handlers.NewTimeHandler(time.RFC1123))
	rr.Handler(helpers.PublicRouteRegex, http.MethodGet, http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// user related routes
	rr.HandlerFunc(helpers.UsersRouteRegex, http.MethodGet, handlers.HandleGetUsersRoute)
	rr.HandlerFunc(helpers.UsersRouteRegex, http.MethodPost, handlers.HandleCreateUserRoute)
	rr.HandlerFunc(helpers.UserRouteRegex, http.MethodGet, handlers.HandleGetUserRoute)
	rr.HandlerFunc(helpers.UserRouteRegex, http.MethodPut, handlers.HandleUpdateUserRoute)
	rr.HandlerFunc(helpers.UserRouteRegex, http.MethodDelete, handlers.HandleRemoveUserRoute)

	return rr
}
