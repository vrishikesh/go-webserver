package middlewares

import (
	"log"
	"net/http"
	"time"
)

type requestLogger struct {
	http.Handler
}

func NewRequestLogger(handle http.Handler) *requestLogger {
	return &requestLogger{Handler: handle}
}

func (l *requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("path [%s], method [%s], execution time [%s]", r.URL.Path, r.Method, time.Since(start))
}
