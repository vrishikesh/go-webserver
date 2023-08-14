package router

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/vrishikesh/go-webserver/handlers"
)

func Router(w http.ResponseWriter, r *http.Request) {
	defer handlers.TimeTracker(time.Now())
	log.Printf("request: path [%s], method [%s]", r.URL.Path, r.Method)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("could not read request body: %s", err)
		fmt.Fprint(w, "something went wrong")
		return
	}

	log.Printf("request body: %s", string(reqBody))
	var res any

	switch {
	case strings.Index(r.URL.Path, "/users") == 0:
		res, err = UserRouter(r, reqBody)
	default:
		res, err = handlers.ResourceNotFound(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: %s", err)
		fmt.Fprint(w, "something went wrong")
		return
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal response: %s", err)
		fmt.Fprint(w, "something went wrong")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Printf("response body: %s", string(resBody))
	fmt.Fprint(w, string(resBody))
}
