package handlers

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func NewTimeHandler(format string) *timeHandler {
	return &timeHandler{format: format}
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	_, err := w.Write([]byte(tm))
	if err != nil {
		log.Printf("could not write byte on stdout: %s", err)
	}
}
