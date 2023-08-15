package app

import (
	"log"
	"os"
)

var Logger *log.Logger

func NewLogger() (*log.Logger, error) {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile), nil
}
