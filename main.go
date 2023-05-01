package mylogger

import (
	"log"
	"time"
)

func LogInfo(message string) {
	log.Printf("INFO %s - %v", time.Now(), message)
}

func LogWarning(message string) {
	log.Printf("Warning %s - %v", time.Now(), message)
}
