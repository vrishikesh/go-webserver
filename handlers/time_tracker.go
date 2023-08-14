package handlers

import (
	"log"
	"time"
)

func TimeTracker(start time.Time) {
	log.Printf("execution time [%s]", time.Since(start))
}
